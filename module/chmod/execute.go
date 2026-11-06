// https://github.com/Zeronetsec/blsys

package chmod

import (
    "fmt"
    "os"
    "strconv"
    "github.com/Zeronetsec/blsys/utils"
    "github.com/Zeronetsec/blsys/utils/invinput"
)

func Execute(args []string) {
    var (
        modeStr string
        target string
    )

    for i := 0; i < len(args); i++ {
        var valPtr *string

        switch args[i] {
            case "--mode":
                valPtr = &modeStr
            case "--target":
                valPtr = &target
            default:
                invinput.InvalidOption(
                    fmt.Sprintf(
                        "--chmod %s",
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

    if modeStr == "" || target == "" {
        invinput.MissingArgument()
        os.Exit(1)
    }

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

    fileMode := os.FileMode(parsed)

    runner(
        target, modeStr,
        func() error {
            return os.Chmod(target, fileMode)
        },
    )
}

// Copyright (c) 2026 Senzdetta