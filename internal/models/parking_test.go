package model

import (
	"testing"

	. "github.com/franela/goblin"
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
