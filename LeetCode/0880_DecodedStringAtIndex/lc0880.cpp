#include <cassert>
#include <string>

#include <iostream>

class Solution {
public:
  std::string decodeAtIndex(std::string s, int k) {
    long long length = 0;
    long long cursor = 0; // follows the encoded string
    while (length < k) {
      if (s[cursor] >= '2' && s[cursor] <= '9') { // if current char is a digit in {2, 9} range,
        length *= (s[cursor] - '0');
      }
      else {
        length++;
      }
      cursor++;
    }

    // Now traverse the string in reverse to find out the exact char to return:
    for (int c = cursor - 1; c >= 0; c--) {
      if (s[c] >= '2' && s[c] <= '9') {
        length /= (s[c] - '0');
        k = k % length;
      }
      else {
        if (k == length || k == 0) {
          return {s[c]};
        }
        length--;
      }
    }
    throw;
  }
};

int main() {
  Solution s;
  assert(s.decodeAtIndex("leet2code3", 10) == "o");
  assert(s.decodeAtIndex("ha22", 5) == "h");
  assert(s.decodeAtIndex("a2345678999999999999999", 1) == "a");
}
