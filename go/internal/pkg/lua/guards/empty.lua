function guard(value)
    if value ~= nil and value ~= "" and (type(value) ~= "table" or next(value) ~= nil) then
        return err(config.err, {config = config, value = value})
    end
    return ""
end
