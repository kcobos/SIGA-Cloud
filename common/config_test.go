package common

import (
	"os"
	"strconv"
	"testing"

	"github.com/franela/goblin"
)

func TestNewConfiguration(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Create configuration", func() {

		c := NewConf("conf.yaml")
		g.It("Configuration has to be equal than in file", func() {
			g.Assert(c.DB.Host).Equal(os.Getenv("DB_HOST"))
			g.Assert(c.DB.Port).Equal(strconv.Atoi(os.Getenv("DB_PORT")))
			g.Assert(c.DB.Name).Equal(os.Getenv("DB_NAME"))
			g.Assert(c.DB.User).Equal(os.Getenv("DB_USER"))
			g.Assert(c.DB.Pass).Equal(os.Getenv("DB_PASS"))
			g.Assert(c.DB.Test).IsTrue()
		})
	})
}
