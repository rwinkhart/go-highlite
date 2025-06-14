//go:build stxDart || stxAll

package syntax

func init() {
	syntaxMap["dart"] = &lazySyntax{init: func() []byte {
		return []byte(`filetype: dart
rules:
    - constant.number: "\\b[-+]?([1-9][0-9]*|0[0-7]*|0x[0-9a-fA-F]+)([uU][lL]?|[lL][uU]?)?\\b"
    - constant.number: "\\b[-+]?([0-9]+\\.[0-9]*|[0-9]*\\.[0-9]+)([EePp][+-]?[0-9]+)?[fFlL]?"
    - constant.number: "\\b[-+]?([0-9]+[EePp][+-]?[0-9]+)[fFlL]?"
    - identifier: "[A-Za-z_][A-Za-z0-9_]*[[:space:]]*[(]"
    - statement: "\\b(break|case|catch|continue|default|else|finally)\\b"
    - statement: "\\b(for|function|get|if|in|as|is|new|return|set|switch|final|await|async|sync)\\b"
    - statement: "\\b(switch|this|throw|try|var|void|while|with|import|library|part|const|export)\\b"
    - constant: "\\b(true|false|null)\\b"
    - type: "\\b(List|String)\\b"
    - type: "\\b(int|num|double|bool)\\b"
    - statement: "[-+/*=<>!~%?:&|]"
    - constant: "/[^*]([^/]|(\\\\/))*[^\\\\]/[gim]*"
    - constant: "\\\\[0-7][0-7]?[0-7]?|\\\\x[0-9a-fA-F]+|\\\\[bfnrt'\"\\?\\\\]"
    - comment:
        start: "//"
        end: "$"
        rules:
            - todo: "TODO:?"
    - comment:
        start: "/\\*"
        end: "\\*/"
        rules:
            - todo: "TODO:?"
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
            - constant.specialChar: "\\\\."`)
	}}
}
