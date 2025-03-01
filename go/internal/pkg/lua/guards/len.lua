function guard(value)
    local length = 0
    if type(value) ~= "string" and type(value) ~= "table" then
        return err(config.err_type, {config = config, value = value})
    end

    length = #value

    if config.min ~= nil and length < config.min then
        return err(config.err_min, {config = config, value = value})
    end

    if config.max ~= nil and length > config.max then
        return err(config.err_max, {config = config, value = value})
    end

    return ""
end
