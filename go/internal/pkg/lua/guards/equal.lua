function guard(value)
    if type(value) ~= "boolean" then
        return err(config.err_type, {config = config, value = value})
    end

    if value ~= config.value then
        return err(config.err, {config = config, value = value})
    end

    return ""
end
