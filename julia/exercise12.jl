function message(msg)
    split(msg, ": ", limit=2)[2] |> strip
end

function log_level(msg)
    split(msg, "]")[1][2:end] |> lowercase
end

function reformat(msg)
    m = message(msg)
    lvl = log_level(msg)
    "$m ($lvl)"
end
