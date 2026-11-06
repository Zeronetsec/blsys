// https://github.com/Zeronetsec/blsys

package install

import (
    "fmt"
    "os"
    "strings"
    "path/filepath"
    "github.com/Zeronetsec/blsys/utils/color"
    "github.com/Zeronetsec/blsys/utils/invinput"
)

func Execute(args []string) {
    var (
        mode string
        packageArg string
        backendArg string
    )

    for i := 0; i < len(args); i++ {
        switch args[i] {
            case "--mode":
                if i+1 < len(args) {
                    mode = strings.ToLower(args[i+1])
                    i++
                }
            case "--packages":
                if i+1 < len(args) {
                    packageArg = args[i+1]
                    i++
                }
            case "--backend":
                if i+1 < len(args) {
                    backendArg = args[i+1]
                    i++
                }
            default:
                if mode == "" && !strings.HasPrefix(args[i], "-") {
                    mode = args[i]
                }
        }
    }

    if mode == "" || packageArg == "" {
        invinput.MissingArgument()
        os.Exit(1)
    }

    if mode == "file" && backendArg != "" {
        fmt.Fprintf(
            os.Stderr,
            "%s[!] %sCustom backend not allowed in file mode!\n",
            color.R, color.N,
        )
        os.Exit(1)
    }

    if mode != "direct" && mode != "file" {
        fmt.Fprintf(
            os.Stderr,
            "%s[!] %sInvalid mode: %s%s%s\n",
            color.R, color.N, color.GG, mode, color.N,
        )
        os.Exit(1)
    }

    prefix := os.Getenv("PREFIX")
    if prefix == "" {
        prefix = "/usr"
    }

    binDir := filepath.Join(prefix, "bin")
    _ = os.MkdirAll(binDir, 0755)

    BackendRegistry = make(map[string]string)
    if mode == "direct" && backendArg != "" {
        parseBackendString(backendArg)
    }

    if mode == "direct" {
        fmt.Printf(
            "%s[*] %sInstalling: %s%s%s\n",
            color.B, color.N, color.GG, packageArg, color.N,
        )
        code := parseAndExecuteItem(packageArg, "", binDir)
        os.Exit(code)
    } else if mode == "file" {
        err := processInstallFile(packageArg, binDir)
        if err != nil {
            fmt.Fprintf(
                os.Stderr,
                "%s[!] %sFailed to process install file: %s%v%s\n",
                color.R, color.N, color.GG, err, color.N,
            )
            os.Exit(1)
        }
        os.Exit(0)
    }
}

// Copyright (c) 2026 Senzdetta