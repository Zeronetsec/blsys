// https://github.com/Zeronetsec/blsys

package move

import (
    "os"
    "path/filepath"
)

func moveItem(src, dest string, force bool) error {
    if err := os.MkdirAll(
        filepath.Dir(dest),
        0755,
    ); err != nil {
        return err
    }

    if force {
        if _, err := os.Stat(dest); err == nil {
            if err := os.RemoveAll(dest); err != nil {
                return err
            }
        }
    }

    err := os.Rename(src, dest)
    if err != nil {
        return err
    }

    return nil
}

// Copyright (c) 2026 Senzdetta