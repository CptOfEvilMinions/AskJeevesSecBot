package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"pkg/geoip"
	"pkg/hash"
	"pkg/kafka"
)

type VPNdata struct {
	Username string  `json:"username"`
	IPaddr   string  `json:"ipaddr"`
	Location *string `json:"location"`
	Hash     *string `json:"Hash"`
}

func run() {
	kafkaConsumer := kafka.ConsumerInit()
	kafkaProducer := kafka.ProducerInit()

	kafkaConsumer.SubscribeTopics([]string{"myTopic", "^aRegex.*[Tt]opic"}, nil)

	// Iterate throuugh all messages in topic
	for run == true {
		ev := kafkaConsumer.Poll(0)
		switch e := ev.(type) {
		case *kafka.Message:
			// Print message
			fmt.Printf("%% Message on %s:\n%s\n", e.TopicPartition, string(e.Value))

			// Extract JSON data
			vpn := new(VPNdata)
			if err := json.NewDecoder(resp.Body).Decode(vpn); err != nil {
				log.Fatalln(err)
			}
			if len(t.Error) > 0 {
				log.Fatalln(t.Error)
			}

			// Set Location
			vpn.Location = geoip.IPaddrLocationLookup(ipAddr)

			// Set VPN hash
			vpn.Hash = hash.VpnHash(vpn.username, vpn.IPaddr, location)

			vpnJSON, err := json.Marshal(vpn)
			if err != nil {
				log.Fatalln(err.Error)
			}

			// Enrich VPN data - push to Kafka
			ProduceMessagesToTopic(kafkaProducer, vpnJson, "vpn-enriched")

		case kafka.PartitionEOF:
			fmt.Printf("%% Reached %v\n", e)
		case kafka.Error:
			fmt.Fprintf(os.Stderr, "%% Error: %v\n", e)
			run = false
		default:
			fmt.Printf("Ignored %v\n", e)
		}
	}

	// Close connection to Kafka
	kafkaConsumer.Close()

	// Wait for message deliveries before shutting down
	kafkaProducer.Flush(15 * 1000)
}
