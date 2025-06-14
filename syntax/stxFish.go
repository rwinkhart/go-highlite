//go:build stxFish || stxAll

package syntax

func init() {
	syntaxMap["fish"] = &lazySyntax{init: func() []byte {
		return []byte(`filetype: fish
rules:
    - constant: "\\b[0-9]+\\b"
    - statement: "\\b(and|begin|break|case|continue|else|end|for|function|if|in|not|or|return|select|shift|switch|while)\\b"
    - special: "(\\{|\\}|\\(|\\)|\\;|\\]|\\[|` + "`" + `|\\\\|\\$|<|>|^|!|=|&|\\|)"
    - type: "\\b(bg|bind|block|breakpoint|builtin|cd|count|command|commandline|complete|dirh|dirs|echo|emit|eval|exec|exit|fg|fish|fish_config|fish_ident|fish_pager|fish_prompt|fish_right_prompt|fish_update_completions|fishd|funced|funcsave|functions|help|history|jobs|math|mimedb|nextd|open|popd|prevd|psub|pushd|pwd|random|read|set|set_color|source|status|string|trap|type|ulimit|umask|vared)\\b"
    - type: "\\b((g|ig)?awk|bash|dash|find|\\w{0,4}grep|kill|killall|\\w{0,4}less|make|pkill|sed|sh|tar)\\b"
    - type: "\\b(base64|basename|cat|chcon|chgrp|chmod|chown|chroot|cksum|comm|cp|csplit|cut|date|dd|df|dir|dircolors|dirname|du|env|expand|expr|factor|false|fmt|fold|head|hostid|id|install|join|link|ln|logname|ls|md5sum|mkdir|mkfifo|mknod|mktemp|mv|nice|nl|nohup|nproc|numfmt|od|paste|pathchk|pinky|pr|printenv|printf|ptx|pwd|readlink|realpath|rm|rmdir|runcon|seq|(sha1|sha224|sha256|sha384|sha512)sum|shred|shuf|sleep|sort|split|stat|stdbuf|stty|sum|sync|tac|tail|tee|test|time|timeout|touch|tr|true|truncate|tsort|tty|uname|unexpand|uniq|unlink|users|vdir|wc|who|whoami|yes)\\b"
    - statement: "--[a-z-]+"
    - statement: "\\ -[a-z]+"
    - identifier: "(?i)\\$\\{?[0-9A-Z_!@#$*?-]+\\}?"
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
        rules: []
    - comment:
        start: "#"
        end: "$"
        rules:
            - todo: "(TODO|XXX|FIXME):?"`)
	}}
}
