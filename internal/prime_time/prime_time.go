package primetime

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"log"
	"net"

	"github.com/braidonw/protohackers-go/internal/server"
)

func Run() {
	cfg := server.NewConfig(10001)
	s := PrimeTime{cfg}
	server.RunTCPServer(s)
}

type PrimeTime struct {
	*server.Config
}

func (PrimeTime) Setup() context.Context { return context.TODO() }

type request struct {
	Method string `json:"method"`
	Number int    `json:"number"`
}

type response struct {
	Method string `json:"method"`
	Prime  bool   `json:"prime"`
}

func isPrime(n int) bool {
	if n <= 3 {
		return n > 1
	} else if n%2 == 0 || n%3 == 0 {
		return false
	}

	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}

	return true
}

func (PrimeTime) Handle(_ctx context.Context, conn net.Conn) {
	addr := conn.RemoteAddr()
	log.Printf("Accepted connection from: %v", addr)

	defer func() {
		conn.Close()
		log.Printf("Closed connection from: %v", addr)
	}()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		in := scanner.Bytes()

		var out []byte
		var req request

		err := json.Unmarshal(in, &req)
		if err != nil || req.Method != "isPrime" {
			log.Printf("Error unmarshalling request: %v", err)
			out = []byte(`{"error": "invalid request"}`)
		} else {
			resp, err := handleRequest(req)
			if err != nil {
				log.Printf("Error handling request: %v", err)
				out = []byte(`{"error": "invalid request"}`)
			}

			out, err = json.Marshal(resp)
			if err != nil {
				log.Printf("Error marshalling response: %v", err)
				out = []byte(`{"error": "invalid request"}`)
			}
		}

		out = append(out, byte('\n'))

		if _, err := conn.Write(out); err != nil {
			log.Printf("Error writing response: %v", err)
			return
		}

		log.Printf("%#q ⇒ %#q (%v)", in, out[:len(out)-1], addr)
	}
}

func handleRequest(req request) (response, error) {
	correctMethod := req.Method == "isPrime"

	if !correctMethod {
		return response{}, errors.New("invalid request")
	}

	isAPrime := isPrime(req.Number)
	return response{Method: req.Method, Prime: isAPrime}, nil
}
