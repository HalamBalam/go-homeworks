package messages

import (
	"strconv"
	"time"
)

// Message - сообщение, сохраняемое на сервере.
type Message struct {
	ID   int
	Time time.Time
	Text string
}

// String - возвращает текстовое представление сообщения.
func (m *Message) String() string {
	return strconv.Itoa(m.ID) + ":" + m.Time.Format("02.01.2006 15:04:05") + ` - "` + m.Text + `"`
}
