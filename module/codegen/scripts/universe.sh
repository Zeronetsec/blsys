# https://github.com/Zeronetsec/blsys

source <(command blsys --codegen 'app/appvar')
source <(command blsys --codegen 'app/automove_app')

source <(command blsys --codegen 'config/getmanifest')

source <(command blsys --codegen 'format/argsfmt')

source <(command blsys --codegen 'iofs/chmod')
source <(command blsys --codegen 'iofs/copy')
source <(command blsys --codegen 'iofs/makedir')
source <(command blsys --codegen 'iofs/makefile')
source <(command blsys --codegen 'iofs/makelink')
source <(command blsys --codegen 'iofs/move')

source <(command blsys --codegen 'package/finsdep')

source <(command blsys --codegen 'path/ensure_path')
source <(command blsys --codegen 'path/pathvar')

source <(command blsys --codegen 'runtime/strictmode')
source <(command blsys --codegen 'runtime/traperror')
source <(command blsys --codegen 'runtime/unbound_error')

source <(command blsys --codegen 'system/cmdcheck')

source <(command blsys --codegen 'visual/color')

# Copyright (c) 2026 Senzdetta