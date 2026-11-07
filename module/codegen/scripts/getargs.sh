# https://github.com/Zeronetsec/blsys

function blsys::getargs() {
    local manifest="${__root__}/manifest.meta"
    local var="${1:-}"

    command blsys --manifest "${manifest}" --get "${var}"
    return ${?}
}

# Copyright (c) 2026 Senzdetta