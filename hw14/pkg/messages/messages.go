package messages

import (
	"strconv"
	"time"
)

type Message struct {
	ID   int
	Time time.Time
	Text string
}

func (m *Message) ToString() string {
	return strconv.Itoa(m.ID) + ":" + m.Time.Format("02.01.2006 15:04:05") + ` - "` + m.Text + `"`
}
