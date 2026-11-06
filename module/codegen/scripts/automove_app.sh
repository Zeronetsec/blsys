# https://github.com/Zeronetsec/blsys

function blsys::automove_app() {
    local appname="$(
        command blsys \
            --manifest "${__root__}/manifest.meta" \
            --get 'appname'
    )"

    local opt="$(command blsys --optdir)"

    if [[ -d "${opt}/${appname}" ]]; then
        command blsys --remove "${opt}/${appname}"
    fi

    command blsys \
        --move \
        --type 'dir' \
        --source "${__root__}" \
        --dest "${opt}/${appname}"
}

# Copyright (c) 2026 Senzdetta