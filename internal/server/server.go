package server

import (
	"context"
	"net"
)

type Server interface {
	Configuration

	Setup() context.Context
	Handle(ctx context.Context, conn net.Conn)
}
