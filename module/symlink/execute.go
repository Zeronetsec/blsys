// https://github.com/Zeronetsec/blsys

package symlink

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
        src string
        dest string
        modeStr string
    )

    for i := 0; i < len(args); i++ {
        switch args[i] {
            case "--force":
                force = true
                continue
            case "--source", "--dest", "--mode":
                if i+1 >= len(args) || (len(
                    args[i+1],
                ) > 1 && args[i+1][:2] == "--") {
                    invinput.MissingArgument()
                    os.Exit(1)
                }

                val := args[i+1]
                i++

                switch args[i-1] {
                    case "--source":
                        src = val
                    case "--dest":
                        dest = val
                    case "--mode":
                        modeStr = val
                }
            default:
                invinput.InvalidOption(
                    fmt.Sprintf(
                        "--symlink %s",
                        args[i],
                    ),
                )
                os.Exit(1)
        }
    }

    if src == "" || dest == "" {
        invinput.MissingArgument()
        os.Exit(1)
    }

    if _, err := os.Lstat(dest); err == nil && !force {
        fmt.Fprintf(
            os.Stderr,
            "%s[!] %sDestination: %s%s %sis already exist!\n",
            color.R, color.N, color.GG, dest, color.N,
        )
        os.Exit(1)
    }

    var fileMode os.FileMode
    hasMode := false

    if modeStr != "" {
        parsed, err := strconv.ParseUint(
            modeStr, 8, 32,
        )

        if err != nil {
            utils.Error(
                fmt.Errorf(
                    "invalid mode format '%s': %w",
                    modeStr, err,
                ),
            )
        }

        fileMode = os.FileMode(parsed)
        hasMode = true
    }

    runner(
        src, dest,
        func() error {
            return createSymlink(
                src, dest, force,
                hasMode, fileMode,
            )
        },
    )
}

// Copyright (c) 2026 Senzdetta