// https://github.com/Zeronetsec/blsys

package bindir

import (
    "os"
    "fmt"
)

func Execute(args []string) {
    bin := "/usr/bin"

    if prefix := os.Getenv("PREFIX"); prefix != "" {
        bin = prefix + "/bin"
    }

    fmt.Print(bin)
}

// Copyright (c) 2026 Senzdetta