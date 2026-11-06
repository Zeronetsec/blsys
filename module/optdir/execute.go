// https://github.com/Zeronetsec/Blsys

package optdir

import (
    "os"
    "fmt"
)

func Execute(args []string) {
    opt := "/opt"

    if _, err := os.Stat("/etc/os-release"); err != nil {
        opt = os.Getenv("PREFIX") + "/opt"
    }

    fmt.Print(opt)
}

// Copyright (c) 2026 Senzdetta