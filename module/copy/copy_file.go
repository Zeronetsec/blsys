// https://github.com/Zeronetsec/blsys

package copy

import (
    "io"
    "os"
    "path/filepath"
)

func copyFile(src, dest string, mode os.FileMode) error {
    in, err := os.Open(src)
    if err != nil {
        return err
    }
    defer in.Close()

    if err := os.MkdirAll(
        filepath.Dir(dest),
        0755,
    ); err != nil {
        return err
    }

    out, err := os.Create(dest)
    if err != nil {
        return err
    }

    if _, err := io.Copy(out, in); err != nil {
        out.Close()
        return err
    }

    if err := out.Close(); err != nil {
        return err
    }

    return os.Chmod(
        dest,
        mode.Perm(),
    )
}

// Copyright (c) 2026 Senzdetta