#include <iostream>
#include <vector>

// Runtime complexity:
// Auxiliary space complexity:
class Solution {
public:
    int maxTaskAssign(std::vector<int>& tasks, std::vector<int>& workers,
                      int pills, int strength) {
        int ans = 0;
        // TODO 
        return ans;

    }
};

int main() {
    struct Testcase {
        std::vector<int> tasks;
        std::vector<int> workers;
        int pills;
        int strength;
        int expected;
    };
    Testcase testcases[] = {
        {/*tasks=*/{3, 2, 1},
         /*workers=*/{0, 3, 3},
         /*pills=*/1,
         /*strength=*/1,
         /*expected=*/3
        },
        {/*tasks=*/{5, 4},
         /*workers=*/{0, 0, 0},
         /*pills=*/1,
         /*strength=*/5,
         /*expected=*/1
        },
        {/*tasks=*/{10, 15, 30},
         /*workers=*/{0, 10, 10, 10, 10},
         /*pills=*/3,
         /*strength=*/10,
         /*expected=*/2
        },
    };
    Solution s{};
    int numGood = 0, numBad = 0;
    for (Testcase& tc: testcases) {
        auto actual = s.maxTaskAssign(tc.tasks, tc.workers, tc.pills, tc.strength);
        if (actual != tc.expected) {
            std::cout << "Testcase failed. Got: " << actual
                << ", want: " << tc.expected << "\n";
            ++numBad;
        } else {
            ++numGood;
        }
    }
    std::cout << (numBad == 0 ? "[OK]" : "[FAIL]") << " "
        << numGood << "/" << (numBad + numGood) << " testcases passed successfully.\n";
}
