# https://github.com/Zeronetsec/blsys

function blsys::makedir() {
    local input_data="${1}"
    if [[ -z "${input_data}" ]]; then
        return 1
    fi

    local default_mode="0755"
    local line
    local mode
    local target
    local inside_bracket=0

    while read -r line; do
        line="${line#"${line%%[![:space:]]*}"}"
        line="${line%"${line##*[![:space:]]}"}"

        if [[ "${line}" == *"("* ]]; then
            inside_bracket=1
            continue
        fi

        if [[ "${line}" == *")"* ]]; then
            inside_bracket=0
            continue
        fi

        if (( inside_bracket )); then
            [[ -z "${line}" ]] && continue
            [[ "${line}" =~ ^# ]] && continue

            if [[ "${line}" == *"|"* ]]; then
                mode="${line%%|*}"
                target="${line#*|}"
            else
                mode="default"
                target="${line}"
            fi

            mode="${mode#"${mode%%[![:space:]]*}"}"
            mode="${mode%"${mode##*[![:space:]]}"}"

            target="${target#"${target%%[![:space:]]*}"}"
            target="${target%"${target##*[![:space:]]}"}"

            [[ -z "${mode}" || "${mode}" == "default" ]] && {
                mode="${default_mode}"
            }

            if [[ ! "${mode}" =~ ^[0-7]{3,4}$ ]]; then
                echo -e "\x1b[1;31m[!] \x1b[0mInvalid mode: \x1b[0;32m${mode} \x1b[0mfor \x1b[0;32m${target}\x1b[0m" >&2
                return 1
            fi

            [[ -z "${target}" ]] && continue

            command blsys \
                --create \
                --type 'dir' \
                --target "${target}" \
                --mode "${mode}" \
                --force
        fi
    done <<< "${input_data}"
    return 0
}

# Copyright (c) 2026 Senzdetta