// https://github.com/Zeronetsec/blsys

package invinput

import (
    "fmt"
    "os"
    "github.com/Zeronetsec/blsys/utils/color"
)

func MissingArgument() {
    fmt.Fprintf(
        os.Stderr,
        "%s[!] %sMissing argument!\n",
        color.R, color.N,
    )

    fmt.Fprintf(
        os.Stderr,
        "%s[!] %sTry: %sblsys --help%s\n",
        color.R, color.N, color.GG, color.N,
    )
}

// Copyright (c) 2026 Senzdetta