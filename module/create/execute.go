// https://github.com/Zeronetsec/blsys

package create

import (
    "fmt"
    "os"
    "strconv"
    "github.com/Zeronetsec/blsys/utils"
    "github.com/Zeronetsec/blsys/utils/color"
    "github.com/Zeronetsec/blsys/utils/invinput"
)

func Execute(args []string) {
    var (
        force bool
        typ string
        target string
        modeStr string
    )

    for i := 0; i < len(args); i++ {
        var valPtr *string

        switch args[i] {
            case "--force":
                force = true
                continue
            case "--type":
                valPtr = &typ
            case "--target":
                valPtr = &target
            case "--mode":
                valPtr = &modeStr
            default:
                invinput.InvalidOption(
                    fmt.Sprintf(
                        "--create %s",
                        args[i],
                    ),
                )
                os.Exit(1)
        }

        if i+1 >= len(args) || (len(
            args[i+1],
        ) > 1 && args[i+1][:2] == "--") {
            invinput.MissingArgument()
            os.Exit(1)
        }

        i++
        *valPtr = args[i]
    }

    if typ == "" || target == "" {
        invinput.MissingArgument()
        os.Exit(1)
    }

    if _, err := os.Stat(target); err == nil && !force {
        fmt.Fprintf(
            os.Stderr,
            "%s[!] %sTarget: %s%s %sis already exist!\n",
            color.R, color.N, color.GG, target, color.N,
        )
        os.Exit(1)
    }

    var fileMode os.FileMode
    if modeStr != "" {
        parsed, err := strconv.ParseUint(
            modeStr, 8, 32,
        )

        if err != nil {
            utils.Error(
                fmt.Errorf(
                    "invalid mode format: %s (%w)",
                    modeStr, err,
                ),
            )
        }

        fileMode = os.FileMode(parsed)
    } else {
        if typ == "dir" {
            fileMode = 0755
        } else {
            fileMode = 0644
        }
    }

    switch typ {
        case "file":
            runner(
                "file", target,
                func() error {
                    return createFile(
                        target, fileMode,
                    )
                },
            )
        case "dir":
            runner(
                "directory", target,
                func() error {
                    return os.MkdirAll(
                        target, fileMode,
                    )
                },
            )
        default:
            invinput.InvalidOption(
                fmt.Sprintf(
                    "--create --type %s",
                    typ,
                ),
            )
            os.Exit(1)
    }
}

// Copyright (c) 2026 Senzdetta