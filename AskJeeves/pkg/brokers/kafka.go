package brokers

import (
	"fmt"
	"strconv"

	"github.com/CptOfEvilMinions/AskJeevesSecBot/pkg/config"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// KafkaConsumerInit input: config input
// KafkaConsumerInit output: Returns Kafka consumer connector
func ConsumerInit(cfg *config.Config) (*kafka.Consumer, error) {

	fmt.Println("=============== Kafka ===============")
	fmt.Println("Hostname:", cfg.Kafka.Hostname)
	fmt.Println("GroupId:", cfg.Kafka.GroupId)
	fmt.Println("Offset:", cfg.Kafka.Offset)
	fmt.Println("Port:", cfg.Kafka.Port)
	fmt.Println("Consumer topic:", cfg.Kafka.ConsumerTopic)

	kafkaServer := fmt.Sprintf("%s:%s", cfg.Kafka.Hostname, strconv.Itoa(cfg.Kafka.Port))
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": kafkaServer,
		"group.id":          cfg.Kafka.GroupId,
		"auto.offset.reset": cfg.Kafka.Offset,
	})
	if err != nil {
		return nil, err
	}

	// Subscribe to topic
	c.SubscribeTopics([]string{cfg.Kafka.ConsumerTopic}, nil)

	return c, err
}
