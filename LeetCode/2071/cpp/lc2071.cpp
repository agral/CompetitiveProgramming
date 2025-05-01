#include <algorithm>
#include <iostream>
#include <map>
#include <vector>

class Solution {
public:
    int maxTaskAssign(std::vector<int>& tasks, std::vector<int>& workers,
                      int pills, int strength) {
        int ans = 0;
        int l = 0;
        int r = std::min(tasks.size(), workers.size());
        std::ranges::sort(tasks);
        std::ranges::sort(workers);

        // Returns true if numTasks can be completed using at most numPills
        auto canComplete = [&](int numTasks, int numPills) {
            std::map<int, int> sorted_workers;
            // take `numTasks` strongest workers:
            for (int i = workers.size() - numTasks; i < workers.size(); i++) {
                ++sorted_workers[workers[i]];
            }

            // process `numTasks` easiest tasks, starting from the hardest one:
            for (int i = numTasks-1; i >= 0; i--) {
                // Find the least strong worker with own strength >= tasks[i].
                // If such a worker is found, then no magic pill needs to be used,
                // which is the preferred option.
                auto it = sorted_workers.lower_bound(tasks[i]);
                if (it != sorted_workers.end()) {
                    // Use this worker.
                    it->second -= 1;
                    if (it->second == 0) {
                        sorted_workers.erase(it);
                    }
                } else if (numPills > 0) {
                    // Find the least strong worker with own strength >= tasks[i] - `strength`.
                    // This is so that this worker can still complete the task,
                    // when helped with a magic pill:
                    it = sorted_workers.lower_bound(tasks[i] - strength);
                    if (it != sorted_workers.end()) {
                        // Use this worker and one magic pill.
                        it->second -= 1;
                        --numPills;
                        if (it->second == 0) {
                            sorted_workers.erase(it);
                        }
                    } else {
                        // The current task can't be done as no worker is strong enough,
                        // even when helped by the magic pill.
                        return false;
                    }
                } else {
                    // There are no magic pills left, and the current task can't be done
                    // as no worker is strong enough on his own.
                    return false;
                }
            }
            // All the `numTasks` tasks have been completed successfully.
            return true;
        };

        while (l <= r) {
            int mid = (l + r) / 2;
            if (canComplete(mid, pills)) {
                ans = mid;
                l = mid + 1;
            } else {
                r = mid - 1;
            }
        }

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
