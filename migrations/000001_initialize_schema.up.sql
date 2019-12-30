CREATE TYPE gender_t AS ENUM ('male', 'female');

CREATE TABLE bookings
(
    "id"             VARCHAR(32) NOT NULL,
    "first_name"     TEXT        NOT NULL,
    "last_name"      TEXT        NOT NULL,
    "gender"         gender_t    NOT NULL,
    "birth_date"     DATE        NOT NULL,
    "launchpad_id"   TEXT        NOT NULL,
    "launch_date"    DATE        NOT NULL,
    "destination_id" TEXT        NOT NULL,
    "created_at"     TIMESTAMP   NOT NULL DEFAULT NOW(),

    PRIMARY KEY ("id")
);

CREATE INDEX bookings_launchpad_id   ON bookings("launchpad_id");
CREATE INDEX bookings_destination_id ON bookings("destination_id");
