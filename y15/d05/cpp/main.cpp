#include "helper.h"
#include <iostream>
#include <ostream>
#include <regex>
#include <vector>
using namespace std;

int main(int argc, char *argv[]) {
  if (argc < 2) {
    cout << "Provide the input file name" << endl;
    return 1;
  }

  vector<string> lines = read_lines(argv[1]);
  // Part 1
  int valid_count1 = 0;
  for (const auto &line : lines) {
    // regex pattern("\\d+");
    vector<string> vowels = re_match(line, regex("(a|e|i|o|u)"));
    if (vowels.size() < 3)
      continue;
    vector<string> twice = re_match(line, regex("(\\w)\\1"));
    if (twice.size() == 0)
      continue;
    vector<string> bad_strs = re_match(line, regex("(ab|cd|pq|xy)"));
    if (bad_strs.size() > 0)
      continue;
    valid_count1++;
  }
  cout << "Part 1: " << valid_count1 << endl;

  // Part 2
  int valid_count2 = 0;
  for (const auto &line : lines) {
    // regex pattern("\\d+");
    vector<string> pairs = re_match(line, regex("(\\w{2}).*\\1"));
    if (pairs.size() == 0)
      continue;
    vector<string> onebetween = re_match(line, regex("(\\w)\\w\\1"));
    if (onebetween.size() == 0)
      continue;
    valid_count2++;
  }
  cout << "Part 2: " << valid_count2 << endl;


  return 0;
}
