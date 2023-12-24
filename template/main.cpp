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
  for (const auto line : lines) {
    cout << line << endl;
  }

  return 0;
}
