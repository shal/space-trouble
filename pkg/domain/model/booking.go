package model

type Gender string

const (
	Male   Gender = "male"
	Female Gender = "female"
)

type Booking struct {
	ID            string `db:"id"`
	FirstName     string `db:"first_name"`
	LastName      string `db:"last_name"`
	Gender        Gender `db:"gender"`
	BirthDate     Date   `db:"birth_date"`
	LaunchpadID   string `db:"launchpad_id"`
	DestinationID string `db:"destination_id"`
	LaunchDate    Date   `db:"launch_date"`
}
