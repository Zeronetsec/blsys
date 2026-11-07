# https://github.com/Zeronetsec/blsys

function blsys::move() {
    local input_data="${1}"
    [[ -z "${input_data}" ]] && return 1

    local line
    local rest
    local type_target
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

        if [[ "${line}" == ")" ]]; then
            inside_bracket=0
            continue
        fi

        if (( inside_bracket )); then
            [[ -z "${line}" ]] && continue
            [[ "${line}" =~ ^# ]] && continue

            if [[ "${line}" == *"|"* ]]; then
                type_target="${line%%|*}"
                rest="${line#*|}"
            else
                echo -e "\x1b[1;31m[!] \x1b[0mMissing type: \x1b[0;32m${line}\x1b[0m" >&2
                return 1
            fi

            type_target="${type_target#"${type_target%%[![:space:]]*}"}"
            type_target="${type_target%"${type_target##*[![:space:]]}"}"

            if [[ "${type_target}" != "file" && "${type_target}" != "dir" ]]; then
                echo -e "\x1b[1;31m[!] \x1b[0mInvalid type: \x1b[0;32m${type_target}\x1b[0m" >&2
                return 1
            fi

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

            if [[ -z "${src}" || -z "${dest}" ]]; then
                echo -e "\x1b[1;31m[!] \x1b[0mSource/Dest is empty!" >&2
                return 1
            fi

            args=(--move --type "${type_target}" --source "${src}" --dest "${dest}" --force)
            command blsys "${args[@]}"
        fi
    done <<< "${input_data}"
    return 0
}

# Copyright (c) 2026 Senzdetta