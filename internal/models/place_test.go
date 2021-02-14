package model

import (
	"testing"

	"github.com/franela/goblin"
	"github.com/kcobos/SIGA-Cloud/internal/errors"
)

func TestNewPlace(t *testing.T) {
	g := goblin.Goblin(t)

	parking1 := NewParking()
	parking2 := NewParking()

	parking2.ChangeStatus("free")

	place1, _ := NewPlace(37.181024, -3.592526, "Cuesta de las Cabras, nº 8", []*Parking{parking1})
	place2, errPlace2 := NewPlace(37.181024, -3.592526, "Cuesta de las Cabras, nº 8", []*Parking{})
	place3, errPlace3 := NewPlace(37.181024, -3.592526, "Cuesta de las Cabras, nº 8", []*Parking{parking1})
	place4, _ := NewPlace(37.181024, -3.592526, "Cuesta de las Cabras, nº 8", []*Parking{parking2})

	g.Describe("Set up parking place (#27, HU #3)", func() {
		g.It("Place has a unique ID", func() {
			g.Assert(place1.ID).Equal(int64(0))
		})
		g.It("Place has a unique location", func() {
			lat, lon := place1.Location()
			g.Assert(lat).Equal(37.181024)
			g.Assert(lon).Equal(-3.592526)
		})
		g.It("Place has a unique address", func() {
			g.Assert(place1.Address).Equal("Cuesta de las Cabras, nº 8")
		})
		g.It("Place has at least one parking sensor", func() {
			parkings, freeParkings := place1.GetParkings()
			g.Assert(parkings).Equal(1)
			g.Assert(freeParkings).Equal(0)
		})
		g.It("Parking sensor is attached to the place", func() {
			g.Assert(parking1.PlaceID).Equal(int64(0))
		})
		g.It("Place must have at least one parking sensor", func() {
			g.Assert(place2).Equal((*Place)(nil))
			g.Assert(errPlace2).Equal(&errors.NeedsAParkingSensor{})
		})
		g.It("There can be no places that share parking sensors", func() {
			g.Assert(place3).Equal((*Place)(nil))
			g.Assert(errPlace3).Equal(&errors.ParkingSensorIsAlreadyAttached{})
		})
		g.It("Update free parking lots of the place", func() {
			parkings, freeParkings := place4.GetParkings()
			g.Assert(parkings).Equal(1)
			g.Assert(freeParkings).Equal(1)
		})
	})
}
