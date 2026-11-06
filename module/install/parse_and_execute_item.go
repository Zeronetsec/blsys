// https://github.com/Zeronetsec/blsys

package install

import (
    "strings"
    "fmt"
    "path/filepath"
    "github.com/Zeronetsec/blsys/utils/color"
)

func parseAndExecuteItem(
    rawInput, defaultBackend, binDir string,
) int {
    rawInput = strings.TrimSpace(rawInput)
    if rawInput == "" {
        return 1
    }

    targetAlias := ""
    input := rawInput

    if strings.Contains(rawInput, "->") {
        parts := strings.SplitN(rawInput, "->", 2)
        input = strings.TrimSpace(parts[0])

        targetAlias = strings.ReplaceAll(parts[1], " ", "")
        targetAlias = strings.ReplaceAll(targetAlias, "\t", "")
    }

    processedPkg := ""

    if strings.Contains(input, "::") {
        parts := strings.Split(input, "::")
        allowSkip := false

        for _, part := range parts {
            part = strings.TrimSpace(part)
            if part == "" {
                continue
            }

            if part == "-" {
                allowSkip = true
                continue
            }

            matches := BackendMacroRegex.FindStringSubmatch(part)
            currentBackend := defaultBackend
            pkgTarget := part

            if len(matches) == 3 {
                backendAlias := matches[1]
                pkgTarget = strings.TrimSpace(matches[2])

                if registeredCmd, exists := BackendRegistry[
                    backendAlias,
                ]; exists {
                    currentBackend = registeredCmd
                }
            }

            fmt.Printf(
                "%s-> %sTry: %s%s%s",
                color.DG, color.N, color.GG, pkgTarget, color.N,
            )

            exitCode := executeBackend(
                pkgTarget, currentBackend,
            )

            if exitCode == 0 {
                fmt.Printf(
                    " %s- %sexit: %s0%s\n",
                    color.DG, color.N, color.GG, color.N,
                )
                processedPkg = pkgTarget
                break
            } else {
                fmt.Printf(
                    " %s- %serror: %s%d%s\n",
                    color.DG, color.N, color.R, exitCode, color.N,
                )
            }
        }

        if processedPkg == "" && allowSkip {
            fmt.Printf(
                "%s-> %sSkipped...",
                color.DG, color.N,
            )
            return 0
        }
    } else {
        matches := BackendMacroRegex.FindStringSubmatch(input)
        currentBackend := defaultBackend
        pkgTarget := input

        if len(matches) == 3 {
            backendAlias := matches[1]
            pkgTarget = strings.TrimSpace(matches[2])

            if registeredCmd, exists := BackendRegistry[
                backendAlias,
            ]; exists {
                currentBackend = registeredCmd
            }
        }

        fmt.Printf(
            "%s-> %sTry: %s%s%s",
            color.DG, color.N, color.GG, pkgTarget, color.N,
        )

        exitCode := executeBackend(
            pkgTarget, currentBackend,
        )

        if exitCode == 0 {
            fmt.Printf(
                " %s- %sexit: %s0%s\n",
                color.DG, color.N, color.GG, color.N,
            )
            processedPkg = pkgTarget
        } else {
            fmt.Printf(
                " %s- %serror: %s%d%s\n",
                color.DG, color.N, color.R, exitCode, color.N,
            )
        }
    }

    if processedPkg != "" && targetAlias != "" {
        targetPath := filepath.Join(
            binDir, targetAlias,
        )

        errCode := createSymlink(
            processedPkg, targetAlias, binDir,
        )

        if errCode == 0 {
            fmt.Printf(
                "%s-> %sSymlink: %s%s %s-> %s%s%s\n",
                color.DG, color.N, color.GG, targetAlias,
                color.DG, color.GG, targetPath, color.N,
            )
        }

        return errCode
    } else if processedPkg == "" {
        return 1
    }

    return 0
}

// Copyright (c) 2026 Senzdetta