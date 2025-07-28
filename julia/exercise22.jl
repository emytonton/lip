function demote(n)
    if n isa Float64
        return UInt8(ceil(n))
    elseif n isa Integer
        return Int8(n)
    else
        throw(MethodError(demote, (n,)))
    end
end

function preprocess(coll)
    if coll isa Vector
        demoted_grades = [demote(n) for n in coll]
        return reverse(demoted_grades)
    elseif coll isa Set
        demoted_grades = [demote(n) for n in coll]
        return sort(demoted_grades, rev=true)
    else
        throw(MethodError(preprocess, (coll,)))
    end
end