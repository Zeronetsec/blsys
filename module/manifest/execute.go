// https://github.com/Zeronetsec/Blsys

package manifest

import (
    "fmt"
    "os"
    "strings"
    "github.com/Zeronetsec/Blsys/utils"
)

func Execute(args []string) {
    var (
        manifestFile string
        targetKey string
    )

    for i := 0; i < len(args); i++ {
        switch args[i] {
            case "--manifest":
                if i+1 < len(args) {
                    manifestFile = args[i+1]
                    i++
                }
            case "--get":
                if i+1 < len(args) {
                    targetKey = args[i+1]
                    i++
                }
            default:
                if manifestFile == "" && !strings.HasPrefix(
                    args[i], "-",
                ) {
                    manifestFile = args[i]
                } else if targetKey == "" && !strings.HasPrefix(
                    args[i], "-",
                ) {
                    targetKey = args[i]
                }
        }
    }

    if manifestFile == "" && len(args) > 0 && !strings.HasPrefix(
        args[0], "-",
    ) {
        manifestFile = args[0]
    }

    if targetKey == "" && len(args) > 1 && !strings.HasPrefix(
        args[1], "-",
    ) {
        targetKey = args[1]
    }

    if manifestFile == "" || targetKey == "" {
        utils.Error(
            fmt.Errorf(
                "missing required arguments",
            ),
        )
    }

    data, err := parseManifestFile(manifestFile)
    if err != nil {
        utils.Error(
            fmt.Errorf(
                "failed to parse manifest: %w",
                err,
            ),
        )
    }

    val, exists := data[targetKey]
    if !exists {
        os.Exit(1)
    }

    fmt.Print(val)
}

// Copyright (c) 2026 Senzdetta