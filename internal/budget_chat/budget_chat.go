package budgetchat

import "github.com/braidonw/protohackers-go/internal/server"

const (
	welcomeMessage = "Welcome to BudgetChat! What shall I call you?"
)

type BudgetChat struct {
	*server.Config
}
