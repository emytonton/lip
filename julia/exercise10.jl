function create_inventory(items)
    inventory = Dict{String, Int}()
    for item in items
        inventory[item] = get(inventory, item, 0) + 1
    end
    return inventory
end

function add_items(inventory, items)
    for item in items
        inventory[item] = get(inventory, item, 0) + 1
    end
    return inventory
end

function decrement_items(inventory, items)
    for item in items
        if haskey(inventory, item) && inventory[item] > 0
            inventory[item] -= 1
        end
    end
    return inventory
end

function remove_item(inventory, item)
    delete!(inventory, item)
    return inventory
end

function list_inventory(inventory)
    pairs_list = Pair{String, Int}[]
    for (key, value) in inventory
        if value > 0
            push!(pairs_list, key => value)
        end
    end
    sort!(pairs_list, by = p -> p.first)
    return pairs_list
end