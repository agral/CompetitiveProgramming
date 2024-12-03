#include <cassert>
#include <string>
#include <sstream>
#include <vector>

#include <iostream>

class Solution {
public:
    std::string addSpaces(std::string s, std::vector<int>& spaces) {
        std::stringstream ans;
        std::vector<int>::iterator it = spaces.begin();
        const int SZ = s.length();
        for (int k = 0; k < SZ; ++k) {
            if ((it != spaces.end()) && (k == *it)) {
                ans << " ";
                it++;
            }
            ans << s[k];
        }

        return ans.str();
    }
};

int main() {
    Solution s;
    {
        std::string st = {"LeetcodeHelpsMeLearn"};
        std::vector<int> spaces = {8, 13, 15};
        assert(s.addSpaces(st, spaces) == "Leetcode Helps Me Learn");
    }
    {
        std::string st = {"icodeinpython"};
        std::vector<int> spaces = {1, 5, 7, 9};
        assert(s.addSpaces(st, spaces) == "i code in py thon");
    }
    {
        std::string st = {"spacing"};
        std::vector<int> spaces = {0, 1, 2, 3, 4, 5, 6};
        assert(s.addSpaces(st, spaces) == " s p a c i n g");
    }
}
