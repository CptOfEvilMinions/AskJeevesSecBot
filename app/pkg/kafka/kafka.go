package kafka

import (
	"fmt"
	"os"
	"strconv"

	"github.com/CptOfEvilMinions/AskJeevesSecBot/pkg/config"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// KafkaConsumerInit input:
// KafkaConsumerInit output:
func ConsumerInit(cfg *config.Config) (*kafka.Consumer, error) {

	fmt.Println(cfg.Kafka.Hostname)
	fmt.Println(cfg.Kafka.GroupId)
	fmt.Println(cfg.Kafka.Offset)
	fmt.Println(cfg.Kafka.Port)
	fmt.Println(cfg.Kafka.ConsumerTopic)

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

// KafkaProducerInit input:
// KafkaProducerInit output:
func ProducerInit(cfg *config.Config) (*kafka.Producer, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return nil, err
	}

	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers":    "localhost",
		"client.id":            hostname,
		"default.topic.config": kafka.ConfigMap{"acks": "all"}})
	if err != nil {
		return nil, err
	}

	defer p.Close()

	return p, err
}

// ProduceMessagesToTopic input:
// ProduceMessagesToTopic output:
// https://docs.confluent.io/current/clients/go.html
// func ProduceMessagesToTopic( p *kafka.Producer, message string, topic, string) (bool, error) {
// 	// Push message to Kafka
// 	delivery_chan := make(chan kafka.Event, 10000)
// 	err := p.Produce(&kafka.Message{
// 		TopicPartition: kafka.TopicPartition{Topic: topic, Partition: kafka.PartitionAny},
// 		Value: []byte(message)},
// 		delivery_chan,
// 	)

// 	e := <-delivery_chan
// 	m := e.(*kafka.Message)

// 	if m.TopicPartition.Error != nil
// 		fmt.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
// 		return false, err
// 	} else {
// 		fmt.Printf("Delivered message to topic %s [%d] at offset %v\n", *m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)

// 	}
// 	close(delivery_chan)

// 	return true, nil

// }
