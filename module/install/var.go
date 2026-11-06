// https://github.com/Zeronetsec/blsys

package install

import (
    "regexp"
)

var (
    BackendRegistry = make(map[string]string)
    BackendMacroRegex = regexp.MustCompile(
        `@([a-zA-Z0-9_-]+)\((.*?)\)`,
    )
)

// Copyright (c) 2026 Senzdetta