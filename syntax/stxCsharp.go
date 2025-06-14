//go:build stxCsharp || stxAll

package syntax

func init() {
	syntaxMap["csharp"] = &lazySyntax{init: func() []byte {
		return []byte(`filetype: csharp
rules:
    - identifier.macro: "class +[A-Za-z0-9]+ *((:) +[A-Za-z0-9.]+)?"
    - identifier.var: "@[A-Za-z]+"
    - identifier: "[A-Za-z_][A-Za-z0-9_]*[[:space:]]*[()]"
    - type: "\\b(bool|byte|sbyte|char|decimal|double|float|IntPtr|int|uint|long|ulong|object|short|ushort|string|base|this|var|void)\\b"
    - statement: "\\b(alias|as|case|catch|checked|default|do|dynamic|else|finally|fixed|for|foreach|goto|if|is|lock|new|null|return|switch|throw|try|unchecked|while)\\b"
    - statement: "\\b(abstract|async|class|const|delegate|enum|event|explicit|extern|get|implicit|in|internal|interface|namespace|operator|out|override|params|partial|private|protected|public|readonly|ref|sealed|set|sizeof|stackalloc|static|struct|typeof|unsafe|using|value|virtual|volatile|yield)\\b"
    - statement: "\\b(from|where|select|group|info|orderby|join|let|in|on|equals|by|ascending|descending)\\b"
    - special: "\\b(break|continue)\\b"
    - constant.bool: "\\b(true|false)\\b"
    - symbol.operator: "[\\-+/*=<>?:!~%&|]"
    - constant.number: "\\b([0-9._]+|0x[A-Fa-f0-9_]+|0b[0-1_]+)[FL]?\\b"
    - constant.string:
        start: "\""
        end: "\""
        skip: "\\\\."
        rules:
            - constant.specialChar: "\\\\([btnfr]|'|\\\"|\\\\)"
            - constant.specialChar: "\\\\u[A-Fa-f0-9]{4}"
    - constant.string:
        start: "'"
        end: "'"
        skip: "\\\\."
        rules:
            - constant.specialChar: "\\\\([btnfr]|'|\\\"|\\\\)"
            - constant.specialChar: "\\\\u[A-Fa-f0-9]{4}"
    - comment:
        start: "//"
        end: "$"
        rules:
            - todo: "(TODO|XXX|FIXME):?"
    - comment:
        start: "/\\*"
        end: "\\*/"
        rules:
            - todo: "(TODO|XXX|FIXME):?"`)
	}}
}
