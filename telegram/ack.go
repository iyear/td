package telegram

import (
	"context"
	"time"

	"go.uber.org/zap"

	"github.com/gotd/td/internal/mt"
)

func (c *Client) ackLoop(ctx context.Context) {
	c.wg.Add(1)
	defer c.wg.Done()

	log := c.log.Named("ack")

	const (
		maxBatchSize      = 20
		forcedSendTimeout = time.Second * 15
	)

	var (
		buf = make([]int64, 0, maxBatchSize)

		// TODO(ernado): remove side-effect.
		timer = time.NewTimer(forcedSendTimeout)
	)

	send := func() {
		defer func() { buf = buf[:0] }()

		if err := c.writeServiceMessage(ctx, &mt.MsgsAck{MsgIds: buf}); err != nil {
			c.log.Error("Failed to ACK", zap.Error(err))
			return
		}

		log.Info("ACK", zap.Int64s("message_ids", buf))
	}

	for {
		select {
		case <-ctx.Done():
			return
		case <-timer.C:
			if len(buf) > 0 {
				send()
			}
		case msgID := <-c.ackSendChan:
			buf = append(buf, msgID)
			if len(buf) == maxBatchSize {
				send()
				timer.Reset(forcedSendTimeout)
			}
		}
	}
}