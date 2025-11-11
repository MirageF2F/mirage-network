package handler

import (
	"context"

	"github.com/miragef2f/go-peer/pkg/crypto/asymmetric"
	"github.com/miragef2f/hidden-lake/pkg/request"
	"github.com/miragef2f/hidden-lake/pkg/response"
)

type IHandlerF func(
	context.Context,
	asymmetric.IPubKey,
	request.IRequest,
) (response.IResponse, error)
