package model

type Chatroom struct {
	Name      string
	Messages  []Message
	Observers map[string]struct {
		Username string
		Message  chan *Message
	}
}
