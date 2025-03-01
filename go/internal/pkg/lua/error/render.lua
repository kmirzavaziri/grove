function err(err, values)
    if type(err) ~= "string" or err == "" then
        return "Guard failed"
    end

    return (err:gsub("%${(.-)}", function(key)
        local v = get_nested(values, key)

        if type(v) == "table" then
            return table.concat(v, ", ")
        end

        return tostring(v)
    end))
end

function get_nested(v, kk)
    for k in kk:gmatch("[^.]+") do
        if type(v) ~= "table" then
            return nil
        end

        v = v[k] or v[tonumber(k)]
    end

    return v
end
