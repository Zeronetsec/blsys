// https://github.com/Zeronetsec/blsys

package main

import (
    "os"
    "github.com/Zeronetsec/blsys/utils/invinput"
    "github.com/Zeronetsec/blsys/module/help"
    "github.com/Zeronetsec/blsys/module/copy"
    "github.com/Zeronetsec/blsys/module/move"
    "github.com/Zeronetsec/blsys/module/remove"
    "github.com/Zeronetsec/blsys/module/create"
    "github.com/Zeronetsec/blsys/module/fsclear"
    "github.com/Zeronetsec/blsys/module/chmod"
    "github.com/Zeronetsec/blsys/module/symlink"
    "github.com/Zeronetsec/blsys/module/lookpath"
    "github.com/Zeronetsec/blsys/module/codegen"
    "github.com/Zeronetsec/blsys/module/tmpdir"
    "github.com/Zeronetsec/blsys/module/bindir"
    "github.com/Zeronetsec/blsys/module/platform"
    "github.com/Zeronetsec/blsys/module/geteuid"
    "github.com/Zeronetsec/blsys/module/manifest"
    "github.com/Zeronetsec/blsys/module/find"
    "github.com/Zeronetsec/blsys/module/runcmd"
    "github.com/Zeronetsec/blsys/module/optdir"
    "github.com/Zeronetsec/blsys/module/install"
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