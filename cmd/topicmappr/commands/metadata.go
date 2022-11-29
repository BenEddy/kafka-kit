package commands

import (
	"context"
	"encoding/json"
	"fmt"
	"regexp"
	"sort"
	"time"

	"github.com/DataDog/kafka-kit/v4/kafkaadmin"
	"github.com/DataDog/kafka-kit/v4/kafkazk"
	"github.com/DataDog/kafka-kit/v4/mapper"
)

type metadataProvider interface {
	getMatchingTopicNames(topics []string) ([]string, error)
	checkMetaAge(maxAge int) error
	getBrokerMeta(useZk bool) (mapper.BrokerMetaMap, []error)
	getPartitionMaps(topics []string) (*mapper.PartitionMap, error)
	getPartitionMeta() (mapper.PartitionMetaMap, error)
}

type defaultMetadataProvider struct {
	zk kafkazk.Handler
	ka kafkaadmin.KafkaAdmin
}

func (m *defaultMetadataProvider) getMatchingTopicNames(topics []string) ([]string, error) {
	tState, err := m.ka.DescribeTopics(context.Background(), topics)
	if err != nil {
		return []string{}, err
	}
	return tState.List(), nil
}

// checkMetaAge checks the age of the stored partition and broker storage
// metrics data against the tolerated metrics age parameter.
func (m *defaultMetadataProvider) checkMetaAge(maxAge int) error {
	age, err := m.zk.MaxMetaAge()
	if err != nil {
		return fmt.Errorf("Error fetching metrics metadata: %s\n", err)
	}

	if age > time.Duration(maxAge)*time.Minute {
		return fmt.Errorf("Metrics metadata is older than allowed: %s\n", age)
	}
	return nil
}

// getBrokerMeta returns a map of brokers and broker metadata for those
// registered in the cluster state. Optionally, broker metrics can be popularted
// via ZooKeeper.
func (m *defaultMetadataProvider) getBrokerMeta(useZk bool) (mapper.BrokerMetaMap, []error) {
	// Get broker states.
	brokerStates, err := m.ka.DescribeBrokers(context.Background(), false)
	if err != nil {
		return nil, []error{err}
	}

	brokers, _ := mapper.BrokerMetaMapFromStates(brokerStates)

	// Populate metrics.
	if useZk {
		if errs := kafkazk.LoadMetrics(m.zk, brokers); errs != nil {
			return nil, errs
		}
	}

	return brokers, nil
}

func (m *defaultMetadataProvider) getPartitionMaps(topics []string) (*mapper.PartitionMap, error) {
	// Get the topic states.
	tState, err := m.ka.DescribeTopics(context.Background(), topics)
	if err != nil {
		return nil, err
	}

	// Translate it to a mapper object.
	pm, _ := mapper.PartitionMapFromTopicStates(tState)
	sort.Sort(pm.Partitions)

	return pm, nil
}

// ensureBrokerMetrics takes a map of reference brokers and a map of discovered
// broker metadata. Any non-missing brokers in the broker map must be present
// in the broker metadata map and have a non-true MetricsIncomplete value.
func ensureBrokerMetrics(bm mapper.BrokerMap, bmm mapper.BrokerMetaMap) []error {
	errs := []error{}
	for id, b := range bm {
		// Missing brokers won't be found in the brokerMeta.
		if !b.Missing && id != mapper.StubBrokerID && bmm[id].MetricsIncomplete {
			errs = append(errs, fmt.Errorf("Metrics not found for broker %d\n", id))
		}
	}
	return errs
}

// getPartitionMeta returns a map of topic, partition metadata persisted in
// ZooKeeper (via an external mechanism*). This is primarily partition size
// metrics data used for the storage placement strategy.
func (m *defaultMetadataProvider) getPartitionMeta() (mapper.PartitionMetaMap, error) {
	result := mapper.PartitionMetaMap{}

	topics, err := m.ka.DescribeTopics(context.Background(), []string{".*"})
	if err != nil {
		return nil, err
	}

	for _, topic := range topics {
		if _, ok := result[topic.Name]; !ok {
			result[topic.Name] = map[int]*mapper.PartitionMeta{}
		}

		for _, partition := range topic.PartitionStates {
			result[topic.Name][int(partition.ID)] = &mapper.PartitionMeta{
				Size: 1000,
			}
		}
	}

	return result, nil
}

// removeTopics takes a PartitionMap and []*regexp.Regexp of topic name patters.
// Any topic names that match any provided pattern will be removed from the
// PartitionMap and a []string of topics that were found and removed is returned.
func removeTopics(pm *mapper.PartitionMap, r []*regexp.Regexp) []string {
	var removedNames []string

	if len(r) == 0 {
		return removedNames
	}

	// Create a new PartitionList, populate non-removed topics, substitute the
	// existing PartitionList in the PartitionMap.
	newPL := mapper.PartitionList{}

	// Track what's removed.
	removed := map[string]struct{}{}

	// Traverse the partition map.
	for _, p := range pm.Partitions {
		for i, re := range r {
			// If the topic matches any regex pattern, add it to the removed set.
			if re.MatchString(p.Topic) {
				removed[p.Topic] = struct{}{}
				break
			}

			// We've checked all patterns.
			if i == len(r)-1 {
				// Else, it wasn't marked for removal; add it to the new PartitionList.
				newPL = append(newPL, p)
			}
		}
	}

	pm.Partitions = newPL

	for t := range removed {
		removedNames = append(removedNames, t)
	}

	return removedNames
}

type staticMetadataProvider struct {
	Brokers    mapper.BrokerMetaMap `json:"brokers"`
	Partitions []partition          `json:"partitions"`
}

type partition struct {
	Topic     string  `json:"topic"`
	Partition int     `json:"partition"`
	Replicas  []int   `json:"replicas"`
	Size      float64 `json:"size"` // In bytes
}

func newStaticMetadataProvider(data []byte) (*staticMetadataProvider, error) {
	meta := staticMetadataProvider{}

	if err := json.Unmarshal(data, &meta); err != nil {
		return nil, err
	}

	return &meta, nil
}

func (s *staticMetadataProvider) getMatchingTopicNames(topics []string) ([]string, error) {
	re, err := kafkaadmin.StringsToRegex(topics)
	if err != nil {
		return []string{}, nil
	}

	matching := map[string]bool{}
	for _, p := range s.Partitions {
		for _, r := range re {
			if r.MatchString(p.Topic) {
				matching[p.Topic] = true
			}
		}
	}

	result := make([]string, len(matching))
	for topic, _ := range matching {
		result = append(result, topic)
	}
	return result, nil
}

func (s *staticMetadataProvider) checkMetaAge(maxAge int) error {
	return nil
}

func (s *staticMetadataProvider) getBrokerMeta(useZk bool) (mapper.BrokerMetaMap, []error) {
	return s.Brokers, nil
}

func (s *staticMetadataProvider) getPartitionMaps(topics []string) (*mapper.PartitionMap, error) {
	re, err := kafkaadmin.StringsToRegex(topics)
	if err != nil {
		return &mapper.PartitionMap{}, nil
	}

	m := mapper.NewPartitionMap()
	for _, p := range s.Partitions {
		var keep bool
		for _, r := range re {
			if r.MatchString(p.Topic) {
				keep = true
			}
		}

		if keep {
			m.Partitions = append(m.Partitions, mapper.Partition{
				Topic:     p.Topic,
				Partition: p.Partition,
				Replicas:  p.Replicas,
			})
		}
	}

	sort.Sort(m.Partitions)
	return m, nil
}

func (s *staticMetadataProvider) getPartitionMeta() (mapper.PartitionMetaMap, error) {
	m := mapper.PartitionMetaMap{}
	for _, p := range s.Partitions {
		if _, ok := m[p.Topic]; !ok {
			m[p.Topic] = map[int]*mapper.PartitionMeta{}
		}
		m[p.Topic][p.Partition] = &mapper.PartitionMeta{
			Size: p.Size,
		}

	}
	return m, nil
}
