// https://github.com/Zeronetsec/blsys

package move

import (
    "fmt"
    "os"
    "github.com/Zeronetsec/blsys/utils"
    "github.com/Zeronetsec/blsys/utils/color"
    "github.com/Zeronetsec/blsys/utils/invinput"
)

func Execute(args []string) {
    var (
        force bool
        typ string
        src string
        dest string
    )

    for i := 0; i < len(args); i++ {
        switch args[i] {
            case "--force":
                force = true
                continue
            case "--type", "--source", "--dest":
                if i+1 >= len(args) || (len(
                        args[i+1],
                ) > 1 && args[i+1][:2] == "--") {
                    invinput.MissingArgument()
                    os.Exit(1)
                }

                val := args[i+1]
                i++

                switch args[i-1] {
                    case "--type":
                        typ = val
                    case "--source":
                        src = val
                    case "--dest":
                        dest = val
                }
            default:
                invinput.InvalidOption(
                    fmt.Sprintf(
                        "--move %s",
                        args[i],
                    ),
                )
                os.Exit(1)
        }
    }

    if typ == "" || src == "" || dest == "" {
        invinput.MissingArgument()
        os.Exit(1)
    }

    info, err := os.Stat(src)
    if err != nil {
        utils.Error(
            fmt.Errorf(
                "cannot access source: %w",
                err,
            ),
        )
    }

    if _, err := os.Stat(dest); err == nil && !force {
        fmt.Fprintf(
            os.Stderr,
            "%s[!] %sDestination: %s%s %sis already exist!\n",
            color.R, color.N, color.GG, dest, color.N,
        )
        os.Exit(1)
    }

    switch typ {
        case "file":
            if info.IsDir() {
                utils.Error(
                    fmt.Errorf(
                        "source is a directory",
                    ),
                )
            }
            runner(
                "file",
                src, dest,
                func() error {
                    return moveItem(
                        src, dest,
                        force,
                    )
                },
            )
        case "dir":
            if !info.IsDir() {
                utils.Error(
                    fmt.Errorf(
                        "source is a file",
                    ),
                )
            }
            runner(
                "directory",
                src, dest,
                func() error {
                    return moveItem(
                        src, dest,
                        force,
                    )
                },
            )
        default:
            invinput.InvalidOption(
                fmt.Sprintf(
                    "--move --type %s",
                    typ,
                ),
            )
            os.Exit(1)
    }
}

// Copyright (c) 2026 Senzdetta