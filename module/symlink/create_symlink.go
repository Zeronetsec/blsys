// https://github.com/Zeronetsec/Blsys

package symlink

import (
    "os"
    "path/filepath"
)

func createSymlink(
    src, dest string,
    force bool,
    hasMode bool,
    mode os.FileMode,
) error {
    if err := os.MkdirAll(
        filepath.Dir(dest),
        0755,
    ); err != nil {
        return err
    }

    if _, err := os.Lstat(dest); err == nil {
        if force {
            if err := os.Remove(dest); err != nil {
                return err
            }
        }
    }

    if err := os.Symlink(src, dest); err != nil {
        return err
    }

    if hasMode {
        _ = os.Chmod(dest, mode)
    }

    return nil
}

// Copyright (c) 2026 Senzdetta