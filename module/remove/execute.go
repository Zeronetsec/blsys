// https://github.com/Zeronetsec/blsys

package remove

import (
    "fmt"
    "os"
    "github.com/Zeronetsec/blsys/utils"
    "github.com/Zeronetsec/blsys/utils/invinput"
)

func Execute(args []string) {
    if len(args) == 0 {
        invinput.MissingArgument()
        os.Exit(1)
    }

    target := args[0]

    info, err := os.Lstat(target)
    if err != nil {
        utils.Error(
            fmt.Errorf(
                "cannot access target: %w",
                err,
            ),
        )
    }

    var kind string
    var removeFn func() error

    switch {
        case info.Mode()&os.ModeSymlink != 0:
            kind = "symlink"
            removeFn = func() error {
                return os.Remove(target)
            }
        case info.IsDir():
            kind = "directory"
            removeFn = func() error {
                return os.RemoveAll(target)
            }
        default:
            kind = "file"
            removeFn = func() error {
                return os.Remove(target)
            }
    }

    runner(
        kind, target,
        removeFn,
    )
}

// Copyright (c) 2026 Senzdetta