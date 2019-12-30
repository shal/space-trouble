package command

import "github.com/opencars/space-trouble/pkg/domain/model"

type Command interface {
	Prepare()
	Validate() error
}

func Process(c Command) error {
	c.Prepare()

	return model.Validate(c, "command")
}
