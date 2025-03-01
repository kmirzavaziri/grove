function guard(value)
    if type(value) ~= "string" then
        return err(config.err_type, {config = config, value = value})
    end

    if not string.match(value, "^[1-9]%d%d%d%d%d%d%d%d%d%d%d$") then
        return err(config.err, {config = config, value = value})
    end
    return ""
end
