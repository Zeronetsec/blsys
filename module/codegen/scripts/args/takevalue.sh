# https://github.com/Zeronetsec/blsys

function args::takevalue() {
    local -n __out__="${1}"

    if [[ -n "${3:-}" && "${3}" != --* ]]; then
        __out__="${3}"
        return 0
    fi

    return 1
}

# Copyright (c) 2026 Senzdetta