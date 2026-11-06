// https://github.com/Zeronetsec/blsys

package manifest

import (
    "bufio"
    "os"
    "strings"
)

func parseManifestFile(filePath string) (
    map[string]string, error,
) {
    file, err := os.Open(filePath)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    result := make(map[string]string)
    var (
        currentGroup string
        inArrayBlock bool
        arrayKey string
        arrayBuffer []string
    )

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := strings.TrimSpace(scanner.Text())

        if line == "" || strings.HasPrefix(
            line, "//",
        ) || strings.HasPrefix(line, "#") {
            continue
        }

        if inArrayBlock {
            if strings.Contains(line, "]") {
                idx := strings.Index(line, "]")
                contentBefore := strings.TrimSpace(line[:idx])

                if contentBefore != "" {
                    arrayBuffer = append(
                        arrayBuffer,
                        extractArrayValues(contentBefore)...,
                    )
                }

                inArrayBlock = false

                result[arrayKey] = strings.Join(
                    arrayBuffer, " ",
                )

                arrayKey = ""
                arrayBuffer = nil
                continue
            }

            arrayBuffer = append(
                arrayBuffer,
                extractArrayValues(line)...,
            )

            continue
        }

        if strings.HasPrefix(line, "group") {
            parts := strings.Fields(line)
            if len(parts) >= 2 {
                groupName := strings.TrimSuffix(
                    parts[1], ":",
                )
                currentGroup = groupName
            }

            if strings.Contains(line, "}") {
                currentGroup = ""
            }

            continue
        }

        if line == "}" {
            currentGroup = ""
            continue
        }

        parts := strings.SplitN(line, ":", 2)
        if len(parts) != 2 {
            continue
        }

        key := strings.TrimSpace(parts[0])
        valPart := strings.TrimSpace(parts[1])

        fullKey := key
        if currentGroup != "" {
            fullKey = currentGroup + "." + key
        }

        if strings.HasPrefix(valPart, "[") {
            inArrayBlock = true
            arrayKey = fullKey
            arrayBuffer = nil

            if strings.HasSuffix(valPart, "]") {
                inner := strings.TrimSuffix(
                    strings.TrimPrefix(
                        valPart, "[",
                    ), "]",
                )

                result[arrayKey] = strings.Join(
                    extractArrayValues(inner), " ",
                )

                inArrayBlock = false
                arrayKey = ""
            } else {
                inner := strings.TrimPrefix(
                    valPart, "[",
                )

                if inner != "" {
                    arrayBuffer = append(
                        arrayBuffer,
                        extractArrayValues(inner)...,
                    )
                }
            }

            continue
        }

        cleanedVal := cleanScalarValue(valPart)
        result[fullKey] = cleanedVal
    }

    if err := scanner.Err(); err != nil {
        return nil, err
    }

    return result, nil
}

// Copyright (c) 2026 Senzdetta