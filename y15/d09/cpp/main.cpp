#include "helper.h"
#include <algorithm>
#include <iostream>
#include <map>
#include <ostream>
#include <string>
#include <utility>
#include <vector>
using namespace std;

void permute(vector<string>& arr, int i, vector<vector<string>>& result) {
  if (i == arr.size()) {
    result.push_back(arr);
    return;
  }

  for (size_t j = i; j < arr.size() ; j++) {
    // swap
    swap(arr[i], arr[j]);
    // permute subtree 
    permute(arr, i+1, result);
    // backtrack
    swap(arr[i], arr[j]);
  }
}

int main(int argc, char *argv[]) {
  if (argc < 2) {
    cout << "Provide the input file name" << endl;
    return 1;
  }

  vector<string> lines = read_lines(argv[1]);
  map<string, map<string, int>> graph;
  
  for(const auto& line: lines) {
    vector<string> parts = split_string(line, " ");
    string u = parts[0];
    string v = parts[2];
    int d = stoi(parts[parts.size() - 1]);

    graph[u][v] = d;
    graph[v][u] = d;
  }

  vector<string> nodes;

  for (const auto& [k,v]: graph) {
    nodes.push_back(k);
  }

  vector<vector<string>> routes;
  permute(nodes, 0, routes);


  vector<int> distances;
  for(const auto& route: routes) {
    int route_dist = 0;
    for (size_t i = 0; i < route.size() - 1; i++) {
      string u = route[i];
      string v = route[i + 1];
      route_dist += graph[u][v];
    }

    distances.push_back(route_dist);
  }

  int part1 = *min_element(distances.begin(), distances.end());
  cout << "Part 1: " << part1 << endl;

  int part2 = *max_element(distances.begin(), distances.end());
  cout << "Part 2: " << part2 << endl;
  return 0;
}
