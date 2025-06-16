function get_item(stack, position)
    return stack[position]
end

function set_item!(stack, position, replacement_card)
    stack[position] = replacement_card
    return stack
end

function insert_item_at_top!(stack, new_card)
    push!(stack, new_card)
    return stack
end

function remove_item!(stack, position)
    deleteat!(stack, position)
    return stack
end

function remove_item_from_top!(stack)
    pop!(stack)
    return stack
end

function insert_item_at_bottom!(stack, new_card)
    pushfirst!(stack, new_card)
    return stack
end

function remove_item_at_bottom!(stack)
    popfirst!(stack)
    return stack
end

function check_size_of_stack(stack, stack_size)
    return length(stack) == stack_size
end
