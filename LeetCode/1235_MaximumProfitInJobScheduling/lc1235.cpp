#include <algorithm>
#include <cassert>
#include <vector>

class Job {
public:
    Job(int startTime, int endTime, int profit)
    : m_startTime(startTime), m_endTime(endTime), m_profit(profit) {}
public:
    int m_startTime;
    int m_endTime;
    int m_profit;
};

class Solution {
public:
    int jobScheduling(std::vector<int>& startTime, std::vector<int>& endTime, std::vector<int>& profit) {
        const int SIZE = startTime.size();
        std::vector<Job> jobs;

        for (std::size_t i = 0; i < SIZE; ++i) {
            jobs.push_back({startTime[i], endTime[i], profit[i]});
        }

        // sort the Jobs in ascending order w.r.t. start time:
        std::ranges::sort(jobs, [](const Job& lhs, const Job& rhs){
            return lhs.m_startTime < rhs.m_startTime;
        });

        // also need to update startTime array - this will be used
        // in pickOptimalSchedule, and a contiguous vector is required
        // to use std::lower_bound().
        for (std::size_t i = 0; i < SIZE; ++i) {
            startTime[i] = jobs[i].m_startTime;
        }

        dp.clear();
        dp.resize(SIZE + 1, 0);
        return pickOptimalSchedule(jobs, startTime, 0);
    }
private:
    int pickOptimalSchedule(const std::vector<Job>& jobs,
                            std::vector<int>& startTime, int i) {
        if (i >= jobs.size()) {
            return 0;
        }
        if (dp[i] > 0) {
            return dp[i];
        }

        const int threshold = std::lower_bound(
            startTime.begin() + i + 1, startTime.end(), jobs[i].m_endTime
        ) - startTime.begin();
        const int when_picked = jobs[i].m_profit + 
                pickOptimalSchedule(jobs, startTime, threshold);
        const int when_skipped = pickOptimalSchedule(jobs, startTime, i + 1);
        dp[i] = std::max(when_picked, when_skipped);
        return dp[i];
    }

    std::vector<int> dp;
};

int main() {
  Solution s;
  {
    std::vector<int> startTime = {1, 1, 3, 3};
    std::vector<int> endTime = {3, 4, 5, 6};
    std::vector<int> profit = {20, 20, 100, 70, 60};
    assert(s.jobScheduling(startTime, endTime, profit) == 120);
  }
  {
    std::vector<int> startTime = {1, 2, 3, 4, 6};
    std::vector<int> endTime = {3, 5, 10, 6, 9};
    std::vector<int> profit = {20, 20, 100, 70, 60};
    assert(s.jobScheduling(startTime, endTime, profit) == 150);
  }
  {
    std::vector<int> startTime = {1, 1, 1};
    std::vector<int> endTime = {2, 3, 4};
    std::vector<int> profit = {5, 6, 4};
    assert(s.jobScheduling(startTime, endTime, profit) == 6);
  }
}
