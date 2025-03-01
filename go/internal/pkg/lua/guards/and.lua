function guard(value)
    for _, v in ipairs(config.guards) do
        local e = run(v, value)

        if e ~= "" then
            return err(config.err, {config = config, value = value, nested_err = e})
        end
    end

    return ""
end

function run(guard_func, value)
    (_VERSION == "Lua 5.1" and loadstring(guard_func) or load(guard_func))()
    return guard(value)
end
