package sqlstore_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/opencars/space-trouble/pkg/domain/model"
	"github.com/opencars/space-trouble/pkg/domain/query"
	"github.com/opencars/space-trouble/pkg/store/sqlstore"
)

func TestBookingRepository_Create(t *testing.T) {
	s, teardown := sqlstore.TestDB(t, conf)
	defer teardown("bookings")

	booking := model.TestBooking(t)
	assert.NoError(t, s.Booking().Create(context.TODO(), booking))
}

func TestBookingRepository_List(t *testing.T) {
	s, teardown := sqlstore.TestDB(t, conf)
	defer teardown("booking")

	booking := model.TestBooking(t)
	assert.NoError(t, s.Booking().Create(context.TODO(), booking))

	actual, err := s.Booking().List(context.TODO(), &query.List{})
	assert.NoError(t, err)
	assert.Equal(t, booking, actual)
}
