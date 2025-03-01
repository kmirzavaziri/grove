function guard(value)
    local e = run(config.guard, value)

    if e == "" then
        return err(config.err, {config = config, value = value})
    end

    return ""
end

function run(guard_func, value)
    (_VERSION == "Lua 5.1" and loadstring(guard_func) or load(guard_func))()
    return guard(value)
end
