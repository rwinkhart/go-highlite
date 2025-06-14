//go:build stxPascal || stxAll

package syntax

func init() {
	syntaxMap["pascal"] = &lazySyntax{init: func() []byte {
		return []byte(`filetype: pascal
rules:
    - type: "\\b(?i:(string|ansistring|widestring|shortstring|char|ansichar|widechar|boolean|byte|shortint|word|smallint|longword|cardinal|longint|integer|int64|single|currency|double|extended))\\b"
    - statement: "\\b(?i:(and|asm|array|begin|break|case|const|constructor|continue|destructor|div|do|downto|else|end|file|for|function|goto|if|implementation|in|inline|interface|label|mod|not|object|of|on|operator|or|packed|procedure|program|record|repeat|resourcestring|set|shl|shr|then|to|type|unit|until|uses|var|while|with|xor))\\b"
    - statement: "\\b(?i:(as|class|dispose|except|exit|exports|finalization|finally|inherited|initialization|is|library|new|on|out|property|raise|self|threadvar|try))\\b"
    - statement: "\\b(?i:(absolute|abstract|alias|assembler|cdecl|cppdecl|default|export|external|forward|generic|index|local|name|nostackframe|oldfpccall|override|pascal|private|protected|public|published|read|register|reintroduce|safecall|softfloat|specialize|stdcall|virtual|write))\\b"
    - constant: "\\b(?i:(false|true|nil))\\b"
    - special:
        start: "asm"
        end: "end"
        rules: []
    - constant.number: "\\$[0-9A-Fa-f]+"
    - constant.number: "\\b[+-]?[0-9]+([.]?[0-9]+)?(?i:e[+-]?[0-9]+)?"
    - constant.string:
        start: "#[0-9]{1,}"
        end: "$"
        rules:
            - constant.specialChar: "\\\\."
    - constant.string:
        start: "'"
        end: "'"
        skip: "\\\\."
        rules:
            - constant.specialChar: "\\\\."
    - preproc:
        start: "{\\$"
        end: "}"
        rules: []
    - comment:
        start: "//"
        end: "$"
        rules: []
    - comment:
        start: "\\(\\*"
        end: "\\*\\)"
        rules: []
    - comment:
        start: "({)(?:[^$])"
        end: "}"
        rules: []`)
	}}
}
