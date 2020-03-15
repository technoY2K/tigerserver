package tigerserver

import (
	"bufio"
	"io"
	"strings"
	"time"
)

// CLI can make calls to the server via terminal client
type CLI struct {
	store PlayerStore
	in    *bufio.Scanner
}

// BlindAlerter interface for any Alert creator
type BlindAlerter interface {
	ScheduleAlertAt(duration time.Duration, amount int)
}

// NewCLI factory function for object
func NewCLI(store PlayerStore, i io.Reader, alerter BlindAlerter) *CLI {
	return &CLI{store: store, in: bufio.NewScanner(i)}
}

// PlayPoker initiates a game
func (c *CLI) PlayPoker() {
	userInput := c.readLine()
	c.store.RecordWin(extractWinner(userInput))
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}

func (c *CLI) readLine() string {
	c.in.Scan()
	return c.in.Text()
}
