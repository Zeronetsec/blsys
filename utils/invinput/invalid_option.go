// https://github.com/Zeronetsec/Blsys

package invinput

import (
    "fmt"
    "os"
    "github.com/Zeronetsec/Blsys/utils/color"
)

func InvalidOption(input string) {
    fmt.Fprintf(
        os.Stderr,
        "%s[!] %sInvalid option: %s%s%s\n",
        color.R, color.N, color.GG, input, color.N,
    )

    fmt.Fprintf(
        os.Stderr,
        "%s[!] %sTry: %sblsys --help%s\n",
        color.R, color.N, color.GG, color.N,
    )
}

// Copyright (c) 2026 Senzdetta