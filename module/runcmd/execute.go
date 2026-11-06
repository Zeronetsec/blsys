// https://github.com/Zeronetsec/Blsys

package runcmd

import (
    "fmt"
    "os"
    "strings"
    "os/exec"
    "github.com/Zeronetsec/Blsys/utils/color"
    "github.com/Zeronetsec/Blsys/utils/invinput"
)

func Execute(args []string) {
    var (
        commandToRun string
        printMessage string
        noSuppress bool
    )

    for i := 0; i < len(args); i++ {
        switch args[i] {
            case "--runcmd":
                if i+1 < len(args) {
                    commandToRun = args[i+1]
                    i++
                }
            case "--print":
                if i+1 < len(args) {
                    printMessage = args[i+1]
                    i++
                }
            case "--no-suppress":
                noSuppress = true
            default:
                if commandToRun == "" && !strings.HasPrefix(
                    args[i], "-",
                ) {
                    commandToRun = args[i]
                }
        }
    }

    if commandToRun == "" {
        invinput.MissingArgument()
        os.Exit(1)
    }

    if noSuppress {
        if printMessage != "" {
            fmt.Printf(
                "%s[*] %s%s%s\n",
                color.B, color.N, printMessage, color.N,
            )
        } else {
            fmt.Fprintf(
                os.Stderr,
                "%s[!] %sWarning message not set!\n",
                color.YY, color.N,
            )

            fmt.Printf(
                "%s[*] %sRunning command...",
                color.B, color.N,
            )
        }

        cmd := exec.Command("bash", "-c", commandToRun)
        cmd.Stdout = os.Stdout
        cmd.Stderr = os.Stderr

        err := cmd.Run()
        if err != nil {
            exitCode := 1
            if exitError, ok := err.(*exec.ExitError); ok {
                if status, ok := exitError.Sys().(interface{
                    ExitStatus() int
                }); ok {
                    exitCode = status.ExitStatus()
                }
            }
            os.Exit(exitCode)
        }
        os.Exit(0)
    }

    if printMessage != "" {
        fmt.Printf(
            "%s[*] %s%s%s",
            color.B, color.N, printMessage, color.N,
        )
    } else {
        fmt.Fprintf(
            os.Stderr,
            "%s[!] %sWarning message not set!\n",
            color.YY, color.N,
        )

        fmt.Printf(
            "%s[*] %sRunning command...",
            color.B, color.N,
        )
    }

    cmd := exec.Command("bash", "-c", commandToRun)
    var stderrBuf strings.Builder
    cmd.Stderr = &stderrBuf

    err := cmd.Run()
    if err != nil {
        exitCode := 1
        if exitError, ok := err.(*exec.ExitError); ok {
            if status, ok := exitError.Sys().(interface{
                ExitStatus() int
            }); ok {
                exitCode = status.ExitStatus()
            }
        }

        fmt.Printf(
            " %s- %serror: %s%d%s\n",
            color.DG, color.N, color.R, exitCode, color.N,
        )

        errOutput := strings.TrimSpace(stderrBuf.String())
        if errOutput != "" {
            fmt.Fprintln(os.Stderr, errOutput)
        }

        os.Exit(exitCode)
    }

    fmt.Printf(
        " %s- %sexit: %s0%s\n",
        color.DG, color.N, color.GG, color.N,
    )

    os.Exit(0)
}

// Copyright (c) 2026 Senzdetta