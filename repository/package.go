package repository

import (
	"github.com/behavioral-ai/core/host"
	"github.com/behavioral-ai/core/messaging"
)

var (
	msg = host.NewSyncMap[string, *messaging.Message]()
)

// LoadMessage - get a message
func LoadMessage(name string) *messaging.Message {
	return msg.Load(name)
}

// StoreMessage - store a message
func StoreMessage(m *messaging.Message) {
	msg.Store(m.Name, m)
}

/*
// ModifyMessage - modify a message
func ModifyMessage(m *messaging.Message) {
	msg.modify(m)
}

// DeleteMessage - delete a message
func DeleteMessage(name string) {
	msg.delete(name)
}


*/
