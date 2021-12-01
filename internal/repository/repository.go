package repository

import (
	"time"

	"github.com/hafizmfadli/bookings-web/internal/models"
)

// "contract" yang harus dipenuhi suatu type agar dianggap sebagai DatabaseRepo
type DatabaseRepo interface {
	AllUsers() bool
	InsertReservation(res models.Reservation) (int, error)
	InsertRoomRestriction(r models.RoomRestriction) error
	SearchAvailabilityByDatesByRoomID(start, end time.Time, roomID int) (bool, error)
	SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error)
}
