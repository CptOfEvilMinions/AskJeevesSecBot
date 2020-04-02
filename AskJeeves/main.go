package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/CptOfEvilMinions/AskJeevesSecBot/pkg/config"

	"github.com/CptOfEvilMinions/AskJeevesSecBot/pkg/brokers"
	"github.com/CptOfEvilMinions/AskJeevesSecBot/pkg/database"
	"github.com/CptOfEvilMinions/AskJeevesSecBot/pkg/geoip"
	"github.com/CptOfEvilMinions/AskJeevesSecBot/pkg/hash"
	"github.com/CptOfEvilMinions/AskJeevesSecBot/pkg/slack"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

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

	// Run MySQL clean up
	// Init background task
	go database.DeleteOldEntries(mysqlConnector, cfg)

	// Init Kafka Consumer
	kafkaConsumer, err := brokers.ConsumerInit(cfg)
	if err != nil {
		fmt.Println(err.Error())
		log.Fatalln(err)
	}

	// Init GeoIP database reader
	geoIPreader, err := geoip.InitGeoIPReader(cfg)
	if err != nil {
		fmt.Println(err.Error())
		log.Fatalln(err)
	}

	// Init Slack API
	slackConnector := slack.InitConnector(cfg)
	fmt.Println(slackConnector)

	// Iterate through all messages in topic
	run := true
	for run {
		// PollInterval * 1000 for seconds
		ev := kafkaConsumer.Poll(cfg.Kafka.PollInterval * 1000)
		switch e := ev.(type) {
		case *kafka.Message:
			// Print message
			fmt.Printf("\n\n%% Message on %s:\n%s\n", e.TopicPartition, string(e.Value))

			// Extract JSON string to struct
			var vpnEntry brokers.VPNdata
			json.Unmarshal([]byte(e.Value), &vpnEntry)

			// Set Location
			location, isoCode, err := geoip.IPaddrLocationLookup(geoIPreader, vpnEntry.SrcIP)
			//vpnEntry.Location, err = geoip.IPaddrLocationLookup(vpnEntry.SrcIP)
			if err != nil {
				fmt.Println(err.Error())
				log.Fatalln(err)
			}

			// Set VPN hash
			vpnHash := hash.VpnHash(vpnEntry.Username, vpnEntry.SrcIP, isoCode)
			fmt.Printf("VPN hash: %s\n", vpnHash)

			// Query database for VPN hash
			// If VpnHash does not exist add it to database, query location, and send slack message to user
			if database.QueryDoesVpnHashExist(mysqlConnector, vpnHash) == false {
				// Add entry to database
				database.AddVpnUserEntry(mysqlConnector, vpnEntry.Username, vpnHash, vpnEntry.SrcIP, isoCode, location)

				// Send Slack message to user
				err := slack.SendUserMessage(cfg, slackConnector, vpnEntry, location, vpnHash)
				if err != nil {
					fmt.Println(err.Error())
					log.Fatalln(err)
				}

			}

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
}