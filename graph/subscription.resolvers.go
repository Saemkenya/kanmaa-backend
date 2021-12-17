package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"kanmaa-backend/graph/generated"
	"kanmaa-backend/graph/model"
)

func (r *subscriptionResolver) PaymentMade(ctx context.Context) (<-chan *model.Payment, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) LessonAttended(ctx context.Context) (<-chan *model.Attendance, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) ExamScored(ctx context.Context) (<-chan *model.Score, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) NuLessonIn5(ctx context.Context) (<-chan *model.Lesson, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) StudentHasArrived(ctx context.Context) (<-chan *model.StudentinStudentout, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) StudentHasDeparted(ctx context.Context) (<-chan *model.StudentinStudentout, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) PaymentConfirmed(ctx context.Context) (<-chan *model.Status, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) BookAdded(ctx context.Context) (<-chan *model.Book, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) BookLentout(ctx context.Context) (<-chan *model.Lend, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) BookReceivedBack(ctx context.Context) (<-chan *model.Lend, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) ProductAdded(ctx context.Context) (<-chan *model.Product, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) ProductDispensed(ctx context.Context) (<-chan *model.DispensedProduct, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) DiaryAdded(ctx context.Context) (<-chan *model.Diary, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) DiarySeen(ctx context.Context) (<-chan *model.Diary, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) WorkerHasArrived(ctx context.Context) (<-chan *model.CheckinCheckout, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) WorkerHasDeparted(ctx context.Context) (<-chan *model.CheckinCheckout, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) VisitorCheckedin(ctx context.Context) (<-chan *model.Visitor, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) VisitorCheckedout(ctx context.Context) (<-chan *model.Visitor, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) AnnounceAdded(ctx context.Context) (<-chan *model.Announcement, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) NewsletterAdded(ctx context.Context) (<-chan *model.Newsletter, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) LessonAdded(ctx context.Context) (<-chan *model.Lesson, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) StudentAdded(ctx context.Context) (<-chan *model.Student, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) UserAdded(ctx context.Context) (<-chan *model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) ItemArchived(ctx context.Context) (<-chan *model.Archive, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) AttendanceConfirmed(ctx context.Context) (<-chan *model.Attendance, error) {
	panic(fmt.Errorf("not implemented"))
}

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type subscriptionResolver struct{ *Resolver }
