// https://github.com/Zeronetsec/blsys

package manifest

import (
    "strings"
)

func extractArrayValues(text string) []string {
    var result []string

    fields := strings.FieldsFunc(
        text, func(r rune) bool {
            return (r == ',' ||
                r == '"' ||
                r == '\'' ||
                r == ' ' ||
                r == '\t')
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