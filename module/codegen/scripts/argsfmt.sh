# https://github.com/Zeronetsec/blsys

function blsys::argsfmt() {
    local args="${1:-'<noset>'}"
    local val="${2:-'<noset>'}"

    if [[ "${1}" == '--print-init' ]]; then
        printf '\n'
        echo -e "\x1b[0mOptions:"
        return 0
    fi

    if [[ "${1}" == '--print-endline' ]]; then
        printf '\n'
        return 0
    fi

    echo -e "\x1b[1;90m* \x1b[0;32m${args} \x1b[1;90m-> \x1b[0;32m${val}\x1b[0m"
}

# Copyright (c) 2026 Senzdetta