#include <iostream>
#include <algorithm>
#include <vector>
#include <string>
#include <sstream> // for istringstream, stringstream

using std::vector;
using std::string;
using std::istringstream;
using std::stringstream;

class Solution {
public:
  // Return a string of words in reverse order concatenated by a single space.
  // e.g. "the     sky      is blue   " --> "blue is the sky"
  string reverseWords(string s) {
    // declare necessary variables:
    vector<string> words; // a vector of strings for holding separate words
    string word; // a temporary variable for storing next space-separated word retrieved from a stream
    istringstream word_stream(s); // a stream for performing tokenization of input string on whitespace.

    // Split up words from input stream to the vector, removing repeated whitespace:
    while (word_stream >> word) {
      words.push_back(word);
    }

    // Corner case: there might be an input string consisting of only whitespace.
    if (words.empty()) {
      return "";
    }

    // Create a single string consisting of vector elements in reversed order,
    // separated by single space:
    stringstream ss;
    ss << words.back();

    for (auto it = words.crbegin() + 1; it != words.crend(); it++) {
      ss << " " << (*it);
    }

    return ss.str();
  }

  // A test harness with pretty printing:
  void test(string s) {
    std::cout << "[" << reverseWords(s) << "]\n";
  }
};

int main() {
  Solution s;
  string input1{"the sky is blue"};
  s.test(input1);

  string input2{"      hello world   "};
  s.test(input2);

  string input3{"a good     example"};
  s.test(input3);
}
