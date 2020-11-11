package parking

import (
	"testing"

	. "github.com/franela/goblin"
)

func TestParkingAttributes(t *testing.T) {
	g := Goblin(t)
	g.Describe("Set up parking sensor", func() {
		p := NewParking(1)
		g.It("Sensor must an unique ID", func() {
			g.Assert(p.id).Equal(1)
			g.Assert(p.placeID).Equal(-1)
			g.Assert(p.status).Equal(statusNotValid)
		})
	})
}
