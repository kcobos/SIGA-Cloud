package model

import (
	"testing"

	"github.com/franela/goblin"
	"github.com/kcobos/SIGA-Cloud/internal/errors"
)

func TestNewParking(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Set up parking sensor (HU #2)", func() {
		parking := NewParking()
		// g.It("Sensor has an unique ID", func() {
		// 	g.Assert(parking.ID).Equal(1)
		// })
		g.It("Sensor has a status and must be unavailable", func() {
			g.Assert(parking.Status).Equal(StatusNotValid)
		})
		g.It("Sensor must not be attached to any place", func() {
			g.Assert(parking.PlaceID).Equal(int64(-1))
		})
	})
}

func TestChangeStatus(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Update parking status (HU #7)", func() {
		p := NewParking()
		g.It("Parking status must change", func() {
			err := p.ChangeStatus("free")
			g.Assert(err).IsNil()
			status, _ := p.GetStatus()
			g.Assert(status).Equal(Free)
		})
		g.It("Parking status must change", func() {
			err := p.ChangeStatus("occupied")
			g.Assert(err).IsNil()
			status, _ := p.GetStatus()
			g.Assert(status).Equal(Occupied)
		})
		g.It("Parking status must not change on invalid status", func() {
			err := p.ChangeStatus("other")
			g.Assert(err).Equal(&errors.StatusNotValid{})
			status, _ := p.GetStatus()
			g.Assert(status).Equal(Occupied)
		})
		g.It("Parking status must be logged", func() {
			g.Assert(len(p.History)).Equal(2)
		})
	})
}
