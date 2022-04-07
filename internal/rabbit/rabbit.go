package rabbit

import (
	"fmt"
	"github.com/AlexMykhailov1/ImageAPI/config"
	"github.com/streadway/amqp"
	"log"
)

// Rabbit struct stores connection to RabbitMQ
type Rabbit struct {
	*amqp.Connection
}

// NewRabbit creates new RabbitMQ connection using env variables from config
// and returns it in Rabbit struct
func NewRabbit(cfg *config.Config) (*Rabbit, error) {
	// Sprint url
	url := fmt.Sprintf("amqp://%s:%s@%s:%s/",
		cfg.RB.User,
		cfg.RB.Password,
		cfg.RB.Host,
		cfg.RB.Port)

	// Connect
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}
	return &Rabbit{conn}, nil
}

// SetUpQueues creates all RabbitMQ queues
func (r *Rabbit) SetUpQueues(cfg *config.Config) error {
	// Open channel from existing connection
	ch, err := r.Connection.Channel()
	if err != nil {
		return err
	}
	// Defer close the channel
	defer func(ch *amqp.Channel) {
		err = ch.Close()
		if err != nil {
			log.Printf("Error closing RabbitMQ channel:%v", err.Error())
		}
	}(ch)

	// Declare a queue for images
	_, err = ch.QueueDeclare(
		cfg.RB.ImgQueue, // name
		false,           // durable
		false,           // delete when unused
		false,           // exclusive
		false,           // no-wait
		nil,             // arguments
	)
	if err != nil {
		return err
	}
	return nil
}
