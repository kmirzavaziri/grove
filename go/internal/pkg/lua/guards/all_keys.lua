function guard(value)
    if type(value) ~= "table" then
        return err(config.err_type, {config = config, value = value})
    end

    for key, _ in pairs(value) do
        if not string.match(key, config.pattern) then
            return err(config.err, {config = config, value = value})
        end
    end

    return ""
end
