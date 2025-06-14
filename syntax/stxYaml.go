//go:build stxYaml || stxAll

package syntax

func init() {
	syntaxMap["yaml"] = &lazySyntax{init: func() []byte {
		return []byte(`filetype: yaml
rules:
    - type: "(^| )!!(binary|bool|float|int|map|null|omap|seq|set|str) "
    - constant:  "\\b(YES|yes|Y|y|ON|on|NO|no|N|n|OFF|off)\\b"
    - constant: "\\b(true|false)\\b"
    - statement: "(:[[:space:]]|\\[|\\]|:[[:space:]]+[|>]|^[[:space:]]*- )"
    - identifier: "[[:space:]][\\*&][A-Za-z0-9]+"
    - type: "[-.\\w]+:"
    - statement: ":"
    - special:  "(^---|^\\.\\.\\.|^%YAML|^%TAG)"
    - constant.string:
        start: "\""
        end: "\""
        skip: "\\\\."
        rules:
            - constant.specialChar: "\\\\."
    - constant.string:
        start: "'"
        end: "'"
        skip: "\\\\."
        rules:
            - constant.specialChar: "\\\\."
    - comment:
        start: "#"
        end: "$"
        rules: []`)
	}}
}
