package service

import (
	"context"

	"github.com/opencars/space-trouble/pkg/domain"
	"github.com/opencars/space-trouble/pkg/domain/command"
	"github.com/opencars/space-trouble/pkg/domain/model"
	"github.com/opencars/space-trouble/pkg/domain/query"
)

type CustomerService struct {
	repo   domain.BookingRepository
	spacex domain.SpacexService
}

func NewCustomerService(repo domain.BookingRepository) *CustomerService {
	return &CustomerService{
		repo: repo,
	}
}

func (s *CustomerService) CreateBooking(ctx context.Context, c *command.CreateBooking) (*model.Booking, error) {
	if err := command.Process(c); err != nil {
		return nil, err
	}

	booking := model.Booking{
		FirstName:     c.FirstName,
		LastName:      c.LastName,
		Gender:        c.Gender,
		BirthDate:     c.BirthDate,
		LaunchpadID:   c.LaunchpadID,
		DestinationID: c.DestinationID,
		LaunchDate:    c.LaunchDate,
	}

	// TODO: Finish logic.
	s.spacex.FindByLaunchpadID(ctx, booking.LaunchpadID)

	if err := s.repo.Create(ctx, &booking); err != nil {
		return nil, err
	}

	return &booking, nil
}

func (s *CustomerService) DeleteBooking(ctx context.Context, c *command.DeleteBooking) error {
	if err := command.Process(c); err != nil {
		return err
	}

	if err := s.repo.Delete(ctx, c.ID); err != nil {
		return err
	}

	return nil
}

func (s *CustomerService) ListBookings(ctx context.Context, q *query.List) ([]model.Booking, error) {
	if err := query.Process(q); err != nil {
		return nil, err
	}

	bookings, err := s.repo.List(ctx, q)
	if err != nil {
		return nil, err
	}

	return bookings, nil
}
