// https://github.com/Zeronetsec/blsys

package help

import (
    "fmt"
    "github.com/Zeronetsec/blsys/utils/color"
)

func Execute(args []string) {
    fmt.Printf(
        "%s[ %sSenzdetta Build System & Scripting Utility %s]%s\n" +
        "%s[ %sDevelop by %s@Senzdetta %s(%sgithub.com/Senzdetta%s) ]%s\n" +
        "%s[ %sPublished at %sgithub.com/Zeronetsec/blsys %s]%s\n" +
        "%s[ %sLicensed under %sApache 2.0 %s]%s\n",
        color.DG, color.N, color.DG, color.N,
        color.DG, color.N, color.GG, color.DG, color.GG, color.DG, color.N,
        color.DG, color.N, color.GG, color.DG, color.N,
        color.DG, color.N, color.GG, color.DG, color.N,
    )

    fmt.Println()

    fmt.Printf(
        "Usage: %sblsys %s<option> [<args>]%s\n",
        color.GG, color.CC, color.N,
    )

    fmt.Println()

    fmt.Printf(
        "%s* %s--help%s\n",
        color.DG, color.GG, color.N,
    )

    fmt.Printf(
        "%s* %s--copy %s--type <dir|file> --source <src> --dest <dest> [--force]%s\n",
        color.DG, color.GG, color.CC, color.N,
    )

    fmt.Printf(
        "%s* %s--move %s--type <dir|file> --source <src> --dest <dest> [--force]%s\n",
        color.DG, color.GG, color.CC, color.N,
    )

    fmt.Printf(
        "%s* %s--remove %s<target>%s\n",
        color.DG, color.GG, color.CC, color.N,
    )

    fmt.Printf(
        "%s* %s--create %s--type <dir|file> --target <path> [--mode <perm>] [--force]%s\n",
        color.DG, color.GG, color.CC, color.N,
    )

    fmt.Printf(
        "%s* %s--fsclear %s--type <dir|file> --target <path>%s\n",
        color.DG, color.GG, color.CC, color.N,
    )

    fmt.Printf(
        "%s* %s--chmod %s--mode <perm> --target <path>%s\n",
        color.DG, color.GG, color.CC, color.N,
    )

    fmt.Printf(
        "%s* %s--symlink %s--source <src> --dest <dest> [--mode <perm>] [--force]%s\n",
        color.DG, color.GG, color.CC, color.N,
    )

    fmt.Printf(
        "%s* %s--lookpath %s<path> [--rtt <idx>]%s\n",
        color.DG, color.GG, color.CC, color.N,
    )

    fmt.Printf(
        "%s* %s--codegen %s<script|--list>%s\n",
        color.DG, color.GG, color.CC, color.N,
    )

    fmt.Printf(
        "%s* %s--tmpdir%s\n",
        color.DG, color.GG, color.N,
    )

    fmt.Printf(
        "%s* %s--bindir%s\n",
        color.DG, color.GG, color.N,
    )

    fmt.Printf(
        "%s* %s--platform %s[--support <manifest>]%s\n",
        color.DG, color.GG, color.CC, color.N,
    )

    fmt.Printf(
        "%s* %s--geteuid%s\n",
        color.DG, color.GG, color.N,
    )

    fmt.Printf(
        "%s* %s--manifest %s<manifest> --get <key>%s\n",
        color.DG, color.GG, color.CC, color.N,
    )

    fmt.Printf(
        "%s* %s--find %s<path> [--type <all|file|dir|symlink>] [--contains <pattern>] [--basename]%s\n",
        color.DG, color.GG, color.CC, color.N,
    )

    fmt.Printf(
        "%s* %s--runcmd %s<cmds...> --print <message> [--no-suppress]%s\n",
        color.DG, color.GG, color.CC, color.N,
    )

    fmt.Printf(
        "%s* %s--optdir%s\n",
        color.DG, color.GG, color.N,
    )

    fmt.Printf(
        "%s* %s--install %s--mode <direct|file> --package <input> [--backend <cmd...>]%s\n",
        color.DG, color.GG, color.CC, color.N,
    )
}

// Copyright (c) 2026 Senzdetta