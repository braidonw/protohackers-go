package primetime

import "github.com/braidonw/protohackers-go/internal/server"

type PrimeTime struct {
	*server.Config
}

type message struct {
	Method string `json:method`
	Number int    `json:number`
}
