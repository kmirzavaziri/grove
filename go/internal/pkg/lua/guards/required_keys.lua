function guard(value)
    if type(value) ~= "table" then
        return err(config.err_type, {config = config, value = value})
    end

    for _, key in ipairs(config.keys) do
        if value[key] == nil then
            return err(config.err, {config = config, value = value, key = key})
        end
    end

    return ""
end
