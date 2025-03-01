function guard(value)
    if type(value) ~= "string" then
        return err(config.err_type, {config = config, value = value})
    end

    for _, enum_value in ipairs(config.values) do
        if value == enum_value then
            return ""
        end
    end

    return err(config.err, {config = config, value = value})
end
