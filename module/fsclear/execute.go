// https://github.com/Zeronetsec/blsys

package fsclear

import (
    "fmt"
    "os"
    "github.com/Zeronetsec/blsys/utils"
    "github.com/Zeronetsec/blsys/utils/invinput"
)

func Execute(args []string) {
    var (
        typ string
        target string
    )

    for i := 0; i < len(args); i++ {
        var valPtr *string

        switch args[i] {
            case "--type":
                valPtr = &typ
            case "--target":
                valPtr = &target
            default:
                invinput.InvalidOption(
                    fmt.Sprintf(
                        "--fsclear %s",
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

    if _, err := os.Stat(target); err != nil {
        utils.Error(
            fmt.Errorf(
                "cannot access target: %w",
                err,
            ),
        )
    }

    switch typ {
        case "file":
            runner(
                "file", target,
                func() error {
                    return clearFile(target)
                },
            )
        case "dir":
            runner(
                "directory", target,
                func() error {
                    return clearDir(target)
                },
            )
        default:
            invinput.InvalidOption(
                fmt.Sprintf(
                    "--fsclear --type %s",
                    typ,
                ),
            )
            os.Exit(1)
    }
}

// Copyright (c) 2026 Senzdetta