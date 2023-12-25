#include "helper.h"
#include <algorithm>
#include <iostream>
#include <string>
#include <vector>
using namespace std;

int main(int argc, char *argv[]) {
  if (argc < 2) {
    cout << "Provide the input file name" << endl;
    return 1;
  }

  vector<string> lines = read_lines(argv[1]);

  int part1 = 0, part2 = 0;

  for (const auto &line : lines) {
    vector<int> dims = s_vec_to_int(split_string(line, "x"));
    int l = dims[0], w = dims[1], h = dims[2];
    // Part 1
    part1 += 2 * (l * w + w * h + h * l) + min(min(l * w, w * h), h * l);
    // Part 2
    part2 += 2 * min(min(l + w, w + h), h + l) + l * w * h;
  }

  cout << "Part 1: " << part1 << endl;
  cout << "Part 2: " << part2 << endl;

  return 0;
}
