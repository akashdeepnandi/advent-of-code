#include "helper.h"
#include <algorithm>
#include <iostream>
#include <ostream>
#include <vector>
using namespace std;

vector<int> parse_instruction(const vector<string> &instruction) {
  int mode = -1, si = 2, ei = 4;
  if (instruction.size() == 4) {
    si = 1, ei = 3;
    mode = 0;
  } else if (instruction[1] == "on") {
    mode = 1;
  } else {
    mode = 2;
  }

  vector<int> r1 = s_vec_to_int(split_string(instruction[si], ","));
  vector<int> r2 = s_vec_to_int(split_string(instruction[ei], ","));

  // returns mode, start_r, start_c, end_r, end_c
  return {mode, r1[0], r1[1], r2[0], r2[1]};
}

int operate_decoration(const vector<string> &lines,
                       function<int(int, int)> operate) {
  int grid[1000][1000] = {0};

  for (const auto &line : lines) {
    const vector<int> instruction = parse_instruction(split_string(line, " "));
    for (int r = instruction[1]; r <= instruction[3]; r++) {
      for (int c = instruction[2]; c <= instruction[4]; c++) {
        grid[r][c] = operate(grid[r][c], instruction[0]);
      }
    }
  }
  int count = 0;
  for (int r = 0; r < 1000; r++) {
    for (int c = 0; c < 1000; c++) {
      count += grid[r][c];
    }
  }

  return count;
}

int main(int argc, char *argv[]) {
  if (argc < 2) {
    cout << "Provide the input file name" << endl;
    return 1;
  }

  vector<string> lines = read_lines(argv[1]);

  auto operate_light = [](int v, int mode) {
    switch (mode) {
    case 0:
      return v == 1 ? 0 : 1;
    case 1:
      return 1;
    case 2:
      return 0;
    default:
      return 0;
    }
  };

  auto operate_brightness = [](int v, int mode) {
    switch (mode) {
    case 0:
      return v + 2;
    case 1:
      return v + 1;
    case 2:
      return max(0, v - 1);
    default:
      return 0;
    }
  };

  cout << "Part 1: " << operate_decoration(lines, operate_light) << endl;
  cout << "Part 2: " << operate_decoration(lines, operate_brightness) << endl;

  return 0;
}
