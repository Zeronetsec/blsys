// https://github.com/Zeronetsec/Blsys

package platform

import (
    "os"
)

func isRealDistro() bool {
    _, err := os.Stat("/etc/os-release")
    return err == nil
}

// Copyright (c) 2026 Senzdetta