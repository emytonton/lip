function transform(ch)
    if ch == '-'
        return "_"
    elseif isspace(ch)
        return ""
    elseif ch == 'ω'
        return "?"
    elseif isuppercase(ch)
        return "-" * lowercase(string(ch))
    elseif isletter(ch) || ch == '_' || ch == '?' || (!isascii(ch) && isprint(ch))
        return string(ch)
    else
        return ""
    end
end

function clean(str)
    if str == ""
        return ""
    end
    chars = transform.(collect(str))
    join(chars)
end

