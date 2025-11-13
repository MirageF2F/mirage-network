package network

import (
	"context"
	"time"

	"github.com/miragef2f/go-peer/pkg/anonymity"
	"github.com/miragef2f/go-peer/pkg/crypto/asymmetric"
	gopeer_logger "github.com/miragef2f/go-peer/pkg/logger"
	"github.com/miragef2f/go-peer/pkg/types"
	"github.com/miragef2f/hidden-lake/pkg/adapters"
	"github.com/miragef2f/hidden-lake/pkg/request"
	"github.com/miragef2f/hidden-lake/pkg/response"
)

type IHiddenLakeNode interface {
	types.IRunner
	GetOriginNode() anonymity.INode

	SendRequest(context.Context, asymmetric.IPubKey, request.IRequest) error
	FetchRequest(context.Context, asymmetric.IPubKey, request.IRequest) (response.IResponse, error)
}

type ISettings interface {
	ISrvSettings
	IQBPSettings

	GetAdapterSettings() adapters.ISettings
}

type IQBPSettings interface {
	GetPowParallel() uint64
	GetQBPConsumers() uint64
	GetFetchTimeout() time.Duration
	GetQueuePeriod() time.Duration
}

type ISrvSettings interface {
	GetLogger() gopeer_logger.ILogger
	GetFmtAppName() string
}
