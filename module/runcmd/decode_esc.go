// https://github.com/Zeronetsec/blsys

import (
    "strconv"
    "strings"
)

func decodeESC(s string) string {
    var b strings.Builder
    b.Grow(len(s))

    isHex := func(c byte) bool {
        return (c >= '0' && c <= '9') ||
            (c >= 'a' && c <= 'f') ||
            (c >= 'A' && c <= 'F')
    }

    for i := 0; i < len(s); i++ {
        if s[i] != '\\' || i+1 >= len(s) {
            b.WriteByte(s[i])
            continue
        }

        i++
        switch s[i] {
            case 'e':
                b.WriteByte(0x1b)
            case 'n':
                b.WriteByte('\n')
            case 't':
                b.WriteByte('\t')
            case 'r':
                b.WriteByte('\r')
            case '\\':
                b.WriteByte('\\')
            case 'x':
                j := i + 1

                for j < len(s) && j < i+3 && isHex(s[j]) {
                    j++
                }

                if j == i+1 {
                    b.WriteString(`\x`)
                    break
                }

                v, _ := strconv.ParseUint(s[i+1:j], 16, 8)
                b.WriteByte(byte(v))

                i = j - 1
            case '0', '1', '2', '3', '4', '5', '6', '7':
                j := i

                for (j < len(s) &&
                    j < i+3 &&
                    s[j] >= '0' &&
                    s[j] <= '7') {
                        j++
                }

                v, _ := strconv.ParseUint(s[i:j], 8, 16)
                b.WriteByte(byte(v))

                i = j - 1
            default:
                b.WriteByte('\\')
                b.WriteByte(s[i])
        }
    }

    return b.String()
}

// Copyright (c) 2026 Senzdetta