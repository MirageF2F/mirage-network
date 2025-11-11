package adapters

import (
	"github.com/miragef2f/go-peer/pkg/anonymity/adapters"
	"github.com/miragef2f/go-peer/pkg/types"
)

type IRunnerAdapter interface {
	types.IRunner
	adapters.IAdapter
}

type ISettings interface {
	GetNetworkKey() string
	GetWorkSizeBits() uint64
	GetMessageSizeBytes() uint64
}
