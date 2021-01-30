package conf

import (
	"testing"

	"github.com/franela/goblin"
)

func TestNewConfiguration(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Create configuration", func() {
		c := NewConf()
		g.It("Configuration has to be equal than in file", func() {
			g.Assert(c.DB.Host).Equal("siga-db")
			g.Assert(c.DB.Port).Equal(5432)
			g.Assert(c.DB.Name).Equal("places")
			g.Assert(c.DB.User).Equal("siga")
			g.Assert(c.DB.Pass).Equal("siga2021")
		})
	})
}
