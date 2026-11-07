# https://github.com/Zeronetsec/blsys

function runtime::traperror() {
    local err="${?}"
    local cmd="${BASH_COMMAND}"
    local lineno="${LINENO}"

    echo -e " \x1b[1;90m- \x1b[0merror: \x1b[1;31m${err}\x1b[0m"
    echo -e "\x1b[0mLine: \x1b[0;32m${lineno}\x1b[0m"
    echo -e "\x1b[0m${cmd}\x1b[0m"
}

trap runtime::traperror ERR EXIT

# Copyright (c) 2026 Senzdetta