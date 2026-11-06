// https://github.com/Zeronetsec/Blsys

package platform

import (
    "os"
    "bufio"
    "strings"
)

func parseManifest(filePath string) ([]string, error) {
    file, err := os.Open(filePath)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var platforms []string
    inPlatformBlock := false

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := strings.TrimSpace(scanner.Text())
        if line == "" || strings.HasPrefix(
            line, "//",
        ) || strings.HasPrefix(
            line, "#",
        ) {
            continue
        }

        if strings.HasPrefix(
            line, "platform",
        ) && strings.Contains(line, "[") {
            inPlatformBlock = true
            idx := strings.Index(line, "[")

            remainder := strings.TrimSpace(line[idx+1:])
            if strings.Contains(remainder, "]") {
                inPlatformBlock = false
                remainder = strings.TrimRight(remainder, "]")
                remainder = strings.TrimRight(remainder, ",")
            }

            if remainder != "" {
                extracted := extractPlatforms(remainder)
                platforms = append(platforms, extracted...)
            }

            continue
        }

        if inPlatformBlock {
            if strings.Contains(line, "]") {
                inPlatformBlock = false
                idx := strings.Index(line, "]")

                prefix := strings.TrimSpace(line[:idx])
                if prefix != "" {
                    extracted := extractPlatforms(prefix)
                    platforms = append(platforms, extracted...)
                }

                continue
            }

            extracted := extractPlatforms(line)
            platforms = append(platforms, extracted...)
        }
    }

    if err := scanner.Err(); err != nil {
        return nil, err
    }

    return platforms, nil
}

// Copyright (c) 2026 Senzdetta