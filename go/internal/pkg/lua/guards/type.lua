function guard(value)
    if value ~= nil and not is_type(value, config.type_name) then
        return err(config.err, {config = config, value = value})
    end

    return ""
end

function is_type(value, type_name)
    if type_name == "list" then
        return is_list(value)
    end

    if type_name == "map" then
        return is_map(value)
    end

    return type(value) == type_name
end

function is_list(value)
    if type(value) ~= "table" then
        return false
    end

    local count = 0

    for k, _ in pairs(value) do
        count = count + 1

        if type(k) ~= "number" or k ~= count then
            return false
        end
    end

    return true
end

function is_map(value)
    if type(value) ~= "table" then
        return false
    end

    for k, _ in pairs(value) do
        if type(k) ~= "string" then
            return false
        end
    end

    return true
end
