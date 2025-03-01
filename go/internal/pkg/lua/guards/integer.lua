function guard(value)
    if type(value) ~= "number" then
        return err(config.err_type, {config = config, value = value})
    end

    if value % 1 ~= 0 then
        return err(config.err, {config = config, value = value})
    end

    return ""
end
