# https://github.com/Zeronetsec/blsys

function blsys::auto_prepare_path() {
    local -a target_path=(
        "$(command blsys --tmpdir)"
        "$(command blsys --optdir)"
        "$(command blsys --bindir)"
        "${HOME}/.local/share"
    )

    local i
    for i in "${target_path[@]}"; do
        if [[ ! -d "${target_path}" ]]; then
            command blsys \
                --create \
                --type 'dir' \
                --target "${i}" \
                --mode 0755
        fi
    done
}

# Copyright (c) 2026 Senzdetta