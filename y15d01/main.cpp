#include "helper.h"
#include <iostream>
#include <vector>
using namespace std;

int main(int argc, char *argv[]) {
  if (argc < 2) {
    cout << "Provide the input file name" << endl;
    return 1;
  }

  vector<string> lines = read_lines(argv[1]);
  int floor = 0;
  int basement_index = -1;
  for (size_t i = 0; i < lines[0].size(); i++) {
    const auto ch = lines[0][i];
    // part 1
    if (ch == '(') {
      floor++;
    } else {
      floor--;
    }

    // part 2
    if (floor == -1 && basement_index == -1) {
      basement_index = i + 1;
    }
  }

  cout << "Floor is: " << floor << endl;
  cout << "Basement Reached at: " << basement_index << endl;
  return 0;
}
