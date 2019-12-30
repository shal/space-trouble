package query

import "github.com/opencars/space-trouble/pkg/domain/model"

type Query interface {
	Prepare()
	Validate() error
}

func Process(q Query) error {
	q.Prepare()

	return model.Validate(q, "query")
}
