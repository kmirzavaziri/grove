function guard(value)
    if type(value) ~= "string" then
        return err(config.err_type, {config = config, value = value})
    end

    if string.sub(value, 1, #config.prefix) ~= config.prefix then
        return err(config.err, {config = config, value = value})
    end

    return ""
end
