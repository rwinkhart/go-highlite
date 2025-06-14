//go:build stxLua || stxAll

package syntax

func init() {
	syntaxMap["lua"] = &lazySyntax{init: func() []byte {
		return []byte(`filetype: lua
rules:
    - statement: "\\b(do|end|while|repeat|until|if|elseif|then|else|for|in|function|local|return)\\b"
    - statement: "\\b(not|and|or)\\b"
    - statement: "\\b(debug|string|math|table|io|coroutine|os|utf8|bit32)\\b\\."
    - statement: "\\b(_ENV|_G|_VERSION|assert|collectgarbage|dofile|error|getfenv|getmetatable|ipairs|load|loadfile|module|next|pairs|pcall|print|rawequal|rawget|rawlen|rawset|require|select|setfenv|setmetatable|tonumber|tostring|type|unpack|xpcall)\\s*\\("
    - identifier: "io\\.\\b(close|flush|input|lines|open|output|popen|read|tmpfile|type|write)\\b"
    - identifier: "math\\.\\b(abs|acos|asin|atan2|atan|ceil|cosh|cos|deg|exp|floor|fmod|frexp|huge|ldexp|log10|log|max|maxinteger|min|mininteger|modf|pi|pow|rad|random|randomseed|sinh|sqrt|tan|tointeger|type|ult)\\b"
    - identifier: "os\\.\\b(clock|date|difftime|execute|exit|getenv|remove|rename|setlocale|time|tmpname)\\b"
    - identifier: "package\\.\\b(config|cpath|loaded|loadlib|path|preload|seeall|searchers|searchpath)\\b"
    - identifier: "string\\.\\b(byte|char|dump|find|format|gmatch|gsub|len|lower|match|pack|packsize|rep|reverse|sub|unpack|upper)\\b"
    - identifier: "table\\.\\b(concat|insert|maxn|move|pack|remove|sort|unpack)\\b"
    - identifier: "utf8\\.\\b(char|charpattern|codes|codepoint|len|offset)\\b"
    - identifier: "coroutine\\.\\b(create|isyieldable|resume|running|status|wrap|yield)\\b"
    - identifier: "debug\\.\\b(debug|getfenv|gethook|getinfo|getlocal|getmetatable|getregistry|getupvalue|getuservalue|setfenv|sethook|setlocal|setmetatable|setupvalue|setuservalue|traceback|upvalueid|upvaluejoin)\\b"
    - identifier: "bit32\\.\\b(arshift|band|bnot|bor|btest|bxor|extract|replace|lrotate|lshift|rrotate|rshift)\\b"
    - identifier: "\\:\\b(close|flush|lines|read|seek|setvbuf|write)\\b"
    - constant: "\\b(false|nil|true)\\b"
    - statement: "(\\b(dofile|require|include)|%q|%!|%Q|%r|%x)\\b"
    - constant.number: "\\b([0-9]+)\\b"
    - symbol: "(\\(|\\)|\\[|\\]|\\{|\\}|\\*\\*|\\*|/|%|\\+|-|\\^|>|>=|<|<=|~=|=|\\.\\.)"
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
    - constant.string:
        start: "\\[\\["
        end: "\\]\\]"
        rules:
            - constant.specialChar: "\\\\."
    - special: "\\\\[0-7][0-7][0-7]|\\\\x[0-9a-fA-F][0-9a-fA-F]|\\\\[abefnrs]|(\\\\c|\\\\C-|\\\\M-|\\\\M-\\\\C-)."
    - comment:
        start: "#"
        end: "$"
        rules: []
    - comment:
        start: "\\-\\-"
        end: "$"
        rules: []
    - comment:
        start: "\\-\\-\\[\\["
        end: "\\]\\]"
        rules: []`)
	}}
}
