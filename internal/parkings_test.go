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
			g.Assert(parkings.lastID).Equal(0)
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
			_, err := parkings.Len()
			g.Assert(err).Equal(&errors.NotInitialized{})
		})
	})
}
