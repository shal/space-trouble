package http

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/opencars/space-trouble/pkg/httputil"
	"github.com/opencars/space-trouble/pkg/logger"

	"github.com/opencars/space-trouble/pkg/domain/model"
)

func handleErr(err error) error {
	logger.Errorf("%s", err)

	var e model.Error
	if errors.As(err, &e) {
		return httputil.NewError(http.StatusBadRequest, e.Error())
	}

	var vErr model.ValidationError
	if errors.As(err, &vErr) {
		errMessage := make([]string, 0)
		for k, vv := range vErr.Messages {
			for _, v := range vv {
				errMessage = append(errMessage, fmt.Sprintf("%s.%s", k, v))
			}
		}

		return httputil.NewError(http.StatusBadRequest, errMessage...)
	}

	return err
}
