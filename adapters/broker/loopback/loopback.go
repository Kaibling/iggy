package loopback

import (
	"context"
	"fmt"

	"github.com/kaibling/apiforge/logging"
	"github.com/kaibling/iggy/entity"
	"github.com/kaibling/iggy/pkg/utility"
)

func NewLoopback(ctx context.Context, c chan []byte, l logging.LogWriter) *Loopback {
	return &Loopback{ctx, c, l}
}

type Loopback struct {
	ctx context.Context
	c   chan []byte
	l   logging.LogWriter
}

func (l *Loopback) Subscribe(channelName string) error {
	for {
		select {
		case <-l.ctx.Done():
			l.l.LogLine("shuting down worker")
			return l.ctx.Err()
		case newMessage := <-l.c:
			t, err := utility.DecodeToStruct[entity.Task](newMessage)
			if err != nil {
				// log
				fmt.Println(err.Error())
			}
			fmt.Println(utility.Pretty(t))

		}

	}

}
func (l *Loopback) Publish(channelName string, message []byte) error {
	l.c <- message
	return nil
}
