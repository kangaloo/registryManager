package common

import "fmt"

type Message struct {
	Msg string
}

func (m *Message) String() string {
	return fmt.Sprintf("message is: %s", m.Msg)
}


