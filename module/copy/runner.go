// https://github.com/Zeronetsec/Blsys

package copy

import (
    "os"
    "fmt"
    "github.com/Zeronetsec/Blsys/utils/color"
)

func runner(kind, src, dest string, fn func() error) {
    fmt.Printf(
        "%s[*] %sCopying %s: %s%s %s-> %s%s%s",
        color.B, color.N, kind, color.GG, src,
        color.DG, color.GG, dest, color.N,
    )

    if err := fn(); err != nil {
        fmt.Printf(
            " %s- %serror: %s1%s\n",
            color.DG, color.N, color.R, color.N,
        )

        fmt.Fprintf(
            os.Stderr,
            "%s[!] %sFailed to copy %s: %s%v%s\n",
            color.R, color.N, kind, color.GG, err, color.N,
        )

        os.Exit(1)
    }

    fmt.Printf(
        " %s- %sexit: %s0%s\n",
        color.DG, color.N, color.GG, color.N,
    )
}

// Copyright (c) 2026 Senzdetta