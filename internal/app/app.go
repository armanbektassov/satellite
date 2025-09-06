package app

import (
	"net/http"
)

type App struct {
	serviceProvider *serviceProvider
	httpServer      *http.Server
}
