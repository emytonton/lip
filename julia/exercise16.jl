function get_vector_of_wagons(args...)
    return collect(args)
end

function fix_vector_of_wagons(each_wagons_id, missing_wagons)
    reordered = vcat(each_wagons_id[3:end], each_wagons_id[1:2])
    loc_idx = findfirst(==(1), reordered)
    vcat(
        reordered[1:loc_idx],
        missing_wagons,
        reordered[(loc_idx+1):end]
    )
end

function add_missing_stops(route, stops...)
    new_route = copy(route)
    stops_vec = [pair[2] for pair in stops]
    new_route = merge(new_route, Dict("stops" => stops_vec))
    
    return new_route
end

function extend_route_information(route; more_route_information...)
    return merge(route, Dict(more_route_information))
end