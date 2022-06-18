package http

func (s *server) configureRouter() {
	v1 := s.router.PathPrefix("/api/v1/").Subrouter()

	v1.Handle("/version", s.Version())
	v1.Handle("/bookings", s.CreateBooking()).Methods("POST")
	v1.Handle("/bookings/{id}", s.DeleteBooking()).Methods("DELETE")
	v1.Handle("/bookings", s.ListBookings()).Methods("GET")
}
