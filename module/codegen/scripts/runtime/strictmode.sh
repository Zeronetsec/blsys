# https://github.com/Zeronetsec/blsys

set -o errexit -o errtrace -o pipefail
shopt -s failglob extglob shift_verbose

if ((
    BASH_VERSINFO[0] > 4 ||
    (BASH_VERSINFO[0] == 4 && BASH_VERSINFO[1] >= 4)
)); then
    shopt -s inherit_errexit
fi

# Copyright (c) 2026 Senzdetta