# Task 1: Define the TreasureChest type
"""
    TreasureChest{T}

A parametric composite type to hold a `password` and some `treasure` of type `T`.
"""
struct TreasureChest{T}
    password::String
    treasure::T
end

# Task 2: Define the get_treasure() function
"""
    get_treasure(password_attempt, chest)

Check if the `password_attempt` matches the `chest`'s password.
If correct, returns the treasure. Otherwise, returns `nothing`.
"""
function get_treasure(password_attempt::String, chest::TreasureChest)
    return password_attempt == chest.password ? chest.treasure : nothing
end

# Task 3: Define the multiply_treasure() function
"""
    multiply_treasure(multiplier, chest)

Creates a new `TreasureChest` with the same password, but with its
treasure multiplied by the `multiplier`.
"""
# Method for when the treasure is a Number
function multiply_treasure(multiplier, chest::TreasureChest{<:Number})
    new_treasure = chest.treasure * multiplier
    return TreasureChest(chest.password, new_treasure)
end

# Method for when the treasure is a String
function multiply_treasure(multiplier::Integer, chest::TreasureChest{<:String})
    # Create a vector containing `multiplier` copies of the original treasure string.
    new_treasure = [chest.treasure for _ in 1:multiplier]
    return TreasureChest(chest.password, new_treasure)
end