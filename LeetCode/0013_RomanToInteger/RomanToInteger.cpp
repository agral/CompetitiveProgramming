// Solution to LeetCode problem #0013, "Roman to Integer"
// https://leetcode.com/problems/roman-to-integer/
// by https://gralin.ski/

#include <iostream>
#include <string>

class Solution {
 public:
  /*
   * Converts a roman number in range 1-3999 to an int
   * Expects the number to contain only the characters I,V,X,L,C,D,M.
   */
  int romanToInt(std::string s) {
    int decoded_number = 0;
    auto it_current_char = s.cbegin();
    char current_char, next_char;
    while (it_current_char != s.cend()) {
      current_char = *it_current_char;
      next_char = '-';
      if (it_current_char + 1 < s.cend()) {
        next_char = *(it_current_char + 1);
      }

      //std::cout << "Processing character: " << current_char << ", next char is " << next_char << " ...\n";
      if (current_char == 'I') {
        // I can be placed before V, the combined token IV has the value of 4:
        if (next_char == 'V') {
          ++it_current_char; // one extra character has been consumed for this token.
          decoded_number += 4;
        }
        // I can also be placed before X, the combined IX token has the value of 9:
        else if (next_char == 'X') {
          ++it_current_char; // one extra character has been consumed for this token.
          decoded_number += 9;
        }
        else {
          decoded_number += 1;
        }
      }
      else if (current_char == 'V') {
        decoded_number += 5;
      }

      else if (current_char == 'X') {
        // X can be placed before L, the combined token XL has the value of 40:
        if (next_char == 'L') {
          ++it_current_char; // one extra character has been consumed for this token.
          decoded_number += 40;
        }
        // X can also be placed before C, the combined XC token has the value of 90:
        else if (next_char == 'C') {
          ++it_current_char; // one extra character has been consumed for this token.
          decoded_number += 90;
        }
        else {
          decoded_number += 10;
        }
      }
      else if (current_char == 'L') {
        decoded_number += 50;
      }

      else if (current_char == 'C') {
        // C can be placed before D, the combined token CD has the value of 400:
        if (next_char == 'D') {
          ++it_current_char; // one extra character has been consumed for this token.
          decoded_number += 400;
        }
        // C can also be placed before M, the combined CM token has the value of 900:
        else if (next_char == 'M') {
          ++it_current_char; // one extra character has been consumed for this token.
          decoded_number += 900;
        }
        else {
          decoded_number += 100;
        }
      }
      else if (current_char == 'D') {
        decoded_number += 500;
      }
      else if (current_char == 'M') {
        decoded_number += 1000;
      }
      else {
        throw "Unexpected character";
      }

      ++it_current_char;
    }
    return decoded_number;
  }
};

int main() {
  Solution s;

  {
    std::cout << "I (1111)\n";
    if (1111 == s.romanToInt("MCXI")) {
      std::cout << "  Passed\n";
    } else {
      std::cout << "Failed\n";
    }
  }
  {
    std::cout << "MMMDLV (3555)\n";
    if (3555 == s.romanToInt("MMMDLV")) {
      std::cout << "  Passed\n";
    } else {
      std::cout << "Failed\n";
    }
  }
  {
    std::cout << "LXXX (80)\n";
    if (80 == s.romanToInt("LXXX")) {
      std::cout << "  Passed\n";
    } else {
      std::cout << "Failed\n";
    }
  }
  {
    std::cout << "MMCMLXXXIX (2989)\n";
    if (2989 == s.romanToInt("MMCMLXXXIX")) {
      std::cout << "  Passed\n";
    } else {
      std::cout << "Failed\n";
    }
  }
  {
    std::cout << "MMXXII (2022)\n";
    if (2022 == s.romanToInt("MMXXII")) {
      std::cout << "  Passed\n";
    } else {
      std::cout << "Failed\n";
    }
  }
}
