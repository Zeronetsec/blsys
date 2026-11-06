// https://github.com/Zeronetsec/blsys

package fsclear

import (
    "os"
)

func clearFile(path string) error {
    info, err := os.Stat(path)
    if err != nil {
        return err
    }

    file, err := os.OpenFile(
        path,
        os.O_WRONLY|os.O_TRUNC,
        info.Mode(),
    )

    if err != nil {
        return err
    }

    defer file.Close()

    if err := os.Chmod(
        path,
        info.Mode(),
    ); err != nil {
        return err
    }

    return nil
}

// Copyright (c) 2026 Senzdetta