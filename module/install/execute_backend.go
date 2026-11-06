// https://github.com/Zeronetsec/blsys

package install

import (
    "strings"
    "os/exec"
)

func executeBackend(pkg, customBackend string) int {
    if pkg == "" {
        return 1
    }

    if customBackend != "" {
        var fullCmd string
        if strings.Contains(customBackend, "{}") {
            fullCmd = strings.ReplaceAll(
                customBackend, "{}", pkg,
            )
        } else {
            fullCmd = customBackend + " " + pkg
        }

        cmd := exec.Command("bash", "-c", fullCmd)
        if err := cmd.Run(); err != nil {
            if exitErr, ok := err.(*exec.ExitError); ok {
                return exitErr.ExitCode()
            }
            return 1
        }
        return 0
    }

    runCmd := func(name string, args ...string) int {
        cmd := exec.Command(name, args...)
        if err := cmd.Run(); err != nil {
            if exitErr, ok := err.(*exec.ExitError); ok {
                return exitErr.ExitCode()
            }
            return 1
        }
        return 0
    }

    hasCmd := func(name string) bool {
        _, err := exec.LookPath(name)
        return err == nil
    }

    if hasCmd("pacman") {
        if hasCmd("paru") {
            return runCmd(
                "paru", "-S", "--noconfirm", pkg,
            )
        } else if hasCmd("yay") {
            return runCmd(
                "yay", "-S", "--noconfirm", pkg,
            )
        }
        return runCmd(
            "pacman", "-S", "--noconfirm", pkg,
        )
    } else if hasCmd("apt") {
        return runCmd(
            "apt", "install", "-y", pkg,
        )
    } else if hasCmd("dnf") {
        return runCmd(
            "dnf", "install", "-y", pkg,
        )
    } else if hasCmd("zypper") {
        return runCmd(
            "zypper", "in", "-y", pkg,
        )
    } else if hasCmd("apk") {
        return runCmd(
            "apk", "add", pkg,
        )
    } else if hasCmd("flatpak") {
        return runCmd(
            "flatpak", "install", "-y", pkg,
        )
    } else if hasCmd("snap") {
        return runCmd(
            "snap", "install", pkg,
        )
    }

    return 1
}

// Copyright (c) 2026 Senzdetta