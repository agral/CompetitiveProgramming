#include <cassert>
#include <map>
#include <string>
#include <vector>

using Chars = std::vector<char>;

class Solution {
public:
    std::vector<std::string> letterCombinations(std::string digits) {
        if (digits.empty()) {
            return {};
        }
        m_digits = digits;
        m_chars = {};
        m_ans = {};
        backtrack(0);
        return m_ans;
    }

private:
    Chars m_chars;
    std::string m_digits;
    std::vector<std::string> m_ans;
    static std::map<char, std::vector<char>> digitToLetters;
    void backtrack(int index) {
        if (index == m_digits.size()) {
            m_ans.push_back({m_chars.begin(), m_chars.end()});
        }
        else {
            for (char letter: digitToLetters[m_digits[index]]) {
                m_chars.push_back(letter);
                backtrack(index + 1);
                m_chars.pop_back();
            }
        }
    }
};

std::map<char, std::vector<char>> Solution::digitToLetters = {
    {'2', {'a', 'b', 'c'}},
    {'3', {'d', 'e', 'f'}},
    {'4', {'g', 'h', 'i'}},
    {'5', {'j', 'k', 'l'}},
    {'6', {'m', 'n', 'o'}},
    {'7', {'p', 'q', 'r', 's'}},
    {'8', {'t', 'u', 'v'}},
    {'9', {'w', 'x', 'y', 'z'}},
};


int main() {
    Solution s;
    {
        std::vector<std::string> ans = s.letterCombinations("23");
        assert(ans == std::vector<std::string>({"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}));
    }
    {
        std::vector<std::string> ans = s.letterCombinations("");
        assert(ans == std::vector<std::string>{});
    }
    {
        std::vector<std::string> ans = s.letterCombinations("2");
        assert(ans == std::vector<std::string>({"a", "b", "c"}));
    }
}
