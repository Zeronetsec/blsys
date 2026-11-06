// https://github.com/Zeronetsec/Blsys

package platform

import (
    "os"
    "strings"
)

func isProot() bool {
    if os.Getenv(
        "PROOT_TMP_DIR",
    ) != "" || os.Getenv(
        "PROOT_L2S_DIR",
    ) != "" {
        return true
    }

    data, err := os.ReadFile("/proc/self/status")
    if err != nil {
        return false
    }

    for _, line := range strings.Split(string(data), "\n") {
        if strings.HasPrefix(line, "TracerPid:") {
            return strings.TrimSpace(
                strings.TrimPrefix(
                    line, "TracerPid:",
                ),
            ) != "0"
        }
    }

    return false
}

// Copyright (c) 2026 Senzdetta