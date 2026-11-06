// https://github.com/Zeronetsec/blsys

package find

import (
    "fmt"
    "os"
    "strings"
    "io/fs"
    "path/filepath"
    "github.com/Zeronetsec/blsys/utils"
    "github.com/Zeronetsec/blsys/utils/invinput"
)

func Execute(args []string) {
    var (
        searchPath string
        targetType string = "all"
        containsPattern string
        onlyBasename bool
    )

    for i := 0; i < len(args); i++ {
        switch args[i] {
            case "--find":
                if i+1 < len(args) {
                    searchPath = args[i+1]
                    i++
                }
            case "--type":
                if i+1 < len(args) {
                    targetType = strings.ToLower(args[i+1])
                    i++
                }
            case "--contains":
                if i+1 < len(args) {
                    containsPattern = args[i+1]
                    i++
                }
            case "--basename":
                onlyBasename = true
            default:
                if searchPath == "" && !strings.HasPrefix(
                    args[i], "-",
                ) {
                    searchPath = args[i]
                }
        }
    }

    if searchPath == "" {
        invinput.MissingArgument()
        os.Exit(1)
    }

    err := filepath.WalkDir(
        searchPath, func(
            path string,
            d fs.DirEntry,
            err error,
        ) error {
            if err != nil {
                return nil
            }

            if path == searchPath {
                return nil
            }

            info, err := d.Info()
            if err != nil {
                return nil
            }

            isDir := d.IsDir()
            isSymlink := (info.Mode() & os.ModeSymlink) != 0
            isFile := !isDir && !isSymlink

            matchType := false
            switch targetType {
                case "file":
                    if isFile {
                        matchType = true
                    }
                case "dir":
                    if isDir {
                        matchType = true
                    }
                case "symlink":
                    if isSymlink {
                        matchType = true
                    }
                case "all":
                    matchType = true
            }

            if !matchType {
                return nil
            }

            if containsPattern != "" {
                matchedBase, _ := filepath.Match(
                    containsPattern, d.Name(),
                )

                matchedFull, _ := filepath.Match(
                    containsPattern, path,
                )

                if !matchedBase && !matchedFull {
                    return nil
                }
            }

            outputVal := path
            if onlyBasename {
                outputVal = filepath.Base(path)
            }

            fmt.Println(outputVal)
            return nil
        },
    )

    if err != nil {
        utils.Error(
            fmt.Errorf(
                "error walking path: %w",
                err,
            ),
        )
    }
}

// Copyright (c) 2026 Senzdetta