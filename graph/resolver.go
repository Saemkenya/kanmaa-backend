package graph

import (
	"kanmaa-backend/graph/model"
	"sync"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Rooms map[string]*model.Chatroom
	mu    sync.Mutex // nolint: structcheck
	// messages []*model.Message
}
