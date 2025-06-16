
function rank_range()
    return 1:8
end


function file_range()
    return 'A':'H'
end


function ranks()
    return collect(rank_range())
end


function files()
    return collect(file_range())
end
