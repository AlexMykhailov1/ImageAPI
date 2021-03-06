package rabbit

import (
	"github.com/AlexMykhailov1/ImageAPI/config"
	"github.com/AlexMykhailov1/ImageAPI/pkg/optimizer"
	"github.com/streadway/amqp"
	"log"
)

// ConsumeImgID consumes image names from a queue and sends it for the optimization
func (r *Rabbit) ConsumeImgID(cfg *config.Config) error {
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

	// Receive messages from channel
	msgs, err := ch.Consume(
		cfg.RB.ImgQueue, // queue
		"",              // consumer
		true,            // auto-ack
		false,           // exclusive
		false,           // no-local
		false,           // no-wait
		nil,             // args
	)
	if err != nil {
		return err
	}

	forever := make(chan bool)

	go func() {
		for data := range msgs {
			log.Printf("Received a message: %s", data.Body)
			msg := string(data.Body)
			go func() {
				err = optimizer.SaveLessQuality(cfg, msg)
				if err != nil {
					log.Printf("Error processing image:%v", err.Error())
					return
				}
			}()
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
	return nil
}
