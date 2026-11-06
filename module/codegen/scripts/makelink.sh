# https://github.com/Zeronetsec/blsys

function blsys::makelink() {
    local input_data="${1}"
    [[ -z "${input_data}" ]] && return 1

    local line
    local rest
    local mode
    local src
    local dest
    local inside_bracket=0
    local -a args

    while read -r line; do
        line="${line#"${line%%[![:space:]]*}"}"
        line="${line%"${line##*[![:space:]]}"}"

        if [[ "${line}" == "("* ]]; then
            inside_bracket=1
            continue
        fi

        if [[ "${line}" == ")"* ]]; then
            inside_bracket=0
            continue
        fi

        if (( inside_bracket )); then
            [[ -z "${line}" ]] && continue
            [[ "${line}" =~ ^# ]] && continue

            if [[ "${line}" == *"|"* ]]; then
                mode="${line%%|*}"
                rest="${line#*|}"
            else
                mode="default"
                rest="${line}"
            fi

            mode="${mode#"${mode%%[![:space:]]*}"}"
            mode="${mode%"${mode##*[![:space:]]}"}"
            [[ -z "${mode}" ]] && mode="default"

            if [[ "${rest}" != *"->"* ]]; then
                echo -e "\x1b[1;31m[!] \x1b[0mInvalid syntax: \x1b[0;32m${line}\x1b[0m" >&2
                return 1
            fi

            src="${rest%%->*}"
            dest="${rest#*->}"

            src="${src#"${src%%[![:space:]]*}"}"
            src="${src%"${src##*[![:space:]]}"}"

            dest="${dest#"${dest%%[![:space:]]*}"}"
            dest="${dest%"${dest##*[![:space:]]}"}"

            if [[ "${mode}" != "default" && ! "${mode}" =~ ^[0-7]{3,4}$ ]]; then
                echo -e "\x1b[1;31m[!] \x1b[0mInvalid mode: \x1b[0;32m${mode} \x1b[0mfor \x1b[0;32m${src}\x1b[0m" >&2
                return 1
            fi

            if [[ -z "${src}" || -z "${dest}" ]]; then
                echo -e "\x1b[1;31m[!] \x1b[0mSource/Dest is empty!" >&2
                return 1
            fi

            args=(--symlink --source "${src}" --dest "${dest}")
            [[ "${mode}" != "default" ]] && args+=(--mode "${mode}")
            args+=(--force)

            command blsys "${args[@]}"
        fi
    done <<< "${input_data}"
    return 0
}

# Copyright (c) 2026 Senzdetta