package commands

import (
	"fmt"
	"os"
	"strings"

	"github.com/jamiealquiza/envy"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "topicmappr",
}

var rdkafka = []string{}

// Execute rootCmd.
func Execute() {
	envy.ParseCobra(rootCmd, envy.CobraConfig{Prefix: "TOPICMAPPR", Persistent: true, Recursive: false})

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().String("kafka-addr", "localhost:9092", "Kafka bootstrap address")
	rootCmd.PersistentFlags().String("zk-addr", "localhost:2181", "ZooKeeper connect string")
	rootCmd.PersistentFlags().String("zk-prefix", "", "ZooKeeper prefix (if Kafka is configured with a chroot path prefix)")
	rootCmd.PersistentFlags().String("zk-metrics-prefix", "topicmappr", "ZooKeeper namespace prefix for Kafka metrics")
	rootCmd.PersistentFlags().Bool("ignore-warns", false, "Produce a map even if warnings are encountered")
	rootCmd.PersistentFlags().StringSliceVar(&rdkafka, "rdkafka", []string{}, "Set librdkafka configuration property (see https://github.com/edenhill/librdkafka/blob/master/CONFIGURATION.md)")
}

func rdkafkaOpts() map[string]string {
	opts := map[string]string{}

	for _, opt := range rdkafka {
		parts := strings.Split(opt, "=")
		opts[parts[0]] = parts[1]
	}

	return opts
}
