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

template <typename T> vector<T> reverse_vector(const vector<T> &orig) {
  return vector<T>(orig.rbegin(), orig.rend());
}

#endif // !HELPER_H
