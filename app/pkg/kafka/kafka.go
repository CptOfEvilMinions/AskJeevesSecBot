package kafka

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// KafkaConsumerInit input:
// KafkaConsumerInit output:
func ConsumerInit() (*kafka.Consumer, error) {

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		return nil, err
	}

	return c, err
}

// KafkaProducerInit input:
// KafkaProducerInit output:
func ProducerInit() (*kafka.Producer, error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost"
		"client.id": socket.gethostname(),
		"default.topic.config": kafka.ConfigMap{"acks": "all"}
	})
	if err != nil {
		return nil, err
	}

	defer p.Close()

	return p, err
}


// ProduceMessagesToTopic input:
// ProduceMessagesToTopic output:
// https://docs.confluent.io/current/clients/go.html
func ProduceMessagesToTopic( p *kafka.Producer, message string, topic, string) (bool, error) {
	// Push message to Kafka
	delivery_chan := make(chan kafka.Event, 10000)
	err := p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: topic, Partition: kafka.PartitionAny},
		Value: []byte(message)},
		delivery_chan,
	)

	e := <-delivery_chan
	m := e.(*kafka.Message)


	if m.TopicPartition.Error != nil 
		fmt.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
		return false, err
	} else {
		fmt.Printf("Delivered message to topic %s [%d] at offset %v\n", *m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
		
	}
	close(delivery_chan)

	return true, nil
	
} 