# https://github.com/Zeronetsec/blsys

function blsys::finsdep () {
    local dep="${__root__}/packages.txt"

    if [[ -n "${1}" ]]; then
        dep="${1}"
    fi

    if [[ ! -f "${dep}" ]]; then
        echo -e "\x1b[1;31m[!] \x1b[0mFile: \x1b[0;32m${dep} \x1b[0mnot found!"
        return 1
    fi

    command blsys \
        --install \
        --mode 'file' \
        --packages "${dep}"

    return ${?}
}

# Copyright (c) 2026 Senzdetta