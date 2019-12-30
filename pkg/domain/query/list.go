package query

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"

	"github.com/opencars/space-trouble/pkg/domain/model"
)

type List struct {
	Offset string
	Limit  string
}

func (q *List) Prepare() {}

func (q *List) Validate() error {
	return validation.ValidateStruct(q,
		validation.Field(
			&q.Offset,
			validation.Required.Error(model.Required),
		),
		validation.Field(
			&q.Limit,
			validation.Required.Error(model.Required),
		),
	)
}
