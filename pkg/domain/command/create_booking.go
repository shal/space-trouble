package command

import (
	"github.com/opencars/space-trouble/pkg/domain/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type CreateBooking struct {
	FirstName     string       `json:"first_name"`
	LastName      string       `json:"last_name"`
	Gender        model.Gender `json:"gender"`
	BirthDate     model.Date   `json:"birth_date"`
	LaunchpadID   string       `json:"launchpad_id"`
	DestinationID string       `json:"destination_id"`
	LaunchDate    model.Date   `json:"launch_date"`
}

func (c *CreateBooking) Prepare() {}

func (c *CreateBooking) Validate() error {
	return validation.ValidateStruct(c,
		validation.Field(
			&c.FirstName,
			validation.Required.Error(model.Required),
			validation.Length(0, 64).Error(model.Invalid),
		),
		validation.Field(
			&c.LastName,
			validation.Required.Error(model.Required),
			validation.Length(0, 64).Error(model.Invalid),
		),
		validation.Field(
			&c.Gender,
			validation.Required.Error(model.Required),
			validation.In(model.Male, model.Female).Error(model.Invalid),
		),
		validation.Field(
			&c.BirthDate,
			validation.Required.Error(model.Required),
			validation.Date(model.DateLayout).Error(model.Invalid),
		),
		validation.Field(
			&c.LaunchpadID,
			validation.Required.Error(model.Required), // TODO: Add spacex format validation.
		),
		validation.Field(
			&c.DestinationID,
			validation.Required.Error(model.Required), // TODO: Add format validation.
		),
		validation.Field(
			&c.LaunchDate,
			validation.Required.Error(model.Required),
			validation.Date(model.DateLayout).Error(model.Invalid),
		),
	)
}
