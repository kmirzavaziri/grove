function guard(value)
    if type(value) ~= "string" then
        return err(config.err_type, {config = config, value = value})
    end

    if string.sub(value, -#config.suffix) ~= config.suffix then
        return err(config.err, {config = config, value = value})
    end

    return ""
end
