function table.slice(tbl, first, last)
  local sliced = {}

  for i = first or 1, last or #tbl, 1 do
    sliced[#sliced + 1] = tbl[i]
  end

  return sliced
end

function table.print(t)
  for i, v in pairs(t) do
    print(i, v)
  end
end

function string.split(s, sep)
  local t = {}
  local pattern = sep == "" and "." or "[^" .. sep .. "]+"
  for part in s:gmatch(pattern) do
    table.insert(t, part)
  end
  return t
end
