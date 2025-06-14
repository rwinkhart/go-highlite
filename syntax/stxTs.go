//go:build stxTs || stxAll

package syntax

func init() {
	syntaxMap["ts"] = &lazySyntax{init: func() []byte {
		return []byte(`filetype: typescript
rules:
    - constant.number: "\\b[-+]?([1-9][0-9]*|0[0-7]*|0x[0-9a-fA-F]+)([uU][lL]?|[lL][uU]?)?\\b"
    - constant.number: "\\b[-+]?([0-9]+\\.[0-9]*|[0-9]*\\.[0-9]+)([EePp][+-]?[0-9]+)?[fFlL]?"
    - constant.number: "\\b[-+]?([0-9]+[EePp][+-]?[0-9]+)[fFlL]?"
    - identifier: "[A-Za-z_][A-Za-z0-9_]*[[:space:]]*[(]"
    - statement: "\\b(abstract|as|async|await|break|case|catch|class|const|constructor|continue)\\b"
    - statement: "\\b(debugger|declare|default|delete|do|else|enum|export|extends|finally|for|from)\\b"
    - statement: "\\b(function|get|if|implements|import|in|instanceof|interface|is|let|module|namespace)\\b"
    - statement: "\\b(new|of|package|private|protected|public|require|return|set|static|super|switch)\\b"
    - statement: "\\b(this|throw|try|type|typeof|var|void|while|with|yield)\\b"
    - constant: "\\b(false|true|null|undefined|NaN)\\b"
    - type: "\\b(Array|Boolean|Date|Enumerator|Error|Function|Math)\\b"
    - type: "\\b(Number|Object|RegExp|String|Symbol)\\b"
    - type: "\\b(any|boolean|never|number|string|symbol)\\b"
    - statement: "[-+/*=<>!~%?:&|]"
    - constant: "/[^*]([^/]|(\\\\/))*[^\\\\]/[gim]*"
    - constant: "\\\\[0-7][0-7]?[0-7]?|\\\\x[0-9a-fA-F]+|\\\\[bfnrt'\"\\?\\\\]"
    - comment:
        start: "//"
        end: "$"
        rules: []
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
