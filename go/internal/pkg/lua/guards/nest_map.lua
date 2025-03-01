function guard(value)
    if type(value) ~= "table" then
        return err(config.err_type, {config = config, value = value})
    end

    local e = run(config.guard, value[config.key])

    if e ~= "" then
        return err(config.err, {config = config, value = value, nested_err = e})
    end

    return ""
end

function run(guard_func, value)
    (_VERSION == "Lua 5.1" and loadstring(guard_func) or load(guard_func))()
    return guard(value)
end
