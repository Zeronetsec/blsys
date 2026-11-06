# https://github.com/Zeronetsec/blsys

function blsys::finsdep () {
    local dep="${__root__}/packages.txt"

    if [[ -n "${1}" ]]; then
        dep="${1}"
    fi

    command blsys \
        --install \
        --mode 'file' \
        --packages "${dep}"

    return ${?}
}

# Copyright (c) 2026 Senzdetta