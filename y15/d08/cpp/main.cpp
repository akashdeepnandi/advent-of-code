#include "helper.h"
#include <iostream>
#include <ostream>
#include <vector>
#include <string>
using namespace std;

int main(int argc, char *argv[]) {
  if (argc < 2) {
    cout << "Provide the input file name" << endl;
    return 1;
  }

  int total = 0;
  // Part 1
  int actual_total = 0;
  // Part 2
  int extra_total = 0;
  vector<string> lines = read_lines(argv[1]);
  for(const auto& line: lines) {
    int len = line.size();
    // Part 1
    int actual = 0;
    // Part 2
    int extra = 2;

    for (size_t i = 1; i < line.size() - 1;) {
      if(line[i] == '\\') {
        if (line[i+1] == '\\' || line[i + 1] == '"') {
          actual++;
          i += 2;
        } else if (line[i+1] == 'x'){
          actual++;
          i += 4;
        } else {
          actual++;
          i ++;
        }
      } else {
          actual++;
          i ++;
      }
    }

    total += len;

    actual_total += actual;

    for (size_t i = 0; i < line.size(); ) {
      if(line[i] == '\\' || line[i] == '"') {
        extra += 2;
      } else {
        extra ++;
      }

      i++;
    }

    extra_total += extra;
  }

  int part1 = total - actual_total;
  cout << "Part 1: " <<part1 << endl;

  int part2 = extra_total - total;
  cout << "Part 2: " <<part2 << endl;

  return 0;
}
