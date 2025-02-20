#include <iostream>
#include <sstream>
#include <vector>

class Solution {
public:
    std::string findDifferentBinaryString(std::vector<std::string>& nums) {
        // Solution: construct the answer by flipping ith bit of ith input string.
        // This is guaranteed to generate a number of length N that is not yet in nums.
        std::stringstream ans;
        for (int idx = 0; idx < nums.size(); ++idx) {
            ans << (nums[idx][idx] == '0' ? '1' : '0');
        }
        return ans.str();
    }
};

int main() {
    struct Testcase {
        std::vector<std::string> nums;
        bool isValidAnswer(std::string candidate) {
            if (candidate.length() != nums[0].length()) {
                return false;
            }
            for (const std::string& num: nums) {
                if (num == candidate) {
                    return false;
                }
            }
            return true;
        }
    };
    Testcase testcases[] = {
        std::vector<std::string>{"01", "10"},
        std::vector<std::string>{"00", "01"},
        std::vector<std::string>{"111", "011", "001"},
    };
    Solution s{};
    int numGood = 0, numBad = 0, counter = 0;
    for (Testcase& tc: testcases) {
        counter++;
        auto actual = s.findDifferentBinaryString(tc.nums);
        if (!tc.isValidAnswer(actual)) {
            std::cout << "Testcase #" << counter << " failed. Got: " << actual << "\n";
            ++numBad;
        } else {
            ++numGood;
        }
    }
    std::cout << (numBad == 0 ? "[OK]" : "[FAIL]") << " "
        << numGood << "/" << (numBad + numGood) << " testcases passed successfully.\n";
}
