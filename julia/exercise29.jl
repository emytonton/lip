using Dates

"""
    schedule_appointment(appointment::String)

Parse an appointment string in the format "7/25/2025 13:45:00" and return
the corresponding `DateTime` object.
"""
function schedule_appointment(appointment::String)
    df = dateformat"m/d/y H:M:S"
    return DateTime(appointment, df)
end

"""
    has_passed(appointment::DateTime)

Check if an appointment `DateTime` is in the past. Returns `true` if the
appointment has passed, `false` otherwise.
"""
function has_passed(appointment::DateTime)
    return now() > appointment
end

"""
    is_afternoon_appointment(appointment::DateTime)

Check if an appointment is in the afternoon (from 12:00 up to 18:00).
Returns `true` if it is, `false` otherwise.
"""
function is_afternoon_appointment(appointment::DateTime)
    appointment_hour = hour(appointment)
    return 12 <= appointment_hour < 18
end

"""
    describe(appointment::DateTime)

Return a human-readable description of the appointment.
Example: "You have an appointment on Thursday, December 5, 2019 at 09:00."
"""
function describe(appointment::DateTime)
    # Manually build the string to ensure correct padding and format.
    day_name = Dates.dayname(appointment)
    month_name = Dates.monthname(appointment)
    day_num = Dates.day(appointment)
    year_num = Dates.year(appointment)
    
    # Use lpad to zero-pad the hour and minute.
    hour_str = lpad(hour(appointment), 2, '0')
    minute_str = lpad(minute(appointment), 2, '0')

    return "You have an appointment on $day_name, $month_name $day_num, $year_num at $hour_str:$minute_str"
end

"""
    anniversary_date()

Return the `Date` of this year's anniversary of the salon's opening.
The salon opened on September 15, 2012.
"""
function anniversary_date()
    return Date(year(today()), 9, 15)
end