package smoketest

import (
	"context"
	"io"
	"log"
	"net"

	"github.com/braidonw/protohackers-go/internal/server"
)

type SmokeTest struct{ *server.Config }

func (SmokeTest) Setup() context.Context { return context.TODO() }

func (SmokeTest) Handle(ctx context.Context, conn net.Conn) {
	defer conn.Close()

	if _, err := io.Copy(conn, conn); err != nil {
		log.Println(err)
	}
}

func Run() {
	cfg := server.NewConfig(10000)
	cfg.ParseFlags()

	s := SmokeTest{cfg}
	server.RunTCPServer(s)
}
