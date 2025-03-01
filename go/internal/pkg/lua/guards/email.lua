function guard(value)
    if type(value) ~= "string" then
        return err(config.err_type, {config = config, value = value})
    end

    local pattern = "^[A-Za-z0-9%.%%_]+@[A-Za-z0-9%%_%.]+%.[A-Za-z]+$"

    if not string.match(value, pattern) then
        return err(config.err, {config = config, value = value})
    end

    return ""
end
