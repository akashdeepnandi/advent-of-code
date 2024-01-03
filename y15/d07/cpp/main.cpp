#include "helper.h"
#include <iostream>
#include <map>
#include <ostream>
#include <string>
#include <vector>
using namespace std;

int find_val(map<string, vector<string>> &instructions, const string &target) {
  map<string, int> values;
  regex pattern("\\d+");

  auto get_wire_val = [&pattern, &values](string wire) {
    auto m = re_match(wire, pattern);
    if (m.size() > 0) {
      return stoi(m[0]);
    }

    map<string, int>::iterator it = values.find(wire);
    if (it != values.end()) {
      return it->second;
    }

    return -1;
  };

  while (true) {
    bool resolved = true;
    for (map<string, vector<string>>::iterator it = instructions.begin();
         it != instructions.end(); it++) {
      string tid = it->first;
      vector<string> opr = it->second;
      string gate = opr[0], op1 = opr[1], op2 = opr[2];
      int v1 = get_wire_val(op1);
      if (v1 == -1) {
        resolved = false;
        continue;
      }

      if (gate == "") {
        values[tid] = v1;
      } else if (gate == "NOT") {
        values[tid] = 65535 - v1;
      } else {
        int v2 = get_wire_val(op2);
        if (v2 == -1) {
          resolved = false;
          continue;
        }

        if (gate == "AND") {
          values[tid] = v1 & v2;
        } else if (gate == "OR") {
          values[tid] = v1 | v2;
        } else if (gate == "LSHIFT") {
          values[tid] = v1 << v2 & 0xffff;
        } else if (gate == "RSHIFT") {
          values[tid] = v1 >> v2;
        }
      }
    }
    if (resolved) {
      break;
    }
  }

  return values[target];
}

int main(int argc, char *argv[]) {
  if (argc < 2) {
    cout << "Provide the input file name" << endl;
    return 1;
  }

  vector<string> lines = read_lines(argv[1]);
  map<string, vector<string>> instructions;
  for (const auto &line : lines) {
    auto splits = split_string(line, " -> ");
    string opr = splits[0], tid = splits[1];
    vector<string> temp = split_string(opr, " ");
    vector<string> ins;
    ins.push_back(temp.size() == 1 ? "" : temp.size() == 2 ? "NOT" : temp[1]);
    ins.push_back(temp.size() == 2 ? temp[1] : temp[0]);
    ins.push_back(temp.size() == 3 ? temp[2] : "");
    instructions[tid] = ins;
  }

  // part 1
  int part1 = find_val(instructions, "a");
  cout << "Part 1: " << part1 << endl;

  // part 2
  instructions["b"][1] = to_string(part1);
  int part2 = find_val(instructions, "a");
  cout << "Part 2: " << part2 << endl;

  return 0;
}
