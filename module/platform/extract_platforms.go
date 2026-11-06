// https://github.com/Zeronetsec/Blsys

package platform

import (
    "strings"
)

func extractPlatforms(text string) []string {
    var result []string
    fields := strings.FieldsFunc(
        text, func(r rune) bool {
            return (r == ',' ||
                r == '"' ||
                r == '\'' ||
                r == ' ' ||
                r == '\t' ||
                r == '[' ||
                r == ']')
        },
    )

    for _, f := range fields {
        cleaned := strings.TrimSpace(f)
        if cleaned != "" {
            result = append(result, cleaned)
        }
    }

    return result
}

// Copyright (c) 2026 Senzdetta