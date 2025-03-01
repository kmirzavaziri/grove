function guard(value)
    if type(value) ~= "string" then
        return err(config.err_type, {config = config, value = value})
    end

    if not string.find(value, config.substring) then
        return err(config.err, {config = config, value = value})
    end

    return ""
end
