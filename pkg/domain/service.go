package domain

import (
	"context"

	"github.com/opencars/space-trouble/pkg/domain/command"
	"github.com/opencars/space-trouble/pkg/domain/model"
	"github.com/opencars/space-trouble/pkg/domain/query"
)

type BookingRepository interface {
	Create(context.Context, *model.Booking) error
	Delete(context.Context, string) error
	List(context.Context, *query.List) ([]model.Booking, error)
}

type SpacexService interface {
	FindByLaunchpadID(ctx context.Context, launchpadID string) error
}

type CustomerService interface {
	CreateBooking(context.Context, *command.CreateBooking) (*model.Booking, error)
	DeleteBooking(context.Context, *command.DeleteBooking) error
	ListBookings(context.Context, *query.List) ([]model.Booking, error)
}
