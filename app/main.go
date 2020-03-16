package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/CptOfEvilMinions/AskJeevesSecBot/pkg/config"
	"github.com/CptOfEvilMinions/AskJeevesSecBot/pkg/database"
	"github.com/CptOfEvilMinions/AskJeevesSecBot/pkg/geoip"
	"github.com/CptOfEvilMinions/AskJeevesSecBot/pkg/hash"
	myKafka "github.com/CptOfEvilMinions/AskJeevesSecBot/pkg/kafka"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type VPNdata struct {
	Timestamp     time.Time `json:"@timestamp"`
	Host          string    `json:"host"`
	SyslogProgram string    `json:"syslog_program"`
	Message       string    `json:"message"`
	Username      string    `json:"username"`
	SrcIP         string    `json:"src_ip"`
	Location      uint      `json:"location"`
	VpnHash       string    `json:"vpn_hash"`
}

func main() {
	// Generate our config based on the config supplied
	cfg, err := config.NewConfig("config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	// Init database connection
	mysqlConnector, err := database.InitMySQLConnector(cfg)
	if err != nil {
		fmt.Println(err.Error())
		log.Fatalln(err)
	}

	// Init Kafka Consumer
	kafkaConsumer, err := myKafka.ConsumerInit(cfg)
	if err != nil {
		fmt.Println(err.Error())
		log.Fatalln(err)
	}

	// Init Kafka Producer
	// kafkaProducer, err := myKafka.ProducerInit(cfg)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	log.Fatalln(err)
	// }

	// Iterate throuugh all messages in topic
	run := true
	for run {
		// PollInterval * 1000 for seconds
		ev := kafkaConsumer.Poll(cfg.Kafka.PollInterval * 1000)
		switch e := ev.(type) {
		case *kafka.Message:
			// Print message
			fmt.Printf("\n\n%% Message on %s:\n%s\n", e.TopicPartition, string(e.Value))

			// Extract JSON string to struct
			var vpnEntry VPNdata
			json.Unmarshal([]byte(e.Value), &vpnEntry)

			// DEBUG
			//fmt.Printf("json object:::: %v\n", vpnEntry)

			// Set Location
			vpnEntry.Location, err = geoip.IPaddrLocationLookup(vpnEntry.SrcIP)
			if err != nil {
				fmt.Println(err.Error())
				log.Fatalln(err)
			}

			// Set VPN hash
			vpnEntry.VpnHash = hash.VpnHash(vpnEntry.Username, vpnEntry.SrcIP, vpnEntry.Location)
			fmt.Printf("VPN hash: %s\n", vpnEntry.VpnHash)

			// Query database for VPN hash
			result := database.QueryDoesVpnHashExist(mysqlConnector, vpnEntry.VpnHash)

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
	}

	// Close connection to Kafka
	kafkaConsumer.Close()

	// Wait for message deliveries before shutting down
	//kafkaProducer.Flush(15 * 1000)
}
