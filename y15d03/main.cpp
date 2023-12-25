#include "helper.h"
#include <iostream>
#include <map>
#include <ostream>
#include <set>
#include <vector>
using namespace std;

int main(int argc, char *argv[]) {
  if (argc < 2) {
    cout << "Provide the input file name" << endl;
    return 1;
  }

  const string moves = read_lines(argv[1])[0];
  Position santa_pos;
  const map<char, Position> move_dir = {{'^', Position(-1, 0)},
                                        {'v', Position(1, 0)},
                                        {'<', Position(0, -1)},
                                        {'>', Position(0, 1)}};

  set<Position> houses1 = {santa_pos};

  Position both_pos[] = {Position(), Position()};
  set<Position> houses2 = {both_pos[0]};

  for (size_t i = 0; i < moves.size(); i++) {
    // find direction
    const auto &dir = moves[i];
    auto it = move_dir.find(dir);
    Position other = it->second;

    // Part 1
    santa_pos = santa_pos + other;
    houses1.insert(santa_pos);

    // Part 2
    both_pos[i % 2] = both_pos[i % 2] + other;
    houses2.insert(both_pos[i % 2]);
  }

  cout << "Part 1: " << houses1.size() << endl;
  cout << "Part 2: " << houses2.size() << endl;

  return 0;
}
