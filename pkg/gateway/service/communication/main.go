package communication

import (
	"context"
)

type MediaSender interface {
	SendGoodToChannel(ctx context.Context, good Good)
	SendGoodToChat(chatID int64, good Good)
}

type Commutator struct {
	mediaSenders []MediaSender
}

func NewCommutator(senders ...MediaSender) *Commutator {
	return &Commutator{
		mediaSenders: senders,
	}
}

func (c *Commutator) SendGood(ctx context.Context, good Good) { // add wg
	for _, sender := range c.mediaSenders {
		sender.SendGoodToChannel(ctx, good)
	}
}
