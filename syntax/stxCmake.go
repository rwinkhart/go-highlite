//go:build stxCmake || stxAll

package syntax

func init() {
	syntaxMap["cmake"] = &lazySyntax{init: func() []byte {
		return []byte(`filetype: cmake
rules:
    - identifier.var: "^[[:space:]]*[A-Z0-9_]+"
    - preproc: "^[[:space:]]*(include|include_directories|include_external_msproject)\\b"
    - statement: "^[[:space:]]*\\b((else|end)?if|else|(end)?while|(end)?foreach|break)\\b"
    - statement: "\\b(COPY|NOT|COMMAND|PROPERTY|POLICY|TARGET|EXISTS|IS_(DIRECTORY|ABSOLUTE)|DEFINED)\\b[[:space:]]"
    - statement: "[[:space:]]\\b(OR|AND|IS_NEWER_THAN|MATCHES|(STR|VERSION_)?(LESS|GREATER|EQUAL))\\b[[:space:]]"
    - special: "^[[:space:]]*\\b((end)?(function|macro)|return)"
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
    - preproc:
        start: "\\$(\\{|ENV\\{)"
        end: "\\}"
        rules: []
    - identifier.macro: "\\b(APPLE|UNIX|WIN32|CYGWIN|BORLAND|MINGW|MSVC(_IDE|60|71|80|90)?)\\b"
    - comment:
        start: "#"
        end: "$"
        rules:
            - todo: "(TODO|XXX|FIXME):?"`)
	}}
}
