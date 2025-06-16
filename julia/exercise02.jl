const expected_bake_time = 60

function preparation_time(layers)
    return layers * 2
end

function remaining_time(time_in_oven)
    return expected_bake_time - time_in_oven
end

function total_working_time(layers, time_in_oven)
    return preparation_time(layers) + time_in_oven
end

