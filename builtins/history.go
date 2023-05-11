/*
AUTHOR: MELVIN TOWO
COMMAND: History
*/

package builtins

import (
	"errors"
	"fmt"
	"io"
	"strconv"
)

var (
	// Where the commands are stored
	history []string
)

func init() {
	// intializing the var above
	history = make([]string, 0)
}

func AddToHistory(cmd string) {
	// Appending the commmand entered to the history
	history = append(history, cmd)
}

func PrintHistory(w io.Writer, args ...string) error {
	if len(args) > 1 {
		return errors.New("Too many arguments")
	}

	if len(args) == 1 {
		// The function only accepts number arguments, number should be indicating the number of commands in histroy to print
		n, err := strconv.Atoi(args[0])
		if err != nil {
			return errors.New("Arg should be a number")
		}

		if n > len(history) {
			// Setting n to the number of commands in the history if n is bigger than the len of history
			n = len(history)
		}

		for i := len(history) - n; i < len(history); i++ {
			fmt.Fprintln(w, history[i])
		}
	} else {
		// is we arent provided with an argument, we print all  the commands we saved in the history!
		for _, cmd := range history {
			fmt.Fprintln(w, cmd)
		}
	}

	return nil
}
