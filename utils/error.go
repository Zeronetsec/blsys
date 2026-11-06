// https://github.com/Zeronetsec/Blsys

package utils

import (
    "os"
    "fmt"
    "github.com/Zeronetsec/Blsys/utils/color"
)

func Error(err error) {
    fmt.Fprintf(
        os.Stderr,
        "%s[!] %sError: %s%v%s\n",
        color.R, color.N, color.GG, err, color.N,
    )
    os.Exit(1)
}

// Copyright (c) 2026 Senzdetta