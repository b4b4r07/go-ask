// Package ask provides some simple question methods
package ask

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const (
	defaultPrompt  string = `%s [y/n]: `
	defaultRetry   int    = 1
	defaultNewline bool   = true
)

// Q is question object
type Q struct {
	// Prompt is the format of the question text
	Prompt string
	// Retry indicates the number of repetitions of Y/N
	Retry int
	// Newline returns true if allow empty input (e.g. newline)
	Newline bool

	stdin  io.Reader
	stdout io.Writer
	stderr io.Writer
}

// NewQ returns new Q
func NewQ() *Q {
	return &Q{
		Prompt:  defaultPrompt,
		Retry:   defaultRetry,
		Newline: defaultNewline,
		stdin:   os.Stdin,
		stdout:  os.Stdout,
		stderr:  os.Stderr,
	}
}

// Confirm makes a simple closed question
func (q *Q) Confirm(s string) bool {
	r := bufio.NewReader(q.stdin)
	t := q.Retry
	for ; t > 0; t-- {
		fmt.Fprintf(q.stdout, q.Prompt, s)
		res, err := r.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		// Empty input (i.e. "\n")
		if q.Newline && len(res) < 2 {
			continue
		}
		return strings.ToLower(strings.TrimSpace(res))[0] == 'y'
	}
	if q.Retry > 1 && t == 0 {
		fmt.Fprintln(q.stdout, "Retry over")
	}
	return false
}

// Confirmf is Confirm with a trailing "f" wrap fmt.Sprintf
func (q *Q) Confirmf(qfmt string, a ...interface{}) bool {
	return q.Confirm(fmt.Sprintf(qfmt, a...))
}

// Password is not implemented yet
func Password(s string) error {
	return nil
}
