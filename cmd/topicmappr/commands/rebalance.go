package commands

import (
	"fmt"
	"os"
	"sort"

	"github.com/DataDog/kafka-kit/v3/kafkazk"

	"github.com/spf13/cobra"
)

var rebalanceCmd = &cobra.Command{
	Use:   "rebalance",
	Short: "Rebalance partition allotments among a set of topics and brokers",
	Long:  `Rebalance partition allotments among a set of topics and brokers`,
	Run:   rebalance,
}

func init() {
	rootCmd.AddCommand(rebalanceCmd)

	rebalanceCmd.Flags().String("topics", "", "Rebuild topics (comma delim. list) by lookup in ZooKeeper")
	rebalanceCmd.Flags().String("topics-exclude", "", "Exclude topics")
	rebalanceCmd.Flags().String("out-path", "", "Path to write output map files to")
	rebalanceCmd.Flags().String("out-file", "", "If defined, write a combined map of all topics to a file")
	rebalanceCmd.Flags().String("brokers", "", "Broker list to scope all partition placements to ('-1' for all currently mapped brokers, '-2' for all brokers in cluster)")
	rebalanceCmd.Flags().Float64("storage-threshold", 0.20, "Percent below the harmonic mean storage free to target for partition offload (0 targets a brokers)")
	rebalanceCmd.Flags().Float64("storage-threshold-gb", 0.00, "Storage free in gigabytes to target for partition offload (those below the specified value); 0 [default] defers target selection to --storage-threshold")
	rebalanceCmd.Flags().Float64("tolerance", 0.0, "Percent distance from the mean storage free to limit storage scheduling (0 performs automatic tolerance selection)")
	rebalanceCmd.Flags().Int("partition-limit", 30, "Limit the number of top partitions by size eligible for relocation per broker")
	rebalanceCmd.Flags().Int("partition-size-threshold", 512, "Size in megabytes where partitions below this value will not be moved in a rebalance")
	rebalanceCmd.Flags().Bool("locality-scoped", false, "Ensure that all partition movements are scoped by rack.id")
	rebalanceCmd.Flags().Bool("verbose", false, "Verbose output")
	rebalanceCmd.Flags().String("zk-metrics-prefix", "topicmappr", "ZooKeeper namespace prefix for Kafka metrics")
	rebalanceCmd.Flags().Int("metrics-age", 60, "Kafka metrics age tolerance (in minutes)")
	rebalanceCmd.Flags().Bool("optimize-leadership", false, "Rebalance all broker leader/follower ratios")

	// Required.
	rebalanceCmd.MarkFlagRequired("brokers")
	rebalanceCmd.MarkFlagRequired("topics")
}

func rebalance(cmd *cobra.Command, _ []string) {
	bootstrap(cmd)

	// ZooKeeper init.
	zk, err := initZooKeeper(cmd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer zk.Close()

	// Get broker and partition metadata.
	checkMetaAge(cmd, zk)
	brokerMeta := getBrokerMeta(cmd, zk, true)
	partitionMeta := getPartitionMeta(cmd, zk)

	// Get the current partition map.
	partitionMapIn, err := kafkazk.PartitionMapFromZK(Config.topics, zk)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Exclude any topics that are pending deletion.
	pending := stripPendingDeletes(partitionMapIn, zk)

	// Exclude any explicit exclusions.
	excluded := removeTopics(partitionMapIn, Config.topicsExclude)

	// Print topics matched to input params.
	printTopics(partitionMapIn)

	// Print if any topics were excluded due to pending deletion.
	printExcludedTopics(pending, excluded)

	// Get a broker map.
	brokersIn := kafkazk.BrokerMapFromPartitionMap(partitionMapIn, brokerMeta, false)

	// Validate all broker params, get a copy of the broker IDs targeted for
	// partition offloading.
	offloadTargets := validateBrokersForRebalance(cmd, brokersIn, brokerMeta)

	// Sort offloadTargets by storage free ascending.
	sort.Sort(offloadTargetsBySize{t: offloadTargets, bm: brokersIn})

	partitionLimit, _ := cmd.Flags().GetInt("partition-limit")
	partitionSizeThreshold, _ := cmd.Flags().GetInt("partition-size-threshold")
	tolerance, _ := cmd.Flags().GetFloat64("tolerance")
	localityScoped, _ := cmd.Flags().GetBool("locality-scoped")
	verbose, _ := cmd.Flags().GetBool("verbose")

	params := computeReassignmentBundlesParams{
		offloadTargets:         offloadTargets,
		tolerance:              tolerance,
		partitionMap:           partitionMapIn,
		partitionMeta:          partitionMeta,
		brokerMap:              brokersIn,
		partitionLimit:         partitionLimit,
		partitionSizeThreshold: partitionSizeThreshold,
		localityScoped:         localityScoped,
		verbose:                verbose,
	}

	results := computeReassignmentBundles(params)

	// Merge all results into a slice.
	resultsByRange := []reassignmentBundle{}
	for r := range results {
		resultsByRange = append(resultsByRange, r)
	}

	// Sort the rebalance results by range ascending.
	sort.Slice(resultsByRange, func(i, j int) bool {
		switch {
		case resultsByRange[i].storageRange < resultsByRange[j].storageRange:
			return true
		case resultsByRange[i].storageRange > resultsByRange[j].storageRange:
			return false
		}

		return resultsByRange[i].stdDev < resultsByRange[j].stdDev
	})

	// Chose the results with the lowest range.
	m := resultsByRange[0]
	partitionMapOut, brokersOut, relos := m.partitionMap, m.brokers, m.relocations

	// Print parameters used for rebalance decisions.
	printReassignmentParams(cmd, resultsByRange, brokersIn, m.tolerance)

	// Optimize leaders.
	if t, _ := cmd.Flags().GetBool("optimize-leadership"); t {
		partitionMapOut.OptimizeLeaderFollower()
	}

	// Print planned relocations.
	printPlannedRelocations(offloadTargets, relos, partitionMeta)

	// Print map change results.
	printMapChanges(partitionMapIn, partitionMapOut)

	// Print broker assignment statistics.
	errs := printBrokerAssignmentStats(cmd, partitionMapIn, partitionMapOut, brokersIn, brokersOut)

	// Handle errors that are possible to be overridden by the user (aka 'WARN'
	// in topicmappr console output).
	handleOverridableErrs(cmd, errs)

	// Ignore no-ops; rebalances will naturally have a high percentage of these.
	partitionMapIn, partitionMapOut = skipReassignmentNoOps(partitionMapIn, partitionMapOut)

	// Write maps.
	writeMaps(cmd, partitionMapOut, nil)
}
