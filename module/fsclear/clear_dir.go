// https://github.com/Zeronetsec/blsys

package fsclear

import (
    "os"
)

func clearDir(path string) error {
    info, err := os.Stat(path)
    if err != nil {
        return err
    }

    mode := info.Mode()
    if err := os.RemoveAll(path); err != nil {
        return err
    }

    if err := os.MkdirAll(path, mode); err != nil {
        return err
    }

    return os.Chmod(path, mode)
}

// Copyright (c) 2026 Senzdetta