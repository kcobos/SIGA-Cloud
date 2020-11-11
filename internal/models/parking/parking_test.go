package parking

import (
	"testing"

	. "github.com/franela/goblin"
)

func TestNewParking(t *testing.T) {
	g := Goblin(t)
	g.Describe("Set up parking sensor", func() {
		parking := NewParking(1)
		g.It("Sensor must an unique ID", func() {
			g.Assert(parking.id).Equal(1)
			g.Assert(parking.placeID).Equal(-1)
			g.Assert(parking.status).Equal(statusNotValid)
		})
	})
}
