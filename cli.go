package tigerserver

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"time"
)

// CLI can make calls to the server via terminal client
type CLI struct {
	store   PlayerStore
	in      *bufio.Scanner
	out     io.Writer
	alerter BlindAlerter
}

// NewCLI factory function for object
func NewCLI(store PlayerStore, i io.Reader, o io.Writer, alerter BlindAlerter) *CLI {
	return &CLI{store: store, in: bufio.NewScanner(i), out: o, alerter: alerter}
}

// PlayPoker initiates a game
func (c *CLI) PlayPoker() {
	fmt.Fprint(c.out, "Please enter the number of players: ")
	c.scheduleBlindAlerts()
	userInput := c.readLine()
	c.store.RecordWin(extractWinner(userInput))
}

func (c *CLI) scheduleBlindAlerts() {
	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
	blindTime := 0 * time.Second

	for _, blind := range blinds {
		c.alerter.ScheduleAlertAt(blindTime, blind)
		blindTime = blindTime + 10*time.Minute
	}
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}

func (c *CLI) readLine() string {
	c.in.Scan()
	return c.in.Text()
}
