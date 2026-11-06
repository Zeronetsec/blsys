// https://github.com/Zeronetsec/blsys

package lookpath

import (
    "fmt"
    "os"
    "strconv"
    "path/filepath"
    "github.com/Zeronetsec/blsys/utils"
    "github.com/Zeronetsec/blsys/utils/invinput"
)

func Execute(args []string) {
    var (
        targetPath string
        rttStr string
    )

    for i := 0; i < len(args); i++ {
        switch args[i] {
            case "--rtt":
                if i+1 >= len(args) {
                    invinput.MissingArgument()
                    os.Exit(1)
                }
                rttStr = args[i+1]
                i++
            default:
                if (len(args[i]) > 0 &&
                    args[i][0] != '-' &&
                    targetPath == "") {
                        targetPath = args[i]
                } else if args[i][0] == '-' {
                    invinput.InvalidOption(
                        fmt.Sprintf(
                            "--lookpath %s",
                            args[i],
                        ),
                    )
                    os.Exit(1)
                }
        }
    }

    if targetPath == "" {
        invinput.MissingArgument()
        os.Exit(1)
    }

    absPath, err := filepath.Abs(targetPath)
    if err != nil {
        utils.Error(
            fmt.Errorf(
                "cannot resolve absolute path: %w",
                err,
            ),
        )
    }

    realPath, err := filepath.EvalSymlinks(absPath)
    if err != nil {
        utils.Error(
            fmt.Errorf(
                "cannot resolve symlinks: %w",
                err,
            ),
        )
    }

    info, err := os.Stat(realPath)
    currentDir := realPath

    if err == nil && !info.IsDir() {
        currentDir = filepath.Dir(realPath)
    }

    rtt := 0
    if rttStr != "" {
        parsedRtt, err := strconv.Atoi(rttStr)
        if err != nil {
            utils.Error(
                fmt.Errorf(
                    "invalid rtt value: %s (%w)",
                    rttStr, err,
                ),
            )
        }
        rtt = parsedRtt
    }

    for j := 0; j < rtt; j++ {
        currentDir = filepath.Dir(currentDir)
    }

    fmt.Print(currentDir)
}

// Copyright (c) 2026 Senzdetta