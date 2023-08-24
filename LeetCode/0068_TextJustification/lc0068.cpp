#include <cassert>
#include <iostream>
#include <vector>
#include <sstream>
#include <string>

class Solution {
public:
  std::vector<std::string> fullJustify(std::vector<std::string>& words, int maxWidth) {
    std::vector<std::string> ans;

    std::vector<std::string> line_words;
    std::size_t current_word_idx = 0;
    int line_words_length = 0; // holds the total length of all words in current line, without inter-word spaces.
    while (current_word_idx < words.size()) {
      auto& current_word = words[current_word_idx];
      if (current_word.size() + line_words_length + line_words.size() <= maxWidth) {
        // If current word with inter-word space fits, just accumulate:
        line_words.push_back(current_word);
        line_words_length += current_word.size();
      }
      else {
        // Otherwise justify the existing line contents, 
        std::stringstream line_ss;
        int num_spaces_left = maxWidth - line_words_length;

        // Distribute spaces. Corner case: when there's just one word, apply spaces right after word.
        if (line_words.size() == 1) {
          line_ss << line_words[0];
          for (int s = line_words[0].size(); s < maxWidth; s++) {
            line_ss << ' ';
          }
        }
        else {
          // There's more than one word. Distribute spaces evenly across all n-1 inter-word slots:
          int common_spaces = num_spaces_left / (line_words.size() - 1);
          int num_extra_spaces = num_spaces_left % (line_words.size() - 1);
          for (std::size_t i = 0; i < line_words.size(); i++) {
            line_ss << line_words[i];
            //Distribute spaces after every word except for the last one.
            if (i < line_words.size() - 1) {
              for (int s = 0; s < common_spaces; s++) {
                line_ss << ' ';
              }
              if (num_extra_spaces > 0) {
                line_ss << ' ';
                num_extra_spaces -= 1;
              }
            }
          }
        }

        // Store the justified line in answer's array and start a new line with this new word.
        // It is guaranteed in the problem statement that each word's length will not exceed maxWidth,
        // so it can be safely stored in a fresh line.
        ans.push_back(line_ss.str());
        line_words.clear();
        line_words.push_back(current_word);
        line_words_length = current_word.size();
      }

      // After processing this word advance to the next word.
      current_word_idx += 1;
    }
    // Finally, output the last line. Last line consists of words separated by single spaces,
    // then empty spaces until maxWidth is reached.
    int length_counter = 0;
    std::stringstream line_ss;
    for (std::size_t i = 0; i < line_words.size(); i++) {
      if (i > 0) {
        line_ss << ' ';
        length_counter += 1;
      }
      line_ss << line_words[i];
      length_counter += line_words[i].size();
    }
    for (;length_counter < maxWidth; length_counter++) {
      line_ss << ' ';
    }
    //line_ss << "|";
    ans.push_back(line_ss.str());
    return ans;
  }
};

int main() {
  Solution s{};

  std::vector<std::string> words1 = {"This", "is", "an", "example", "of", "text", "justification."};
  std::vector<std::string> expected1 = {
      "This    is    an",
      "example  of text",
      "justification.  "
  };
  auto ans1 = s.fullJustify(words1, 16);
  for (const auto& line: ans1) {
    std::cout << line << "|\n";
  }
  assert(ans1 == expected1);

  std::vector<std::string> words2 = {"What", "must", "be", "acknowledgment", "shall", "be"};
  std::vector<std::string> expected2 = {
      "What   must   be",
      "acknowledgment  ",
      "shall be        "
  };
  auto ans2 = s.fullJustify(words2, 16);
  for (const auto& line: ans2) {
    std::cout << line << "|\n";
  }
  assert(ans2 == expected2);

  std::vector<std::string> words3 = {"Science", "is", "what", "we", "understand", "well", "enough",
      "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"};
  std::vector<std::string> expected3 = {
      "Science  is  what we",
      "understand      well",
      "enough to explain to",
      "a  computer.  Art is",
      "everything  else  we",
      "do                  "
  };
  auto ans3 = s.fullJustify(words3, 20);
  for (const auto& line: ans3) {
    std::cout << line << "|\n";
  }
  assert(ans3 == expected3);
}
