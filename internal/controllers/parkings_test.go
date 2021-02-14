package controller

import (
	"testing"

	"github.com/franela/goblin"
	"github.com/kcobos/SIGA-Cloud/common"
	"github.com/kcobos/SIGA-Cloud/internal/errors"
	model "github.com/kcobos/SIGA-Cloud/internal/models"
)

func TestNewParkingController(t *testing.T) {
	g := goblin.Goblin(t)
	conf := common.NewConf("../../common/conf.yaml")
	database := common.CreateDBConnection(
		conf,
	)
	g.Describe("Controller need to be initialized (#24)", func() {
		parkings := NewParkings(database, conf)

		g.It("Len of list of parking lots must be 0", func() {
			length, err := parkings.Len()
			g.Assert(length).Equal(0)
			g.Assert(err).IsNil()
		})
		g.It("Parking ID starts at 1", func() {
			newID, err := parkings.NewParking()
			g.Assert(newID).Equal(int64(1))
			g.Assert(err).IsNil()
		})
	})
	// g.Describe("Controller not initialized (#24)", func() {
	// 	var parkings Parkings

	// 	g.It("If controller is not initialized, raise an error", func() {
	// 		_, err := parkings.Len()
	// 		g.Assert(err).IsNotNil()
	// 	})
	// })
}

func TestAddNewParking(t *testing.T) {
	g := goblin.Goblin(t)
	conf := common.NewConf("../../common/conf.yaml")
	database := common.CreateDBConnection(
		conf,
	)
	g.Describe("Add new parking (HU #2)", func() {
		parkings := NewParkings(database, conf)

		g.It("When a new parking is added, update the ID", func() {
			for id := int64(1); id < 4; id++ {
				newID, err := parkings.NewParking()
				g.Assert(newID).Equal(id)
				g.Assert(err).IsNil()
			}
		})
		g.It("When a new parking is added, it is not attached to any place", func() {
			for id := int64(1); id < 4; id++ {
				parking, _ := parkings.GetParking(id)
				g.Assert(parking.PlaceID).Equal(int64(-1))
			}
		})
		g.It("The first status of the parking sensor must be unavailable", func() {
			for id := int64(1); id < 4; id++ {
				parking, _ := parkings.GetParking(id)
				_, err := parking.GetStatus()
				g.Assert(err).Equal(&errors.StatusNotValid{})
			}
		})
	})
}

func TestChangeStatus(t *testing.T) {
	g := goblin.Goblin(t)
	conf := common.NewConf("../../common/conf.yaml")
	database := common.CreateDBConnection(
		conf,
	)
	g.Describe("Update parking status (HU #7)", func() {
		parkings := NewParkings(database, conf)
		parkings.NewParking()
		p, _ := parkings.GetParking(1)
		g.Assert(p.History).IsNotNil()

		g.It("Parking status must change", func() {
			err := parkings.ChangeStatus(1, "free")
			g.Assert(err).IsNil()
			p, _ := parkings.GetParking(1)
			status, err := p.GetStatus()
			g.Assert(err).IsNil()
			g.Assert(status).Equal(model.Free)
		})
		g.It("Parking status must change again", func() {
			err := parkings.ChangeStatus(1, "occupied")
			g.Assert(err).IsNil()
			p, _ := parkings.GetParking(1)
			status, err := p.GetStatus()
			g.Assert(err).IsNil()
			g.Assert(status).Equal(model.Occupied)
		})
		g.It("Parking status must be logged", func() {
			statuses, _ := parkings.CountHistory(int64(1))
			g.Assert(statuses).Equal(2)
		})
	})
}
