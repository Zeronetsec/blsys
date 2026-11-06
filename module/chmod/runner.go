// https://github.com/Zeronetsec/Blsys

package chmod

import (
    "fmt"
    "os"
    "github.com/Zeronetsec/Blsys/utils/color"
)

func runner(target, mode string, fn func() error) {
    fmt.Printf(
        "%s[*] %sSet permission: %s%s %s(%s%s%s)%s",
        color.B, color.N, color.GG, target,
        color.DG, color.GG, mode, color.DG, color.N,
    )

    if err := fn(); err != nil {
        fmt.Printf(
            " %s- %serror: %s1%s\n",
            color.DG, color.N, color.R, color.N,
        )

        fmt.Fprintf(
            os.Stderr,
            "%s[!] %sFailed to set permission for: %s%s %s(%s%v%s)%s\n",
            color.R, color.N, color.GG, target,
            color.DG, color.CC, err, color.DG, color.N,
        )

        os.Exit(1)
    }

    fmt.Printf(
        " %s- %sexit: %s0%s\n",
        color.DG, color.N, color.GG, color.N,
    )
}

// Copyright (c) 2026 Senzdetta