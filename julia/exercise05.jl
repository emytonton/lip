function success_rate(speed)
    if speed == 0
        return 0.0
    elseif 1 ≤ speed ≤ 4
        return 1.0
    elseif 5 ≤ speed ≤ 8
        return 0.9
    elseif speed == 9
        return 0.8
    elseif speed == 10
        return 0.77
    end
end

function production_rate_per_hour(speed)
    return speed * 221 * success_rate(speed)
end

function working_items_per_minute(speed)
    return floor(Int, production_rate_per_hour(speed) / 60)
end
