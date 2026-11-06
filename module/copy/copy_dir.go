// https://github.com/Zeronetsec/Blsys

package copy

import (
    "os"
    "errors"
    "strings"
    "path/filepath"
)

func copyDir(src, dest string) error {
    absSrc, _ := filepath.Abs(src)
    absDest, _ := filepath.Abs(dest)

    sep := string(os.PathSeparator)

    if strings.HasPrefix(absDest+sep, absSrc+sep) {
        return errors.New(
            "destination cannot be inside source",
        )
    }

    return filepath.Walk(
        src, func(
            path string,
            info os.FileInfo,
            err error,
        ) error {
            if err != nil {
                return err
            }

            rel, err := filepath.Rel(src, path)
            if err != nil {
                return err
            }

            target := filepath.Join(dest, rel)

            if info.IsDir() {
                return os.MkdirAll(
                    target,
                    info.Mode().Perm()|0700,
                )
            }

            return copyFile(
                path, target,
                info.Mode(),
            )
        },
    )
}

// Copyright (c) 2026 Senzdetta