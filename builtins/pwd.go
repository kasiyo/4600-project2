/*
	AUTHOR: 		Kaia Siripanyo
	SHELL COMMAND:	pwd
*/

package builtins

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

var (
	InvalidArgCount = errors.New("invalid argument count")
	CurrDir, _      = os.Getwd()
)

func PrintWorkingDirectory(w io.Writer, args ...string) error {
	wd := make([]string, 0)
	for i := 0; i < len(CurrDir); i++ {
		wd = append(wd, string(CurrDir[i]))
	}
	_, err := fmt.Fprintln(w, strings.Join(wd, ""))

	if len(args) == 0 {
		return err
	} else {
		return fmt.Errorf("%w: expected zero arguments", InvalidArgCount)
	}
}