package api

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/Cprime50/gin-blog/config"
	db "github.com/Cprime50/gin-blog/db/sqlc"
)

func newTestServer(t *testing.T, store db.Store) *Server {
	conf := config.LoadConfig("test", "../env/")
	server := NewServer(conf, store, nil, nil)
	server.MountHandlers()
	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
