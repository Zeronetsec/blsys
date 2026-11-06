// https://github.com/Zeronetsec/blsys

package codegen

import (
    "fmt"
    "os"
    "io/fs"
    "path/filepath"
    "github.com/Zeronetsec/blsys/utils"
    "github.com/Zeronetsec/blsys/utils/invinput"
)

func Execute(args []string) {
    if len(args) == 0 {
        invinput.MissingArgument()
        os.Exit(1)
    }

    rawInput := args[0]
    if rawInput == "--list" {
        listEmbeddedScripts()
        return
    }

    fileName := rawInput
    if filepath.Ext(fileName) == "" {
        fileName = fileName + ".sh"
    }

    targetPath := filepath.Join("scripts", fileName)
    content, err := fs.ReadFile(embeddedScripts, targetPath)
    if err != nil {
        utils.Error(
            fmt.Errorf(
                "script: %s not found (%w)",
                rawInput, err,
            ),
        )
        os.Exit(1)
    }

    os.Stdout.Write(content)
}

// Copyright (c) 2026 Senzdetta