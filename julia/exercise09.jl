function shift_back(value, amount)
    return value >>> amount
end

function set_bits(value, mask)
    return value | mask
end

function flip_bits(value, mask)
    return xor(value, mask)
end

function clear_bits(value, mask)
    return value & ~mask
end