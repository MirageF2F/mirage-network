
package http

import (
	"net/http"
	"time"

	"github.com/miragef2f/go-peer/pkg/logger"
	"github.com/miragef2f/hidden-lake/pkg/adapters"
)

type IHTTPAdapter interface {
	adapters.IRunnerAdapter

	WithLogger(string, logger.ILogger) IHTTPAdapter
	WithHandlers(map[string]http.HandlerFunc) IHTTPAdapter
	GetOnlines() []string
}

type ISettings interface {
	GetAdapterSettings() adapters.ISettings
	GetAddress() string
	GetReadTimeout() time.Duration
	GetHandleTimeout() time.Duration
}
