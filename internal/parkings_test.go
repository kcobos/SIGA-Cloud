package controller

import (
	"testing"

	. "github.com/franela/goblin"
	"github.com/kcobos/SIGA-Cloud/internal/errors"
)

func TestNewParkingController(t *testing.T) {
	g := Goblin(t)
	g.Describe("Controller need to be initialized", func() {
		parkings := NewParkings()
		g.It("Parking ID starts at 0", func() {
			g.Assert(parkings.lastID).Equal(-1)
		})
		g.It("Len of list of parking lots must be 0", func() {
			g.Assert(len(parkings.parkingList)).Equal(0)
			length, err := parkings.Len()
			g.Assert(length).Equal(0)
			g.Assert(err).IsNil()
		})
	})
	g.Describe("Controller not initialized", func() {
		var parkings Parkings
		g.It("If controller is not initialized, raise an error", func() {
			g.Assert(parkings.lastID).IsZero()
			_, err := parkings.Len()
			g.Assert(err).Equal(&errors.NotInitialized{})
		})
	})
}

func TestAddNewParking(t *testing.T) {
	g := Goblin(t)
	g.Describe("Add new parking", func() {
		parkings := NewParkings()
		g.It("When a new parking is added, update the ID", func() {
			for id := 0; id < 3; id++ {
				new_id, err := parkings.NewParking()
				g.Assert(new_id).Equal(id)
				g.Assert(err).IsNil()
			}
		})
		g.It("When a new parking is added, it is not attached to any place", func() {
			for id := 0; id < 3; id++ {
				g.Assert(parkings.parkingList[id].PlaceID()).Equal(-1)
			}
		})
		g.It("The first status of the parking sensor must be unavailable", func() {
			for id := 0; id < 3; id++ {
				_, err := parkings.parkingList[id].Status()
				g.Assert(err).Equal(&errors.StatusNotValid{})
			}
		})
	})
}
