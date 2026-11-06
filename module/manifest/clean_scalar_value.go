// https://github.com/Zeronetsec/Blsys

package manifest

import (
    "strings"
)

func cleanScalarValue(val string) string {
    val = strings.TrimSpace(val)
    val = strings.TrimSuffix(val, ",")

    if (strings.HasPrefix(
        val, "\"",
    ) && strings.HasSuffix(
        val, "\"",
    )) || (strings.HasPrefix(
        val, "'",
    ) && strings.HasSuffix(
        val, "'",
    )) {
        if len(val) >= 2 {
            return val[1 : len(val)-1]
        }
    }

    return val
}

// Copyright (c) 2026 Senzdetta