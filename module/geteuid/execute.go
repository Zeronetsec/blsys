// https://github.com/Zeronetsec/blsys

package geteuid

import (
    "fmt"
    "runtime"
)

func Execute(args []string) {
    if runtime.GOOS == "windows" {
        fmt.Print("0")
        return
    }

    uid := getSystemEUID()
    fmt.Print(uid)
}

// Copyright (c) 2026 Senzdetta