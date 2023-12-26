#include "helper.h"
#include <iostream>
#include <openssl/md5.h>
#include <ostream>
#include <string>
#include <vector>
using namespace std;

std::string md5(const std::string &input) {
  unsigned char digest[MD5_DIGEST_LENGTH];
  MD5(reinterpret_cast<const unsigned char *>(input.c_str()), input.length(),
      digest);

  char mdString[MD5_DIGEST_LENGTH * 2 + 1];

  for (int i = 0; i < MD5_DIGEST_LENGTH; i++)
    sprintf(&mdString[i * 2], "%02x", (unsigned int)digest[i]);

  return std::string(mdString);
}

int get_iterations(string input, string prefix) {
  for (size_t i = 0;; i++) {
    string hash = md5(input + to_string(i));
    bool found = hash.find(prefix) == 0;
    if (found) {
      return i + 1;
    }
  }
  return 0;
}

int main(int argc, char *argv[]) {
  if (argc < 2) {
    cout << "Provide the input file name" << endl;
    return 1;
  }

  string input = read_lines(argv[1])[0];
  cout << "Part 1: " << get_iterations(input, "00000") << endl;
  cout << "Part 2: " << get_iterations(input, "000000") << endl;

  return 0;
}
