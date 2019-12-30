package command

import (
	"github.com/opencars/space-trouble/pkg/domain/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type DeleteBooking struct {
	ID string
}

func (c *DeleteBooking) Prepare() {}

func (c *DeleteBooking) Validate() error {
	return validation.ValidateStruct(c,
		validation.Field(
			&c.ID,
			validation.Required.Error(model.Required),
			validation.Length(0, 64).Error(model.Invalid),
		),
	)
}
