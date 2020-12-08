package controller

import (
	"fmt"
	"testing"

	. "github.com/franela/goblin"
	"github.com/kcobos/SIGA-Cloud/internal/errors"
	. "github.com/kcobos/SIGA-Cloud/internal/models"
)

func TestNewPlaceController(t *testing.T) {
	g := Goblin(t)
	g.Describe("Controller need to be initialized (#28)", func() {
		places := NewPlaces()
		g.It("Place ID starts at 0", func() {
			g.Assert(places.lastID).Equal(-1)
		})
		g.It("Len of list of place lots must be 0", func() {
			g.Assert(len(places.placeList)).Equal(0)
			length, err := places.Len()
			g.Assert(length).Equal(0)
			g.Assert(err).IsNil()
		})
	})
	g.Describe("Controller not initialized (#28)", func() {
		var places Places
		g.It("If controller is not initialized, raise an error", func() {
			g.Assert(places.lastID).IsZero()
			_, err := places.Len()
			g.Assert(err).Equal(&errors.NotInitialized{})
		})
	})
}

func TestAddNewPlace(t *testing.T) {
	g := Goblin(t)
	g.Describe("Add new place (HU #3)", func() {
		places := NewPlaces()
		var parkings [3]*Parking
		for i := 0; i < 3; i++ {
			parkings[i] = NewParking(i)
		}

		g.It("When a new place is added, update the ID", func() {
			for id := 0; id < 3; id++ {
				newID, err := places.NewPlace(
					37.181024+float64(id), -3.592526+float64(id),
					fmt.Sprintf("Cuesta de las Cabras, nº %d", id),
					[]*Parking{parkings[id]},
				)
				g.Assert(newID).Equal(id)
				g.Assert(err).IsNil()
			}
		})
		g.It("When a new place is added, it must have at least one parking", func() {
			for id := 0; id < 3; id++ {
				parks, freeParks := places.placeList[id].Parkings()
				g.Assert(parks).Equal(1)
				g.Assert(freeParks).Equal(0) // Because parking sensor is just created
			}
		})
		g.It("Parking is attached to the place", func() {
			for id := 0; id < 3; id++ {
				g.Assert(parkings[id].PlaceID()).Equal(id)
			}
		})
	})
	g.Describe("Controller not initialized (#28)", func() {
		var places Places
		g.It("If controller is not initialized, raise an error", func() {
			_, err := places.NewPlace(37.181024, -3.592526, "Cuesta de las Cabras, nº 8", []*Parking{})
			g.Assert(err).Equal(&errors.NotInitialized{})
		})
	})

	g.Describe("Add new place (HU #3). Posible errors", func() {
		places := NewPlaces()
		var parkings [3]*Parking
		for i := 0; i < 3; i++ {
			parkings[i] = NewParking(i)
		}
		for id := 0; id < 3; id++ {
			places.NewPlace(
				37.181024+float64(id), -3.592526+float64(id),
				fmt.Sprintf("Cuesta de las Cabras, nº %d", id),
				[]*Parking{parkings[id]},
			)
		}
		g.It("When a new place is added, update the ID. When fails (Coordinates), don't change ID", func() {
			for id := 0; id < 3; id++ {
				newID, err := places.NewPlace(
					37.181024+float64(id), -3.592526+float64(id),
					fmt.Sprintf("Cuesta de las Cabras, nº %d", id),
					[]*Parking{parkings[id]},
				)
				g.Assert(newID).Equal(-1)
				g.Assert(err).Equal(&errors.PlaceAlreadyExists{})
				g.Assert(places.lastID).Equal(2)
				len, _ := places.Len()
				g.Assert(len).Equal(3)
			}
		})
		g.It("When a new place is added, update the ID. When fails (Address), don't change ID", func() {
			for id := 0; id < 3; id++ {
				newID, err := places.NewPlace(
					37.181024+float64(id)+3, -3.592526+float64(id)+3,
					fmt.Sprintf("Cuesta de las Cabras, nº %d", id),
					[]*Parking{parkings[id]},
				)
				g.Assert(newID).Equal(-1)
				g.Assert(err).Equal(&errors.PlaceAlreadyExists{})
				g.Assert(places.lastID).Equal(2)
				len, _ := places.Len()
				g.Assert(len).Equal(3)
			}
		})
		g.It("When a new place is added, update the ID. When fails (Parkings), don't change ID", func() {
			for id := 0; id < 3; id++ {
				newID, err := places.NewPlace(
					37.181024+float64(id)+3, -3.592526+float64(id)+3,
					fmt.Sprintf("Cuesta de las Cabras, nº %d", id+3),
					[]*Parking{parkings[id]},
				)
				g.Assert(newID).Equal(-1)
				g.Assert(err).Equal(&errors.ParkingSensorIsAlreadyAttached{})
				g.Assert(places.lastID).Equal(2)
				len, _ := places.Len()
				g.Assert(len).Equal(3)
			}
		})
	})
}
