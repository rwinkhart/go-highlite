//go:build stxJava || stxAll

package syntax

func init() {
	syntaxMap["java"] = &lazySyntax{init: func() []byte {
		return []byte(`filetype: java
rules:
    - type: "\\b(boolean|byte|char|double|float|int|long|new|short|this|transient|void)\\b"
    - statement: "\\b(break|case|catch|continue|default|do|else|finally|for|if|return|switch|throw|try|while)\\b"
    - type: "\\b(abstract|class|extends|final|implements|import|instanceof|interface|native|package|private|protected|public|static|strictfp|super|synchronized|throws|volatile)\\b"
    - constant: "\\b(true|false|null)\\b"
    - constant.number: "\\b[0-9]+\\b"
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
            - preproc: "..+"
            - constant.specialChar: "\\\\."
    - comment:
        start: "//"
        end: "$"
        rules: []
    - comment:
        start: "/\\*"
        end: "\\*/"
        rules: []`)
	}}
}
