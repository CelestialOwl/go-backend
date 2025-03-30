package dbrepo

import (
	"context"
	"time"

	"github.com/tsawler/bookings-app/internal/models"
)

func (m *postgresDBRepo) AllUsers() bool {
	return true
}

// InsertReservation inserts a reservation into the database
// and returns an error if something goes wrong.
// It takes a models.Reservation struct as an argument.
// The function uses a context with a timeout of 3 seconds
// to ensure that the database operation does not hang indefinitely.
// It prepares an SQL statement to insert the reservation data
// into the "reservations" table, including fields like first name,
// last name, email, phone number, start date, end date, room ID,
// created at, and updated at timestamps.
// The function executes the SQL statement with the provided reservation data
func (m *postgresDBRepo) InsertReservation(res models.Reservation) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	stmt := `INSERT INTO reservations (first_name, last_name, email, phone, start_date, end_date, room_id, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id`

	var newId int
	// Insert the reservation into the database
	err := m.DB.QueryRowContext(ctx, stmt,
		res.FirstName,
		res.LastName,
		res.Email,
		res.Phone,
		res.StartDate,
		res.EndDate,
		res.RoomID,
		time.Now(),
		time.Now(),
	).Scan(&newId)
	// Check for errors
	if err != nil {
		return 0, nil
	}

	return newId, nil
}

func (m *postgresDBRepo) InsertRoomRestriction(r models.RoomRestriction) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `INSERT INTO room_restrictions (start_date, end_date, room_id, reservation_id, restriction_id, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`

	var newId int
	err := m.DB.QueryRowContext(ctx, stmt,
		r.StartDate,
		r.EndDate,
		r.RoomID,
		r.ReservationID,
		r.RestrictionID,
		time.Now(),
		time.Now(),
	).Scan(&newId)
	if err != nil {
		return 0, err
	}

	return newId, nil
}

// SearchAvailabilityByDates checks if there are any room restrictions
func (m *postgresDBRepo) SearchAvailabilityByDates(start, end time.Time, roomID int) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var numRows int
	query := `SELECT count(id) from room_restrictions WHERE room_id = $1 and $2 < end_date AND $3 > start_date;`

	row := m.DB.QueryRowContext(ctx, query, roomID, start, end)
	err := row.Scan(&numRows)
	if err != nil {
		return false, err
	}
	if numRows == 0 {
		return true, nil
	} else {
		return false, nil
	}
}
