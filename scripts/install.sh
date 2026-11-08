#!/usr/bin/env bash

blsys="${HOME}/go/bin/blsys"
blsys_url="github.com/Zeronetsec/blsys@latest"
msgblsys="Updating"

if [[ ! -f "${blsys}" ]]; then
    msgblsys="Installing"
    echo -e "\x1b[0;33m[!] \x1b[0mBuild system: \x1b[0;32m${blsys} \x1b[0mNot found!" >&2
    if ! command -v go > /dev/null 2>&1; then
        echo -e "\x1b[1;31m[!] \x1b[0mGolang is not installed!" >&2
        echo -e "\x1b[1;31m[!] \x1b[0mPlease installed first!" >&2
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
echo -ne "\x1b[1;34m[*] \x1b[0mChecking platform support: \x1b[0;32m${user_platform}\x1b[0m"
if command blsys --platform --support "${manifest}"; then
    echo -e " \x1b[1;90m- \x1b[0mexit: \x1b[0;32m${?}\x1b[0m"
else
    echo -e " \x1b[1;90m- \x1b[0merror: \x1b[1;31m${?}\x1b[0m"
    echo -e "\x1b[1;31m[!] \x1b[0mUnsupported platform: \x1b[0;32m${user_platform}\x1b[0m" >&2
    exit 1
fi

euid="$(command blsys --geteuid)"
if [[ "${user_platform}" != 'android' ]]; then
    if [[ ${euid} -eq 0 ]]; then
        echo -e "\x1b[1;31m[!] \x1b[0mEUID: \x1b[0;32m${euid} \x1b[0mroot required!" >&2
        exit 1
    fi
fi

unset user_platform
unset euid

if ((
    BASH_VERSINFO[0] < 4 ||
    (BASH_VERSINFO[0] == 4 && BASH_VERSINFO[1] < 4)
)); then
    echo -e "\x1b[0;33m[!] \x1b[0mWarning bash version: \x1b[0;31m${BASH_VERSION} \x1b[1;90m< \x1b[0;32mv4.4\x1b[0m" >&2
fi

source <(command blsys --codegen 'import')

echo -e "\x1b[1;34m[*] \x1b[0mCall: \x1b[0;32mblsys::installer\x1b[0m"
source "${__root__}/.blsys/installer.sh"
blsys::installer "${@}"
exit 0