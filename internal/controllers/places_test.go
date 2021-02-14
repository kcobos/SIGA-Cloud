package controller

import (
	"fmt"
	"testing"

	"github.com/franela/goblin"
	"github.com/kcobos/SIGA-Cloud/common"
	"github.com/kcobos/SIGA-Cloud/internal/errors"
)

func TestNewPlaceController(t *testing.T) {
	g := goblin.Goblin(t)
	conf := common.NewConf("../../common/conf.yaml")
	database := common.CreateDBConnection(
		conf,
	)
	g.Describe("Controller need to be initialized (#28)", func() {
		places := NewPlaces(database, conf)
		parkings := NewParkings(database, conf)

		g.It("Len of list of place lots must be 0", func() {
			length, err := places.Len()
			g.Assert(length).Equal(0)
			g.Assert(err).IsNil()
		})
		g.It("Place ID starts at 1", func() {
			parkings.NewParking()
			newID, err := places.NewPlace(
				37.181024, -3.592526,
				"Cuesta de las Cabras, nº %d",
				[]int64{1}, parkings,
			)
			g.Assert(newID).Equal(int64(1))
			g.Assert(err).IsNil()
		})
	})
	// g.Describe("Controller not initialized (#28)", func() {
	// 	var places Places

	// 	g.It("If controller is not initialized, raise an error", func() {
	// 		_, err := places.Len()
	// 		g.Assert(err).IsNotNil()
	// 	})
	// })
}

func TestAddNewPlace(t *testing.T) {
	g := goblin.Goblin(t)
	conf := common.NewConf("../../common/conf.yaml")
	database := common.CreateDBConnection(
		conf,
	)

	g.Describe("Add new place (HU #3)", func() {
		places := NewPlaces(database, conf)
		parkings := NewParkings(database, conf)

		for i := 1; i < 4; i++ {
			parkings.NewParking()
		}

		g.It("When a new place is added, update the ID", func() {
			for id := int64(1); id < 4; id++ {
				newID, err := places.NewPlace(
					37.181024+float64(id), -3.592526+float64(id),
					fmt.Sprintf("Cuesta de las Cabras, nº %d", id),
					[]int64{id}, parkings,
				)
				g.Assert(err).IsNil()
				g.Assert(newID).Equal(id)
			}
		})
		g.It("When a new place is added, it must have at least one parking", func() {
			newID, err := places.NewPlace(
				37.181024+float64(4), -3.592526+float64(4),
				fmt.Sprintf("Cuesta de las Cabras, nº %d", 4),
				[]int64{}, parkings,
			)
			g.Assert(err).Equal(&errors.NeedsAParkingSensor{})
			g.Assert(newID).Equal(int64(-1))
		})
		g.It("Parking is attached to the place", func() {
			for id := int64(1); id < 4; id++ {
				parking, _ := parkings.GetParking(id)
				g.Assert(parking.PlaceID).Equal(id)
			}
		})
	})
	g.Describe("Controller not initialized (#28)", func() {
		// var places Places
		// var parkings Parkings

		// g.It("If controller is not initialized, raise an error", func() {
		// 	_, err := places.NewPlace(
		// 		37.181024, -3.592526,
		// 		"Cuesta de las Cabras, nº 8",
		// 		[]int64{}, &parkings)
		// 	g.Assert(err).IsNotNil()
		// })
	})

	g.Describe("Add new place (HU #3). Posible errors", func() {
		places := NewPlaces(database, conf)
		parkings := NewParkings(database, conf)

		for i := 1; i < 5; i++ {
			parkings.NewParking()
		}

		for id := int64(1); id < 4; id++ {
			places.NewPlace(
				37.181024+float64(id), -3.592526+float64(id),
				fmt.Sprintf("Cuesta de las Cabras, nº %d", id),
				[]int64{id}, parkings,
			)
		}
		g.It(`When a new place is added, update the ID.
			When fails (Coordinates), don't change ID`, func() {
			newID, err := places.NewPlace(
				37.181024+3, -3.592526+3,
				fmt.Sprintf("Cuesta del Chapiz, nº %d", 4),
				[]int64{4}, parkings,
			)
			g.Assert(err).Equal(&errors.PlaceAlreadyExists{})
			g.Assert(newID).Equal(int64(-1))
			len, _ := places.Len()
			g.Assert(len).Equal(3)
		})
		g.It(`When a new place is added, update the ID.
			When fails (Address), don't change ID`, func() {
			newID, err := places.NewPlace(
				37.179856+4, -3.588329+4,
				fmt.Sprintf("Cuesta de las Cabras, nº %d", 3),
				[]int64{4}, parkings,
			)
			g.Assert(err).Equal(&errors.PlaceAlreadyExists{})
			g.Assert(newID).Equal(int64(-1))
			len, _ := places.Len()
			g.Assert(len).Equal(3)
		})
		g.It(`When a new place is added, update the ID.
			When fails (Parkings), don't change ID`, func() {
			newID, err := places.NewPlace(
				37.179856+4, -3.588329+4,
				fmt.Sprintf("Cuesta del Chapiz, nº %d", 4),
				[]int64{3}, parkings,
			)
			g.Assert(err).Equal(&errors.ParkingSensorIsAlreadyAttached{})
			g.Assert(newID).Equal(int64(-1))
			len, _ := places.Len()
			g.Assert(len).Equal(3)
		})
	})
}
