# https://github.com/Zeronetsec/blsys

function blsys::prettyfmt() {
    local args="${1:-'<noset>'}"
    local val="${2:-'<noset>'}"

    if [[ "${1}" == '--print-init' ]]; then
        echo -e "\x1b[0mOptions:"
        return 0
    fi

    echo -e "\x1b[1;90m* \x1b[0;32m${args} \x1b[1;90m-> \x1b[0;32m${val}\x1b[0m"
}

# Copyright (c) 2026 Senzdetta