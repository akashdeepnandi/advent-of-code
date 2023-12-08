require "helper"

--- FILE READ ---
local file = "input"
local lines = {}
for line in io.lines(file) do
  if line ~= "" then
    table.insert(lines, line)
  end
end
--- FILE READ ---

--- SOLUTION ---
local network = {}
local instructions = {}
local startNodes = {}
for i, l in pairs(lines) do
  if i == 1 then
    instructions = string.split(l, "")
  else
    local t = {}
    for m in l:gmatch("(%w+)") do
      table.insert(t, m)
    end
    local node, left, right = table.unpack(t)
    network[node] = { L = left, R = right }

    if string.sub(node, -1) == "A" then
      table.insert(startNodes, node)
    end
  end
end

-- common function to  get steps, variable reach condition
local getSteps = function(startNode, reachedFn)
  local steps = 0
  while startNode ~= "ZZZ" do
    local currentNode = startNode
    for _, nav in pairs(instructions) do -- nav = "L" | "R"
      steps = steps + 1
      local current = network[currentNode]
      currentNode = current[nav]
      if reachedFn(currentNode) then
        return steps
      end
    end
    startNode = currentNode
  end
end

local reachedFn1 = function(currentNode)
  return currentNode == "ZZZ"
end

-- part 1 ends
print("Part 1", getSteps("AAA", reachedFn1))

-- part 2
local reachedFn2 = function(currentNode)
  return string.sub(currentNode, -1) == "Z"
end

local allSteps = {}

for _, startNode in pairs(startNodes) do
  table.insert(allSteps, getSteps(startNode, reachedFn2))
end

local function gcd(a, b)
  if b == 0 then
    return a
  end
  return gcd(b, a % b)
end

local lcm = function(a, b)
  return (a * b) / gcd(a, b)
end


local function lcmN(nums)
  if #nums == 2 then
    return lcm(nums[1], nums[2])
  end
  return lcm(nums[1], lcmN(table.slice(nums, 2)))
end

print("Part 2", lcmN(allSteps))
--- SOLUTION ---
