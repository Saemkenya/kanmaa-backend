package graph

// // This file will be automatically regenerated based on the schema, any resolver implementations
// // will be copied through when generating and any unknown code will be moved to the end.

// import (
// 	"context"
// 	"fmt"
// 	"kanmaa-backend/graph/generated"
// 	"kanmaa-backend/graph/model"
// 	"math/rand"
// 	"time"

// 	"github.com/99designs/gqlgen/graphql"
// )

// func (r *mutationResolver) Post(ctx context.Context, text string, username string, roomName string) (*model.Message, error) {
// 	fmt.Println("posting new message")
// 	// panic(fmt.Errorf("not implemented"))
// 	r.mu.Lock()
// 	room := r.Rooms[roomName]
// 	if room == nil {
// 		room = &model.Chatroom{
// 			Name: roomName,
// 			Observers: map[string]struct {
// 				Username string
// 				Message  chan *model.Message
// 			}{},
// 		}
// 		r.Rooms[roomName] = room
// 	}
// 	r.mu.Unlock()

// 	message := model.Message{
// 		ID:        randString(8),
// 		CreatedAt: time.Now(),
// 		Text:      text,
// 		CreatedBy: username,
// 	}

// 	room.Messages = append(room.Messages, message)
// 	r.mu.Lock()
// 	for _, observer := range room.Observers {
// 		if observer.Username == "" || observer.Username == message.CreatedBy {
// 			observer.Message <- &message
// 		}
// 	}
// 	r.mu.Unlock()
// 	return &message, nil
// }

// func (r *queryResolver) Room(ctx context.Context, name string) (*model.Chatroom, error) {
// 	fmt.Println("Creating/Appending to room")
// 	// panic(fmt.Errorf("not implemented"))
// 	r.mu.Lock()
// 	room := r.Rooms[name]
// 	if room == nil {
// 		room = &model.Chatroom{
// 			Name: name,
// 			Observers: map[string]struct {
// 				Username string
// 				Message  chan *model.Message
// 			}{},
// 		}
// 		r.Rooms[name] = room
// 	}
// 	r.mu.Unlock()

// 	return room, nil
// }

// func (r *subscriptionResolver) MessageAdded(ctx context.Context, roomName string) (<-chan *model.Message, error) {
// 	fmt.Println("Message added to subscriber channels")
// 	// panic(fmt.Errorf("not implemented"))
// 	r.mu.Lock()
// 	room := r.Rooms[roomName]
// 	if room == nil {
// 		room = &model.Chatroom{
// 			Name: roomName,
// 			Observers: map[string]struct {
// 				Username string
// 				Message  chan *model.Message
// 			}{},
// 		}
// 		r.Rooms[roomName] = room
// 	}
// 	r.mu.Unlock()

// 	id := randString(8)
// 	events := make(chan *model.Message, 1)

// 	go func() {
// 		<-ctx.Done()
// 		r.mu.Lock()
// 		delete(room.Observers, id)
// 		r.mu.Unlock()
// 	}()

// 	r.mu.Lock()
// 	room.Observers[id] = struct {
// 		Username string
// 		Message  chan *model.Message
// 	}{Username: getUsername(ctx), Message: events}
// 	r.mu.Unlock()

// 	return events, nil
// }

// // Mutation returns generated.MutationResolver implementation.
// func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// // Query returns generated.QueryResolver implementation.
// func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// // Subscription returns generated.SubscriptionResolver implementation.
// func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

// type mutationResolver struct{ *Resolver }
// type queryResolver struct{ *Resolver }
// type subscriptionResolver struct{ *Resolver }

// // !!! WARNING !!!
// // The code below was going to be deleted when updating resolvers. It has been copied here so you have
// // one last chance to move it out of harms way if you want. There are two reasons this happens:
// //  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
// //    it when you're done.
// //  - You have helper methods in this file. Move them out to keep these resolver files clean.
// type okey string

// func New() generated.Config {
// 	fmt.Println("generating new config")
// 	return generated.Config{
// 		Resolvers: &Resolver{
// 			Rooms: map[string]*model.Chatroom{},
// 		},
// 		Directives: generated.DirectiveRoot{
// 			User: func(
// 				ctx context.Context,
// 				obj interface{},
// 				next graphql.Resolver,
// 				username string,
// 			) (res interface{}, err error) {
// 				return next(context.WithValue(ctx, okey("username"), username))
// 			},
// 		},
// 	}
// }
// func getUsername(ctx context.Context) string {
// 	fmt.Println("getting user name")
// 	if username, ok := ctx.Value(okey("username")).(string); ok {
// 		return username
// 	}
// 	return ""
// }

// var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// func randString(n int) string {
// 	fmt.Println("generating random string")
// 	b := make([]rune, n)
// 	for i := range b {
// 		b[i] = letterRunes[rand.Intn(len(letterRunes))]
// 	}
// 	return string(b)
// }
