package model

import (
	"testing"

	. "github.com/franela/goblin"
	"github.com/kcobos/SIGA-Cloud/internal/errors"
)

func TestNewParking(t *testing.T) {
	g := Goblin(t)
	g.Describe("Set up parking sensor (HU #2)", func() {
		parking := NewParking(1)
		g.It("Sensor has an unique ID", func() {
			g.Assert(parking.id).Equal(1)
		})
		g.It("Sensor has a status and must be unavailable", func() {
			g.Assert(parking.status).Equal(statusNotValid)
		})
		g.It("Sensor must not be attached to any place", func() {
			g.Assert(parking.placeID).Equal(-1)
		})
	})
}

func TestChangeStatus(t *testing.T) {
	g := Goblin(t)
	g.Describe("Update parking status (HU #7", func() {
		p := NewParking(1)
		g.It("Parking status must change", func() {
			changed, _ := p.ChangeStatus("free")
			g.Assert(changed).IsTrue()
			status, _ := p.Status()
			g.Assert(status).Equal(free)
		})
		g.It("Parking status must change", func() {
			changed, _ := p.ChangeStatus("occupied")
			g.Assert(changed).IsTrue()
			status, _ := p.Status()
			g.Assert(status).Equal(occupied)
		})
		g.It("Parking status must not change on invalid status", func() {
			changed, err := p.ChangeStatus("other")
			g.Assert(changed).IsFalse()
			g.Assert(err).Equal(&errors.StatusNotValid{})
			status, _ := p.Status()
			g.Assert(status).Equal(occupied)
		})
		g.It("Parking status must be logged", func() {
			g.Assert(p.history.Len()).Equal(2)
		})
	})
}
