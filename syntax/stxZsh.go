//go:build stxZsh || stxAll

package syntax

func init() {
	syntaxMap["zsh"] = &lazySyntax{init: func() []byte {
		return []byte(`filetype: zsh
rules:
    - constant.number: "\\b[0-9]+\\b"
    - statement: "\\b(always|break|bye|case|continue|disown|do|done|elif|else|esac|exit|fi|for|function|if|in|local|read|return|select|shift|then|time|until|while)\\b"
    - statement: "(\\{|\\}|\\(|\\)|\\;|\\]|\\[|` + "`" + `|\\\\|\\$|<|>|!|=|&|\\|)"
    - special: "-[Ldefgruwx]\\b"
    - special: "-(eq|ne|gt|lt|ge|le|s|n|z)\\b"
    - statement: "\\b((un)?alias|bindkey|builtin|cd|declare|eval|exec|export|jobs|let|popd|pushd|set|source|typeset|umask|unset)\\b"
    - type: "\\b(add-zsh-hook|autoload|chdir|compinit|dirs|(dis|en)able|echotc|emulate|print|prompt(init)?|(un)?setopt|zle|zmodload|zstyle|whence)\\b"
    - statement: "\\b((g|ig)?awk|find|\\w{0,4}grep|kill|killall|\\w{0,4}less|make|pkill|sed|tar)\\b"
    - statement: "\\b(base64|basename|cat|chcon|chgrp|chmod|chown|chroot|cksum|comm|cp|csplit|cut|date|dd|df|dir|dircolors|dirname|du|echo|env|expand|expr|factor|false|fmt|fold|head|hostid|id|install|join|link|ln|logname|ls|md5sum|mkdir|mkfifo|mknod|mktemp|mv|nice|nl|nohup|nproc|numfmt|od|paste|pathchk|pinky|pr|printenv|printf|ptx|pwd|readlink|realpath|rm|rmdir|runcon|seq|(sha1|sha224|sha256|sha384|sha512)sum|shred|shuf|sleep|sort|split|stat|stdbuf|stty|sum|sync|tac|tail|tee|test|timeout|touch|tr|true|truncate|tsort|tty|uname|unexpand|uniq|unlink|users|vdir|wc|who|whoami|yes)\\b"
    - identifier: "^\\s+(function\\s+)[0-9A-Z_]+\\s+\\(\\)" # (i)
    - identifier: "\\$\\{?[0-9A-Z_!@#$*?-]+\\}?" #(i)
    - constant.string:
        start: "\""
        end: "\""
        skip: "\\\\."
        rules:
            - constant.specialChar: "\\\\."
    - constant.string:
        start: "'"
        end: "'"
        rules: []
    - comment:
        start: "#"
        end: "$"
        rules: []`)
	}}
}
