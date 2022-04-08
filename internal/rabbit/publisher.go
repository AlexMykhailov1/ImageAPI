package rabbit

import (
	"github.com/AlexMykhailov1/ImageAPI/config"
	"github.com/streadway/amqp"
	"log"
)

// SendImgID sends given image name to the specific queue from the config file
func (r *Rabbit) SendImgID(cfg *config.Config, name string) error {
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

	// Send image id to the queue
	err = ch.Publish(
		"",              // exchange
		cfg.RB.ImgQueue, // routing key
		false,           // mandatory
		false,           // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(name),
		})
	return nil
}
