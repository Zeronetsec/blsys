// https://github.com/Zeronetsec/blsys

package tmpdir

import (
    "os"
    "fmt"
)

func Execute(args []string) {
    tmp := os.Getenv("TMPDIR")

    if tmp == "" {
        if prefix := os.Getenv("PREFIX"); prefix != "" {
            tmp = prefix + "/tmp"
        } else {
            tmp = "/tmp"
        }
    }

    fmt.Print(tmp)
}

// Copyright (c) 2026 Senzdetta