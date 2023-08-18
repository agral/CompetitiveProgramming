#include <cassert>
#include <vector>
#include <iostream>
#include <string>

class Solution {
public:
  int compress(std::vector<char>& chars) {
    char accumulator_char = chars[0];
    int accumulator_count = 0;
    int write_idx = 0;
    for (int read_idx = 0; read_idx < chars.size(); read_idx++) {
      if (chars[read_idx] == accumulator_char) {
        accumulator_count += 1;
      } else {
        // write out compressed stream:
        chars[write_idx] = accumulator_char;
        write_idx += 1;
        if (accumulator_count > 1) {
          for (const auto& ch : std::to_string(accumulator_count)) {
            chars[write_idx] = ch;
            write_idx += 1;
          }
        }

        // remember new char in acc, set streak to 1:
        accumulator_char = chars[read_idx];
        accumulator_count = 1;
      }
    }
    // Write the last group too:
    chars[write_idx] = accumulator_char;
    write_idx += 1;
    if (accumulator_count > 1) {
      for (const auto& ch : std::to_string(accumulator_count)) {
        chars[write_idx] = ch;
        write_idx += 1;
      }
    }
    return write_idx;
  }
};

int main() {
  Solution s{};

  std::vector<char> testcase1 = {'a', 'a', 'b', 'b', 'c', 'c', 'c'}; //compressed: a2b2c3
  auto len1 = s.compress(testcase1);
  assert(len1 == 6);
  assert(testcase1[0] == 'a');
  assert(testcase1[1] == '2');
  assert(testcase1[2] == 'b');
  assert(testcase1[3] == '2');
  assert(testcase1[4] == 'c');
  assert(testcase1[5] == '3');

  std::vector<char> testcase2 = {'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'};
  auto len2 = s.compress(testcase2);
  assert(len2 == 4);
  assert(testcase2[0] == 'a');
  assert(testcase2[1] == 'b');
  assert(testcase2[2] == '1');
  assert(testcase2[3] == '2');

  std::vector<char> testcase3 = {'a'};
  auto len3 = s.compress(testcase3);
  assert(len3 == 1);
  assert(testcase3[0] == 'a');

  std::vector<char> testcase4 = {'a', 'b', 'c', 'c', 'c', 'c'};
  auto len4 = s.compress(testcase4);
  assert(len4 == 4);
  assert(testcase4[0] == 'a');
  assert(testcase4[1] == 'b');
  assert(testcase4[2] == 'c');
  assert(testcase4[3] == '4');
}
