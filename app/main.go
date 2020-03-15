package main

import (
	"fmt"
	"log"
	"os"

	"github.com/CptOfEvilMinions/AskJeevesSecBot/pkg/config"
	myKafka "github.com/CptOfEvilMinions/AskJeevesSecBot/pkg/kafka"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type VPNdata struct {
	Username string  `json:"username"`
	IPaddr   string  `json:"ipaddr"`
	Location *string `json:"location"`
	Hash     *string `json:"Hash"`
}

func main() {
	// Generate our config based on the config supplied
	cfg, err := config.NewConfig("config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	// Init Kafka Consumer
	kafkaConsumer, err := myKafka.ConsumerInit(cfg)
	if err != nil {
		fmt.Println(err.Error())
		log.Fatalln(err)
	}

	// Init Kafka Producer
	kafkaProducer, err := myKafka.ProducerInit(cfg)
	if err != nil {
		fmt.Println(err.Error())
		log.Fatalln(err)
	}

	// Iterate throuugh all messages in topic
	run := true
	for run {
		ev := kafkaConsumer.Poll(cfg.Kafka.PollInterval * 1000)
		switch e := ev.(type) {
		case *kafka.Message:
			// Print message
			fmt.Printf("%% Message on %s:\n%s\n", e.TopicPartition, string(e.Value))

			// Extract JSON data
			// vpn := new(VPNdata)

			// if err := json.NewDecoder(vpn).Decode(vpn); err != nil {
			// 	log.Fatalln(err)
			// }
			// if len(t.Error) > 0 {
			// 	log.Fatalln(t.Error)
			// }

			// Set Location
			//vpn.Location = geoip.IPaddrLocationLookup(ipAddr)

			// Set VPN hash
			// vpn.Hash = hash.VpnHash(vpn.Username, vpn.IPaddr, vpn.Location)

			// vpnJSON, err := json.Marshal(vpn)
			// if err != nil {
			// 	log.Fatalln(err.Error)
			// }

			// Enrich VPN data - push to Kafka
			//ProduceMessagesToTopic(kafkaProducer, vpnJson, "vpn-enriched")

		case kafka.PartitionEOF:
			fmt.Printf("%% Reached %v\n", e)
		case kafka.Error:
			fmt.Fprintf(os.Stderr, "%% Error: %v\n", e)
			run = false
		default:
			//continue
			fmt.Printf("Ignored %v\n", e)
		}

		// sleep
		//time.Sleep(3 * time.Second)
	}

	// Close connection to Kafka
	kafkaConsumer.Close()

	// Wait for message deliveries before shutting down
	kafkaProducer.Flush(15 * 1000)
}
