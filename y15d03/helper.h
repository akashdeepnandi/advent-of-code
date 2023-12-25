#include <algorithm>
#include <regex>
#include <string>
#include <vector>
using namespace std;
#ifndef HELPER_H
#define HELPER_H
vector<string> split_string(const string str, string delimiter);
vector<string> re_match(const string input, regex pattern);
vector<string> read_lines(const string filename);
vector<int> s_vec_to_int(const vector<string> &s_vec);

class Position {
private:
  int x, y;

public:
  Position(int _x = 0, int _y = 0) : x(_x), y(_y) {}

  int getX() { return x; }
  int getY() { return y; }

  Position operator+(const Position &other) {
    return Position(x + other.x, y + other.y);
  }
  Position operator-(const Position &other) {
    return Position(x - other.x, y - other.y);
  }

  bool operator<(const Position &other) const {
    if (x < other.x) {
      return true;
    } else if (x == other.x) {
      return y < other.y;
    }
    return false;
  }

  friend ostream &operator<<(ostream &os, const Position &pos) {
    os << "Position: (" << pos.x << ", " << pos.y << ")" << endl;
    return os;
  }
};

template <typename T> vector<T> reverse_vector(const vector<T> &orig) {
  return vector<T>(orig.rbegin(), orig.rend());
}

#endif // !HELPER_H
