#include "helper.h"
#include <iostream>
#include <ostream>
#include <string>
#include <vector>
using namespace std;

string look_and_say(const string &sequence) {
  char current_digit = sequence[0];
  int count = 0;
  string result = "";

  for(const auto& d: sequence) {
    if (current_digit == d) {
      count++;
    } else {
      result += to_string(count);
      result.push_back(current_digit);

      current_digit = d;
      count = 1;
    }
  }
  //add the last one
  result += to_string(count);
  result.push_back(current_digit);

  return result;
}

string evolve_sequence(string sequence, int generation) {
  for (size_t i = 0; i < generation; i++) {
    sequence = look_and_say(sequence);
  }
  return sequence;
}

int main(int argc, char *argv[]) {
  if (argc < 2) {
    cout << "Provide the input file name" << endl;
    return 1;
  }

  vector<string> lines = read_lines(argv[1]);
  string sequence = lines[0];

  string part1 = evolve_sequence(sequence, 40);
  cout << "Part 1: " << part1.size() << endl;

  string part2 = evolve_sequence(sequence, 50);
  cout << "Part 2: " << part2.size() << endl;

  return 0;
}
