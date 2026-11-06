// https://github.com/Zeronetsec/blsys

package geteuid

import (
    "strconv"
    "os/user"
)

func getSystemEUID() int {
    current, err := user.Current()
    if err != nil {
        return 0
    }

    uidInt, err := strconv.Atoi(current.Uid)
    if err != nil {
        return 0
    }

    return uidInt
}