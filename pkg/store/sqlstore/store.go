package sqlstore

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/opencars/space-trouble/pkg/config"
	"github.com/opencars/space-trouble/pkg/domain"
)

type Store struct {
	db *sqlx.DB

	bookingRepository *BookingRepository
}

func (s *Store) Booking() domain.BookingRepository {
	if s.bookingRepository == nil {
		s.bookingRepository = &BookingRepository{
			store: s,
		}
	}

	return s.bookingRepository
}

func New(settings *config.Database) (*Store, error) {
	info := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=%s password=%s",
		settings.Host,
		settings.Port,
		settings.User,
		settings.Name,
		settings.SSLMode,
		settings.Password,
	)

	db, err := sqlx.Connect("postgres", info)
	if err != nil {
		return nil, err
	}

	return &Store{
		db: db,
	}, nil
}
