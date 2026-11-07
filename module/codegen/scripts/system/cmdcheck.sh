# https://github.com/Zeronetsec/blsys

function system::cmdcheck() {
    local cmd="${1:-}"

    if ! command -v ${cmd} &>/dev/null; then
        echo -e "\x1b[1;31m[!] \x1b[0mFailed to installing \x1b[0;32m${cmd}\x1b[0m"
        return 1
    fi

    echo -e "\x1b[0;32m[+] \x1b[0m${cmd^} installed!"
    return 0
}

# Copyright (c) 2026 Senzdetta