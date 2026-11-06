# https://github.com/Zeronetsec/Blsys

__appname__="$(
    command blsys \
        --manifest "${__root__}/manifest.meta" \
        --get 'appname'
)"

__appsource__="$(
    command blsys \
        --manifest "${__root__}/manifest.meta"
        --get 'appsource'
)"

__app_path__="$(command blsys --optdir)/${__appname__}"

# Copyright (c) 2026 Senzdetta