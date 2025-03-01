function guard(value)
    for _, v in ipairs(config.guards) do
        local e = run(v, value)

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
