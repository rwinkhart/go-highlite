//go:build stxFortran || stxAll

package syntax

func init() {
	syntaxMap["fortran"] = &lazySyntax{init: func() []byte {
		return []byte(`filetype: fortran
rules:
    - type:  "(?i)\\b(action|advance|all|allocatable|allocated|any|apostrophe)\\b"
    - type:  "(?i)\\b(append|asis|assign|assignment|associated|character|common)\\b"
    - type:  "(?i)\\b(complex|data|default|delim|dimension|double precision)\\b"
    - type:  "(?i)\\b(elemental|epsilon|external|file|fmt|form|format|huge)\\b"
    - type:  "(?i)\\b(implicit|include|index|inquire|integer|intent|interface)\\b"
    - type:  "(?i)\\b(intrinsic|iostat|kind|logical|module|none|null|only)\\\\b"
    - type:  "(?i)\\b(operator|optional|pack|parameter|pointer|position|private)\\b"
    - type:  "(?i)\\b(program|public|real|recl|recursive|selected_int_kind)\\b"
    - type:  "(?i)\\b(selected_real_kind|subroutine|status)\\b"
    - constant:  "(?i)\\b(abs|achar|adjustl|adjustr|allocate|bit_size|call|char)\\b"
    - constant:  "(?i)\\b(close|contains|count|cpu_time|cshift|date_and_time)\\b"
    - constant:  "(?i)\\b(deallocate|digits|dot_product|eor|eoshift|function|iachar)\\b"
    - constant:  "(?i)\\b(iand|ibclr|ibits|ibset|ichar|ieor|iolength|ior|ishft|ishftc)\\b"
    - constant:  "(?i)\\b(lbound|len|len_trim|matmul|maxexponent|maxloc|maxval|merge)\\b"
    - constant:  "(?i)\\b(minexponent|minloc|minval|mvbits|namelist|nearest|nullify)\\b"
    - constant:  "(?i)\\b(open|pad|present|print|product|pure|quote|radix)\\b"
    - constant:  "(?i)\\b(random_number|random_seed|range|read|readwrite|replace)\\b"
    - constant:  "(?i)\\b(reshape|rewind|save|scan|sequence|shape|sign|size|spacing)\\b"
    - constant:  "(?i)\\b(spread|sum|system_clock|target|transfer|transpose|trim)\\b"
    - constant:  "(?i)\\b(ubound|unpack|verify|write|tiny|type|use|yes)\\b"
    - statement: "(?i)\\b(.and.|case|do|else|else?if|else?where|end|end?do|end?if)\\b"
    - statement: "(?i)\\b(end?select|.eqv.|forall|if|lge|lgt|lle|llt|.neqv.|.not.)\\b"
    - statement: "(?i)\\b(.or.|repeat|select case|then|where|while)\\b"
    - special: "(?i)\\b(continue|cycle|exit|go?to|result|return)\\b"
    - symbol.operator: "[.:;,+*|=!\\%]|/|-|&"
    - symbol.brackets: "[(){}]|\\[|\\]"
    - preproc: "^[[:space:]]*#[[:space:]]*(define|include|(un|ifn?)def|endif|el(if|se)|if|warning|error)"
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
        start: "!"
        end: "$"
        rules:
            - todo: "(TODO|XXX|FIXME):?"`)
	}}
}
