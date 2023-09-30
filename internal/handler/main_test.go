package handler

import (
	"github.com/Bakhram74/advertisement-server.git/internal/config"
	"github.com/Bakhram74/advertisement-server.git/internal/service"
	"github.com/Bakhram74/advertisement-server.git/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
	"time"
)

func newTestHandler(t *testing.T, service *service.Service) *Handler {
	conf := config.Config{
		TokenSymmetricKey:   utils.RandomString(32),
		AccessTokenDuration: time.Minute,
	}

	handler, err := NewHandler(conf, service)
	require.NoError(t, err)

	return handler
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
