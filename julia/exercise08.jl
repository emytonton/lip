
function time_to_mix_juice(juice)
    if juice == "Pure Strawberry Joy"
        return 0.5
    elseif juice == "Energizer" || juice == "Green Garden"
        return 1.5
    elseif juice == "Tropical Island"
        return 3.0
    elseif juice == "All or Nothing"
        return 5.0
    else
        return 2.5
    end
end


function wedges_from_lime(size)
    if size == "small"
        return 6
    elseif size == "medium"
        return 8
    elseif size == "large"
        return 10
    else
        return 0
    end
end


function limes_to_cut(needed, limes)
    limes_cut = 0
    wedges_produced = 0
    
    # Return early if no wedges are needed.
    if needed <= 0
        return 0
    end

    # Loop through available limes until enough wedges are produced.
    for lime_size in limes
        if wedges_produced >= needed
            break # Stop cutting once the need is met.
        end
        wedges_produced += wedges_from_lime(lime_size)
        limes_cut += 1
    end
    
    return limes_cut
end


function order_times(orders)
    # Use a list comprehension for a concise solution.
    return [time_to_mix_juice(juice) for juice in orders]
end



function remaining_orders(time_left, orders)
    time_remaining = Float64(time_left)
    
    # Iterate through orders, subtracting time for each, until time runs out.
    for (index, juice) in enumerate(orders)
        if time_remaining <= 0
            # Time is up, return the rest of the orders from this point.
            return orders[index:end]
        end
        
        mix_time = time_to_mix_juice(juice)
        time_remaining -= mix_time
    end
    
    # If the loop completes, all orders were made in time.
    return String[]

function time_to_mix_juice(juice)
    if juice == "Pure Strawberry Joy"
        return 0.5
    elseif juice == "Energizer" || juice == "Green Garden"
        return 1.5
    elseif juice == "Tropical Island"
        return 3.0
    elseif juice == "All or Nothing"
        return 5.0
    else
        return 2.5
    end
end


function wedges_from_lime(size)
    if size == "small"
        return 6
    elseif size == "medium"
        return 8
    elseif size == "large"
        return 10
    else
        return 0
    end
end

function limes_to_cut(needed, limes)
    limes_cut = 0
    wedges_produced = 0
    
    # Return early if no wedges are needed.
    if needed <= 0
        return 0
    end

    # Loop through available limes until enough wedges are produced.
    for lime_size in limes
        if wedges_produced >= needed
            break # Stop cutting once the need is met.
        end
        wedges_produced += wedges_from_lime(lime_size)
        limes_cut += 1
    end
    
    return limes_cut
end

function order_times(orders)
    # Use a list comprehension for a concise solution.
    return [time_to_mix_juice(juice) for juice in orders]
end


function remaining_orders(time_left, orders)
    time_remaining = Float64(time_left)
    
    # Iterate through orders, subtracting time for each, until time runs out.
    for (index, juice) in enumerate(orders)
        if time_remaining <= 0
            # Time is up, return the rest of the orders from this point.
            return orders[index:end]
        end
        
        mix_time = time_to_mix_juice(juice)
        time_remaining -= mix_time
    end
    
    # If the loop completes, all orders were made in time.
    return String[]
end