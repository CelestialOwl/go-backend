package dbrepo

import (
	"errors"
	"time"

	"github.com/tsawler/bookings-app/internal/models"
)

func (m *TestDBRepo) AllUsers() bool {
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
func (m *TestDBRepo) InsertReservation(res models.Reservation) (int, error) {
	if res.RoomID == 13 {
		return 0, errors.New("room error")
	}

	return 1, nil
}

func (m *TestDBRepo) InsertRoomRestriction(r models.RoomRestriction) (int, error) {
	if r.RoomID == 10000 {
		return 0, errors.New("restriction error")
	}
	return 1, nil
}

// SearchAvailabilityByDatesByRoomId checks if there are any room restrictions
func (m *TestDBRepo) SearchAvailabilityByDatesByRoomId(start, end time.Time, roomID int) (bool, error) {
	return false, nil
}

func (m *TestDBRepo) SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error) {
	var rooms []models.Room

	return rooms, nil
}

func (m *TestDBRepo) GetRoomByID(id int) (models.Room, error) {
	var room models.Room
	if id > 2 {
		return room, errors.New("Some error")
	}
	return room, nil
}

func (m *TestDBRepo) GetUserById(id int) (models.User, error) {
	var u models.User

	return u, nil
}

func (m *TestDBRepo) UpdateUser(u models.User) error {
	return nil
}

func (m *TestDBRepo) Authenticate(email, testPassword string) (int, string, error) {
	return 1, "", nil
}

// AllReservations retrieves all reservations from the database
func (m *TestDBRepo) AllReservations() ([]models.Reservation, error) {
	var reservations []models.Reservation

	return reservations, nil
}

// AllReservations retrieves all reservations from the database
func (m *TestDBRepo) AllNewReservations() ([]models.Reservation, error) {

	var reservations []models.Reservation

	return reservations, nil

}

// return one reservation by id
func (m *TestDBRepo) GetReservationById(id int) (models.Reservation, error) {

	var res models.Reservation

	return res, nil
}
