// https://github.com/Zeronetsec/blsys

package platform

import (
    "os"
    "fmt"
    "strings"
    "github.com/Zeronetsec/blsys/utils"
    "github.com/Zeronetsec/blsys/utils/invinput"
)

func Execute(args []string) {
    var supportFile string

    for i := 0; i < len(args); i++ {
        switch args[i] {
            case "--support":
                if i+1 >= len(args) {
                    invinput.MissingArgument()
                    os.Exit(1)
                }

                supportFile = args[i+1]
                i++
            default:
                invinput.InvalidOption(
                    fmt.Sprintf(
                        "--platform %s",
                        args[i],
                    ),
                )
                os.Exit(1)
        }
    }

    currentPlatform := getNormalizedOS()
    if supportFile == "" {
        fmt.Print(currentPlatform)
        return
    }

    supportedPlatforms, err := parseManifest(supportFile)
    if err != nil {
        utils.Error(
            fmt.Errorf(
                "failed to parse manifest: %w",
                err,
            ),
        )
    }

    isSupported := false
    for _, p := range supportedPlatforms {
        if strings.EqualFold(p, currentPlatform) {
            isSupported = true
            break
        }
    }

    if !isSupported {
        utils.Error(
            fmt.Errorf(
                "platform: %s is not supported by manifest",
                currentPlatform,
            ),
        )
    }
}

// Copyright (c) 2026 Senzdetta