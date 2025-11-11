package tcp

import (
	"time"

	"github.com/miragef2f/go-peer/pkg/logger"
	"github.com/miragef2f/go-peer/pkg/network/connkeeper"
	"github.com/miragef2f/hidden-lake/pkg/adapters"
)

type ITCPAdapter interface {
	adapters.IRunnerAdapter

	WithLogger(string, logger.ILogger) ITCPAdapter
	GetConnKeeper() connkeeper.IConnKeeper
}

type ISettings interface {
	ISrvSettings

	GetAdapterSettings() adapters.ISettings
}

type ISrvSettings interface {
	GetAddress() string
	GetConnNumLimit() uint64
	GetConnKeepPeriod() time.Duration
	GetSendTimeout() time.Duration
	GetRecvTimeout() time.Duration
	GetDialTimeout() time.Duration
	GetWaitTimeout() time.Duration
}
