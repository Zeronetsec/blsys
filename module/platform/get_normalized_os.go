// https://github.com/Zeronetsec/blsys

package platform

import (
    "runtime"
    "strings"
    "os"
)

func getNormalizedOS() string {
    switch runtime.GOOS {
        case "linux":
            if isProot() || isRealDistro() {
                return "linux"
            }

            if strings.Contains(
                os.Getenv("PREFIX"), "com.termux",
            ) {
                return "android"
            }

            if _, err := os.Stat(
                "/system/bin/app_process",
            ); err == nil {
                return "android"
            }

            return "linux"
        case "android":
            if isProot() || isRealDistro() {
                return "linux"
            }

            return "android"
        case "darwin":
            if strings.Contains(
                os.Getenv("HOME"), "/Mobile/Documents",
            ) || runtime.GOARCH == "ios" {
                return "iphone"
            }

            return "macos"
        case "windows":
            return "windows"
        case "chromeos":
            return "chromeos"
        default:
            return runtime.GOOS
    }
}

// Copyright (c) 2026 Senzdetta