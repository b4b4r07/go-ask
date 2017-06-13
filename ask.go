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
	defaultPrompt string = `%s [y/n]: `
	defaultRetry  int    = 1
)

type Q struct {
	PromptFormat string
	Retries      int
	Input        io.Reader
	Output       io.Writer
}

func NewQ() *Q {
	return &Q{
		PromptFormat: defaultPrompt,
		Retries:      defaultRetry,
		Input:        os.Stdin,
		Output:       os.Stdout,
	}
}

func (q *Q) Confirm(s string) bool {
	r := bufio.NewReader(q.Input)
	t := q.Retries
	for ; t > 0; t-- {
		fmt.Fprintf(q.Output, q.PromptFormat, s)
		res, err := r.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		// Empty input (i.e. "\n")
		if len(res) < 2 {
			continue
		}

		return strings.ToLower(strings.TrimSpace(res))[0] == 'y'
	}

	if q.Retries > 1 && t == 0 {
		fmt.Fprintln(q.Output, "Retries over")
	}

	return false
}

func (q *Q) Confirmf(qfmt string, a ...interface{}) bool {
	return q.Confirm(fmt.Sprintf(qfmt, a...))
}

// Password is not implemented yet
func Password(s string) error {
	return nil
}
