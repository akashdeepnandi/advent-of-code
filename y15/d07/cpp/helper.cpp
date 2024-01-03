#include <fstream>
#include <regex>
#include <string>
#include <vector>
using namespace std;

vector<string> split_string(const string &str, const string &delimiter) {
  vector<string> tokens;
  size_t start = 0, end = 0;
  if (delimiter.empty()) {
    for (char ch : str) {
      tokens.push_back(string(1, ch));
    }

    return tokens;
  }

  while ((end = str.find(delimiter, start)) != string::npos) {
    tokens.push_back(str.substr(start, end - start));
    start = end + delimiter.length();
  }

  tokens.push_back(str.substr(start)); // Last token (after the last delimiter)

  return tokens;
}

// Usage:
// string inputString = "123 4234 543534 543534";
// regex pattern("\\d+");
// vector<string> matches = re_match(inputString, pattern);

vector<string> re_match(const string input, regex pattern) {
  vector<string> matches;
  sregex_iterator iter(input.begin(), input.end(), pattern);
  sregex_iterator end; // End iterator

  // Iterate through matches and print each match
  while (iter != end) {
    matches.push_back(iter->str());
    ++iter; // Move to the next match
  }

  return matches;
}

vector<string> read_lines(const string filename) {
  vector<string> lines;

  ifstream input_file(filename);
  if (input_file.is_open()) {
    string line;

    while (getline(input_file, line)) {
      lines.push_back(line);
    }
  }

  return lines;
}

vector<int> s_vec_to_int(const vector<string> &s_vec) {
  vector<int> r_vec;
  for (const auto &elem : s_vec) {
    r_vec.push_back(stoi(elem));
  }
  return r_vec;
}
