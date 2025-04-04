function guard(value)
    if type(value) ~= "table" then
        return err(config.err_type, {config = config, value = value})
    end

    for _, v in ipairs(value) do
        local e = run(config.guard, v)

        if e == "" then
            return ""
        end
    end

    return err(config.err, {config = config, value = value})
end

function run(guard_func, value)
    (_VERSION == "Lua 5.1" and loadstring(guard_func) or load(guard_func))()
    return guard(value)
end
