function guard(value)
    if type(value) ~= "number" then
        return err(config.err_type, {config = config, value = value})
    end

    if config.min ~= nil and value < config.min then
        return err(config.err_min, {config = config, value = value})
    end

    if config.max ~= nil and value > config.max then
        return err(config.err_max, {config = config, value = value})
    end

    return ""
end
