package http

import (
	"encoding/json"
	"net/http"
	"runtime"

	"github.com/gorilla/mux"
	"github.com/opencars/space-trouble/pkg/httputil"

	"github.com/opencars/space-trouble/pkg/domain"
	"github.com/opencars/space-trouble/pkg/domain/command"
	"github.com/opencars/space-trouble/pkg/domain/query"
	"github.com/opencars/space-trouble/pkg/version"
)

type server struct {
	router *mux.Router

	svc domain.CustomerService
}

func newServer(svc domain.CustomerService) *server {
	srv := server{
		router: mux.NewRouter(),
		svc:    svc,
	}

	srv.configureRouter()

	return &srv
}

func (*server) Version() httputil.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		v := struct {
			Version string `json:"version"`
			Go      string `json:"go"`
		}{
			Version: version.Version,
			Go:      runtime.Version(),
		}

		return json.NewEncoder(w).Encode(v)
	}
}

func (s *server) CreateBooking() httputil.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		var c command.CreateBooking
		if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
			return err
		}

		booking, err := s.svc.CreateBooking(r.Context(), &c)
		if err != nil {
			return handleErr(err)
		}

		return json.NewEncoder(w).Encode(booking)
	}
}

func (s *server) DeleteBooking() httputil.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		c := command.DeleteBooking{
			ID: mux.Vars(r)["id"],
		}

		if err := s.svc.DeleteBooking(r.Context(), &c); err != nil {
			return handleErr(err)
		}

		return nil
	}
}

func (s *server) ListBookings() httputil.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		q := query.List{
			Limit:  r.URL.Query().Get("limit"),
			Offset: r.URL.Query().Get("offset"),
		}

		bookins, err := s.svc.ListBookings(r.Context(), &q)
		if err != nil {
			return handleErr(err)
		}

		return json.NewEncoder(w).Encode(bookins)
	}
}
