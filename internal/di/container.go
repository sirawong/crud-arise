package di

import (
	"net/http"

	"github.com/sirawong/crud-arise/pkg/config"
)

type Application struct {
	httpServer *http.Server
	Cfg        *config.Config
}
