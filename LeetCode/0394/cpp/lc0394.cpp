#include <cassert>
#include <deque>
#include <string>
#include <sstream>

class Solution {
public:
    std::string decodeString(std::string s) {
        std::deque<std::pair<std::string, int>> stack;
        std::string builtString;
        int numConv = 0;

        for (const char c: s) {
            if (std::isdigit(c)) {
                numConv *= 10;
                numConv += (c - '0');
            } else {
                if (c == '[') {
                    stack.push_back(std::pair<std::string, int>(builtString, numConv));
                    numConv = 0;
                    builtString = "";
                }
                else if (c == ']') {
                    const auto [stackString, numRepeats] = stack.back();
                    stack.pop_back();
                    builtString = stackString + makeRepeatedString(builtString, numRepeats);
                }
                else {
                    builtString += c;
                }
            }
        }

        return builtString;
    }

private:
    std::string makeRepeatedString(const std::string& base, int numRepeats) {
        std::stringstream ss;
        for (int i = 0; i < numRepeats; i++) {
            ss << base;
        }
        return ss.str();
    }
    std::size_t cursor;
};

int main() {
    Solution s{};
    {
        std::string input = "3[a]2[bc]";
        assert(s.decodeString(input) == "aaabcbc");
    }
    {
        std::string input = "3[a2[c]]";
        assert(s.decodeString(input) == "accaccacc");
    }
    {
        std::string input = "2[abc]3[cd]ef";
        assert(s.decodeString(input) == "abcabccdcdcdef");
    }

}
