import fs from "fs";
const file = fs.readFileSync(process.argv[2], { encoding: "utf8" }).trim();
const lines = file.split("\n");

let graph = {};
let nodes = [];

lines.forEach((line, lineIdx) => {
  // logic
  const parts = line.split(" ");
  let u = parts[0];
  let v = parts[2];
  let d = parseInt(parts[parts.length - 1]);
  if (u in graph === false) {
    graph[u] = {};
  }
  if (v in graph === false) {
    graph[v] = {};
  }
  graph[u][v] = d;
  graph[v][u] = d;

  if (!nodes.includes(u)) {
    nodes.push(u);
  }
  if (!nodes.includes(v)) {
    nodes.push(v);
  }
});

const swap = (arr, i, j) => {
  let t = arr[i];
  arr[i] = arr[j];
  arr[j] = t;
};

const permute = (arr, i, result) => {
  if (i == arr.length) {
    result.push(arr.slice());
    return;
  }

  for (let j = i; j < arr.length; j++) {
    // swap
    swap(arr, i, j);
    // permute the subtree
    permute(arr, i + 1, result);
    // backtrack
    swap(arr, i, j);
  }
};

const get_route_dist = (route) => {
  let route_dist = 0;
  for (let i = 0; i < route.length - 1; i++) {
    const u = route[i];
    const v = route[i + 1];
    route_dist += graph[u][v];
  }

  return route_dist;
};

let routes = [];

permute(nodes, 0, routes);

let distances = routes.map(get_route_dist);

let part1 = Math.min(...distances);
console.log("Part 1: ", part1);
let part2 = Math.max(...distances);
console.log("Part 2: ", part2);
