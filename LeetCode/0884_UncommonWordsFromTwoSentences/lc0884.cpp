#include <cassert>
#include <sstream>
#include <string>
#include <unordered_map>
#include <vector>

class Solution {
public:
    std::vector<std::string> uncommonFromSentences(std::string s1, std::string s2) {
        // split words from both strings; count them in an unordered map:
        std::stringstream ss;
        ss << s1 << " " << s2;
        std::string word;
        std::unordered_map<std::string, int> occurrences;
        while (ss >> word) {
            occurrences[word]++;
        }
        std::vector<std::string> ans;
        for (const auto& [word, count]: occurrences) {
            if (count == 1) {
                ans.push_back(word);
            }
        }
        return ans;
    }
};

int main() {
    Solution s;
    {
        std::vector<std::string> expected = {"sour", "sweet"};
        assert(s.uncommonFromSentences("this apple is sweet", "this apple is sour") == expected);
    }
    {
        std::vector<std::string> expected = {"banana"};
        assert(s.uncommonFromSentences("apple apple", "banana") == expected);
    }
}
