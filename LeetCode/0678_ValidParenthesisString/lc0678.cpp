#include <cassert>
#include <string>
#include <stack>

class Solution {
public:
  bool checkValidString(std::string s) {
    int minLeftParenCount = 0;
    int maxLeftParenCount = 0;
    for (const char& ch: s) {
      if (ch == '(') {
        ++minLeftParenCount;
        ++maxLeftParenCount;
      }
      else if (ch == ')') {
        // Decrement min count; but clamp to 0 if count becomes negative -
        // this corresponds to either some stars being empty strings,
        // or an invalid string; can't tell at this point yet.
        --minLeftParenCount;
        if (minLeftParenCount < 0) {
          minLeftParenCount = 0;
        }

        // Decrement max count. If it becomes negative - the entire string is invalid.
        --maxLeftParenCount;
        if (maxLeftParenCount < 0) {
          return false;
        }
      }
      else if (ch == '*') {
        // Decrement min count; but clamp to 0 if count becomes negative -
        // this too corresponds to either some stars being empty strings,
        // or an invalid string; similarly can't tell at this point yet.
        // This action specifically handles ch being treated as ')' or an empty string.
        --minLeftParenCount;
        if (minLeftParenCount < 0) {
          minLeftParenCount = 0;
        }

        // Increment max count.
        // This action handles ch being treated as '('.
        ++maxLeftParenCount;
      }
    }
    // in the end, the string is valid iff minLeftParenCount is 0.
    return minLeftParenCount == 0;
  }
};

int main() {
  Solution s;
  assert(s.checkValidString("()") == true);
  assert(s.checkValidString("(*)") == true);
  assert(s.checkValidString("(*))") == true);

  assert(s.checkValidString("*********(((((") == false);
}

