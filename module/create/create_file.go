// https://github.com/Zeronetsec/blsys

package create

import (
    "os"
    "time"
    "path/filepath"
)

func createFile(path string, mode os.FileMode) error {
    dir := filepath.Dir(path)
    if err := os.MkdirAll(dir, 0755); err != nil {
        return err
    }

    if _, err := os.Stat(path); err == nil {
        currentTime := time.Now()

        if err := os.Chtimes(
            path, currentTime,
            currentTime,
        ); err != nil {
            return err
        }

        return os.Chmod(path, mode)
    } else if os.IsNotExist(err) {
        file, err := os.OpenFile(
            path,
            os.O_RDWR|os.O_CREATE|os.O_TRUNC,
            mode,
        )

        if err != nil {
            return err
        }

        return file.Close()
    } else {
        return err
    }
}

// Copyright (c) 2026 Senzdetta