// https://github.com/Zeronetsec/Blsys

package main

import (
    "os"
    "github.com/Zeronetsec/Blsys/utils/invinput"
    "github.com/Zeronetsec/Blsys/module/help"
    "github.com/Zeronetsec/Blsys/module/copy"
    "github.com/Zeronetsec/Blsys/module/move"
    "github.com/Zeronetsec/Blsys/module/remove"
    "github.com/Zeronetsec/Blsys/module/create"
    "github.com/Zeronetsec/Blsys/module/fsclear"
    "github.com/Zeronetsec/Blsys/module/chmod"
    "github.com/Zeronetsec/Blsys/module/symlink"
    "github.com/Zeronetsec/Blsys/module/lookpath"
    "github.com/Zeronetsec/Blsys/module/codegen"
    "github.com/Zeronetsec/Blsys/module/tmpdir"
    "github.com/Zeronetsec/Blsys/module/bindir"
    "github.com/Zeronetsec/Blsys/module/platform"
    "github.com/Zeronetsec/Blsys/module/geteuid"
    "github.com/Zeronetsec/Blsys/module/manifest"
    "github.com/Zeronetsec/Blsys/module/find"
    "github.com/Zeronetsec/Blsys/module/runcmd"
    "github.com/Zeronetsec/Blsys/module/optdir"
    "github.com/Zeronetsec/Blsys/module/install"
)

type CommandFunc func(
    args []string,
)

func main() {
    args := os.Args[1:]
    if len(args) == 0 {
        invinput.MissingArgument()
        os.Exit(1)
    }

    commands := map[string]CommandFunc{
        "--help": help.Execute,
        "--copy": copy.Execute,
        "--move": move.Execute,
        "--remove": remove.Execute,
        "--create": create.Execute,
        "--fsclear": fsclear.Execute,
        "--chmod": chmod.Execute,
        "--symlink": symlink.Execute,
        "--lookpath": lookpath.Execute,
        "--codegen": codegen.Execute,
        "--tmpdir": tmpdir.Execute,
        "--bindir": bindir.Execute,
        "--platform": platform.Execute,
        "--geteuid": geteuid.Execute,
        "--manifest": manifest.Execute,
        "--find": find.Execute,
        "--runcmd": runcmd.Execute,
        "--optdir": optdir.Execute,
        "--install": install.Execute,
    }

    cmdKey := args[0]
    subArgs := args[1:]

    if fn, exists := commands[cmdKey]; exists {
        fn(subArgs)
    } else {
        invinput.InvalidOption(cmdKey)
        os.Exit(1)
    }
}

// Copyright (c) 2026 Senzdetta