// https://github.com/Zeronetsec/Blsys

package codegen

import (
    "os"
    "fmt"
    "strings"
    "path/filepath"
    "io/fs"
    "github.com/Zeronetsec/Blsys/utils"
)

func listEmbeddedScripts() {
    var scripts []string

    err := fs.WalkDir(
        embeddedScripts, "scripts",
        func(
            path string,
            d fs.DirEntry,
            err error,
        ) error {
            if err != nil {
                return err
            }

            if !d.IsDir() {
                rel, err := filepath.Rel("scripts", path)
                if err == nil {
                    name := strings.TrimSuffix(
                        rel, filepath.Ext(rel),
                    )
                    scripts = append(scripts, name)
                }
            }
            return nil
        },
    )

    if err != nil {
        utils.Error(
            fmt.Errorf(
                "failed to list embedded scripts: %w",
                err,
            ),
        )
        os.Exit(1)
    }

    for _, s := range scripts {
        fmt.Println(s)
    }
}

// Copyright (c) 2026 Senzdetta