// https://github.com/Zeronetsec/Blsys

package install

import (
    "os"
    "fmt"
    "strings"
    "bufio"
    "github.com/Zeronetsec/Blsys/utils/color"
)

func processInstallFile(filePath, binDir string) error {
    file, err := os.Open(filePath)
    if err != nil {
        return err
    }
    defer file.Close()

    var packageLines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := strings.TrimSpace(scanner.Text())
        if line == "" || strings.HasPrefix(
            line, "#",
        ) || strings.HasPrefix(
            line, "//",
        ) {
            continue
        }

        if strings.Contains(
            line, "=",
        ) && !strings.Contains(
            line, "->",
        ) {
            parts := strings.SplitN(line, "=", 2)

            alias := strings.TrimSpace(parts[0])
            cmdRaw := strings.TrimSpace(parts[1])

            cmdRaw = strings.Trim(cmdRaw, "\"'")
            BackendRegistry[alias] = cmdRaw

            continue
        }

        packageLines = append(packageLines, line)
    }

    if err := scanner.Err(); err != nil {
        return err
    }

    for _, line := range packageLines {
        fmt.Printf(
            "%s[*] %sInstalling: %s%s%s\n",
            color.B, color.N, color.GG, line, color.N,
        )
        parseAndExecuteItem(line, "", binDir)
    }

    return nil
}

// Copyright (c) 2026 Senzdetta