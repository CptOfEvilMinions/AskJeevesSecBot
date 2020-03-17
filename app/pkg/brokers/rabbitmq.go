package brokers

import (
	"github.com/streadway/amqp"
	"fmt"
	"strconv"
	"github.com/CptOfEvilMinions/AskJeevesSecBot/pkg/config"
)

// InitRabbitConnection input:
// InitRabbitConnection output: 
func InitRabbitConnection(cfg *config.Config) (amqp.Queue, error){
	fmt.Println("=============== RabbitMQ ===============")
	fmt.Println(cfg.Kafka.Hostname)
	fmt.Println(cfg.Kafka.GroupId)
	fmt.Println(cfg.Kafka.Offset)
	fmt.Println(cfg.Kafka.Port)
	fmt.Println(cfg.Kafka.ConsumerTopic)

	rabbitMqServer := fmt.Sprintf("amqp://%s:%s@%s:%s/", cfg.RabbitMQ.Username, cfg.RabbitMQ.Password, cfg.RabbitMQ.Hostname, strconv.Itoa(cfg.RabbitMQ.Port))
	
	// Connect to the rabbitMQ instance
	connection, err := amqp.Dial(rabbitMqServer)
	if err != nil {
        panic("could not establish connection with RabbitMQ:" + err.Error())
	}
	defer connection.Close()
	
	// Create a channel from the connection. We'll use channels to access the data in the queue rather than the connection itself.
	channel, err := connection.Channel()
	if err != nil {
        panic("could not open RabbitMQ channel:" + err.Error())
	}
	defer channel.Close()

	//
	queue, err := channel.QueueDeclare(
		cfg.RabbitMQ.QueueName, // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	  )

	return queue, nil
}	


