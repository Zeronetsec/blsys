#!/usr/bin/env bash

blsys="${HOME}/go/bin/blsys"
blsys_url="github.com/Zeronetsec/blsys@latest"
msgblsys="Updating"

if [[ ! -f "${blsys}" ]]; then
    msgblsys="Installing"
    echo -e "\x1b[1;31m[!] \x1b[0mBuild system: \x1b[0;32m${blsys} \x1b[0mNot found!"
    if ! command -v go > /dev/null 2>&1; then
        echo -e "\x1b[1;31m[!] \x1b[0mGolang is not installed!"
        echo -e "\x1b[1;31m[!] \x1b[0mPlease installed first!"
        exit 1
    fi
fi

export GOPROXY=direct
export GOSUMDB=off

echo -ne "\x1b[1;34m[*] \x1b[0m${msgblsys}: \x1b[0;32m${blsys_url}\x1b[0m"
command go install "${blsys_url}" > /dev/null 2>&1
if [[ "${?}" -eq 0 ]] then
    echo -e " \x1b[1;90m- \x1b[0mexit: \x1b[0;32m${?}\x1b[0m"
else
    echo -e " \x1b[1;90m- \x1b[0merror: \x1b[1;31m${?}\x1b[0m"
    exit 1
fi

unset GOPROXY
unset GOSUMDB
unset blsys_url
unset msgblsys

export PATH="${PATH}:${blsys%/blsys}"; readonly PATH
unset blsys

export __root__="$(
    command blsys \
        --lookpath "${BASH_SOURCE[0]}" \
        --rtt 0
)"; readonly __root__

export manifest="${__root__}/manifest.meta"; readonly manifest

user_platform="$(command blsys --platform)"
echo -ne "\x1b[1;34m[*] \x1b[0mChecking platform support for: \x1b[0;32m${user_platform}\x1b[0m"
if command blsys --platform --support "${manifest}"; then
    echo -e " \x1b[1;90m- \x1b[0mexit: \x1b[0;32m${?}\x1b[0m"
else
    echo -e " \x1b[1;90m- \x1b[0merror: \x1b[1;31m${?}\x1b[0m"
    exit 1
fi

euid="$(command blsys --geteuid)"
if [[ "${user_platform}" != 'android' ]]; then
    if [[ ${euid} -eq 0 ]]; then
        echo -e "\x1b[1;31m[!] \x1b[0mEUID: \x1b[0;32m${euid} \x1b[0mroot required!"
        exit 1
    fi
fi

unset user_platform
unset euid

set -o errexit -o errtrace -o pipefail
shopt -s failglob extglob shift_verbose

if ((
    BASH_VERSINFO[0] > 4 ||
    (BASH_VERSINFO[0] == 4 && BASH_VERSINFO[1] >= 4)
)); then
    shopt -s inherit_errexit
fi

source <(command blsys --codegen 'import')

source "${__root__}/.blsys/installer.sh"
blsys::installer "${@}"
exit 0