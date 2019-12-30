package http

import "github.com/opencars/seedwork/httputil"

func (s *server) configureRouter() {
	v1 := s.router.PathPrefix("/api/v1/").Subrouter()
	v1.Use(
		httputil.CustomerTokenMiddleware(),
	)

	v1.Handle("/bookings", s.CreateBooking()).Methods("POST")
	v1.Handle("/bookings/{id}", s.DeleteBooking()).Methods("DELETE")
	v1.Handle("/bookings", s.ListBookings()).Methods("GET")

}
