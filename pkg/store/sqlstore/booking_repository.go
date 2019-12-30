package sqlstore

import (
	"context"

	"github.com/opencars/space-trouble/pkg/domain/model"
	"github.com/opencars/space-trouble/pkg/domain/query"
)

type BookingRepository struct {
	store *Store
}

func (r *BookingRepository) Create(ctx context.Context, b *model.Booking) error {
	_, err := r.store.db.ExecContext(ctx,
		`INSERT INTO bookings (
			id, first_name, last_name, gender, birth_date, launchpad_id, launch_date, destination_id
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8,
		)`,
		b.ID, b.FirstName, b.LastName, b.Gender, b.BirthDate,
		b.LaunchpadID, b.LaunchDate, b.DestinationID,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *BookingRepository) List(ctx context.Context, q *query.List) ([]model.Booking, error) {
	bookings := make([]model.Booking, 0)

	err := r.store.db.SelectContext(ctx, &bookings,
		`SELECT id, first_name, last_name, gender, birth_date, launchpad_id, launch_date, destination_id
		FROM bookings
		ORDER BY id DESC`,
	)
	if err != nil {
		return nil, err
	}

	return bookings, nil
}

func (r *BookingRepository) Delete(ctx context.Context, id string) error {
	_, err := r.store.db.ExecContext(ctx,
		`DELETE FROM bookings WHERE id = $1`,
		id,
	)
	if err != nil {
		return err
	}

	return nil
}
