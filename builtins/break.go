package builtins

import (
    "errors"
    "fmt"
    "io"
    "os"
)

var (
    ErrorBreak = errors.New("break")
)

func Break(w io.Writer, args ...string) error {
    if len(args) > 0 {
        return fmt.Errorf("break: too many arguments")
    }
    return ErrorBreak
}