// https://github.com/Zeronetsec/blsys

package symlink

import (
    "fmt"
    "os"
    "github.com/Zeronetsec/blsys/utils/color"
)

func runner(src, dest string, fn func() error) {
    fmt.Printf(
        "%s[*] %sCreating symlink: %s%s %s-> %s%s%s",
        color.B, color.N, color.GG, src,
        color.DG, color.GG, dest, color.N,
    )

    if err := fn(); err != nil {
        fmt.Printf(
            " %s- %serror: %s1%s\n",
            color.DG, color.N, color.R, color.N,
        )

        fmt.Fprintf(
            os.Stderr,
            "%s[!] %sFailed to create symlink: %s%v%s\n",
            color.R, color.N, color.GG, err, color.N,
        )

        os.Exit(1)
    }

    fmt.Printf(
        " %s- %sexit: %s0%s\n",
        color.DG, color.N, color.GG, color.N,
    )
}

// Copyright (c) 2026 Senzdetta