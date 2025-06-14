//go:build stxBash || stxAll

package syntax

func init() {
	syntaxMap["bash"] = &lazySyntax{init: func() []byte {
		return []byte(`filetype: shell
rules:
    # Numbers
    - constant.number: "\\b[0-9]+\\b"
      # Conditionals and control flow
    - statement: "\\b(case|do|done|elif|else|esac|exit|fi|for|function|if|in|local|read|return|select|shift|then|time|until|while)\\b"
      # Shell commands
    - type: "\\b(cd|echo|export|let|set|umask|unset)\\b"
      # Common linux commands
    - type: "\\b((g|ig)?awk|bash|dash|find|\\w{0,4}grep|kill|killall|\\w{0,4}less|make|pkill|sed|sh|tar|ping|traceroute|service|dpkg|apt|apt-get|apt-cache|aptitude|dpkg-buildpackage|yum)\\b"
      # Coreutils commands
    - type: "\\b(which|sudo|base64|basename|cat|chcon|chgrp|chmod|chown|chroot|cksum|comm|cp|csplit|cut|date|(l)?dd|df|dir|dircolors|dirname|du|env|expand|expr|factor|false|fmt|fold|head|hostid|id|install|join|link|ln|logname|ls|md5sum|mkdir|mkfifo|mknod|mktemp|mv|nice|nl|nohup|nproc|numfmt|od|paste|pathchk|pinky|pr|printenv|printf|ptx|pwd|readlink|realpath|rm|rmdir|runcon|seq|(sha1|sha224|sha256|sha384|sha512)sum|shred|shuf|sleep|sort|split|stat|stdbuf|stty|sum|sync|tac|tail|tee|test|time|timeout|touch|tr|true|truncate|tsort|tty|uname|unexpand|uniq|unlink|users|vdir|wc|who|whoami|yes)\\b"
    - identifier.var: "(^([[:space:]]+)?[A-Za-z0-9_]+)"
    - special: "(\\{|\\}|\\(|\\)|\\;|\\]|\\[|` + "`" + `|\\\\|\\$|<|>|!|=|&|\\|)"
      # Conditional flags
    - statement: "--[a-z-]+"
    - statement: "\\ -[a-z]+"

    - identifier: "\\$\\{?[0-9A-Z_!@#$*?-]+\\}?"
    - identifier: "\\$\\{?[0-9A-Z_!@#$*?-]+\\}?"
    - symbol.brackets: "(\\[|\\]|\\{|\\}|[()])"

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
        rules:
            - todo: "(TODO|XXX|FIXME):?"`)
	}}
}
