package consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"karhub/internal/entity"

	"github.com/gofiber/fiber/v2/log"
)

type reprocess struct {
	rabbit Rabbit
	rep    Reprocess
}

func (c *reprocess) Consume() error {

	msgs, err := c.rabbit.Consume("reprocess-queue")
	if err != nil {
		log.Error(fmt.Sprintf("error to consume: %s", err))
		return err
	}

	for msg := range msgs {

		go func() {
			defer msg.Ack(true)
			var repEntity entity.Reprocess

			err := json.Unmarshal(msg.Body, &repEntity)
			if err != nil {
				log.Error(fmt.Sprintf("error to unmarshal reprocess: %s", err))
				return
			}

			err = c.rep.Reprocess(context.Background(), repEntity)
			if err != nil {
				log.Error(fmt.Sprintf("error to reprocess: %s", err))
				return
			}

		}()

	}

	return nil
}

func NewReprocess(rabbit Rabbit, rep Reprocess) *reprocess {
	return &reprocess{rabbit: rabbit, rep: rep}
}
