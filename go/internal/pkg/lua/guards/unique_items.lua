function guard(value)
    if type(value) ~= "table" then
        return err(config.err_type, {config = config, value = value})
    end

    local seen = {}

    for _, item in ipairs(value) do
        if seen[item] then
            return err(config.err, {config = config, value = value, item = item})
        end
        seen[item] = true
    end

    return ""
end
