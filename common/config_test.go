package common

import (
	"os"
	"testing"

	"github.com/franela/goblin"
)

func TestNewConfiguration(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Create configuration", func() {
		os.Setenv("DB_HOST", "siga-db")
		os.Setenv("DB_PORT", "5432")
		os.Setenv("DB_NAME", "places")
		os.Setenv("DB_USER", "siga")
		os.Setenv("DB_PASS", "siga2021")

		c := NewConf("conf.yaml")
		g.It("Configuration has to be equal than in file", func() {
			g.Assert(c.DB.Host).Equal("siga-db")
			g.Assert(c.DB.Port).Equal(5432)
			g.Assert(c.DB.Name).Equal("places")
			g.Assert(c.DB.User).Equal("siga")
			g.Assert(c.DB.Pass).Equal("siga2021")
			g.Assert(c.DB.Test).IsTrue()
		})
	})
}
