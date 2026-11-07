# https://github.com/Zeronetsec/blsys

function blsys::getargs() {
    local local_manifest="${__root__}/manifest.meta"
    local var="${1:-}"

    command blsys --manifest "${local_manifest}" --get "${var}"
    return ${?}
}

# Copyright (c) 2026 Senzdetta