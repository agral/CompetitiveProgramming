#include <algorithm>
#include <cassert>
#include <string>
#include <sstream>
#include <vector>

class Solution {
public:
    int findMinDifference(std::vector<std::string>& timePoints) {
        // Each timePoint is a string in "HH:MM" format denoting a time point.
        int hour, minute;
        char colon;
        std::vector<int> minutes;
        for (const auto& tp: timePoints) {
            std::stringstream time_point_stream{tp};
            time_point_stream >> hour >> colon >> minute;
            minutes.push_back(60 * hour + minute);
        }

        std::sort(minutes.begin(), minutes.end());

        // It is guaranteed that there are at least two time points in timePoints vector.
        // Take into account that minimum time difference could also be between first and last time point (as the clock loops around):
        int ans = 60 * 24 + minutes[0] - minutes[minutes.size() - 1];

        for (int i = 0; i < minutes.size() - 1; i++) {
            ans = std::min(ans, minutes[i+1] - minutes[i]);
        }

        return ans;
    }
};

int main() {
    Solution s;
    {
        std::vector<std::string> tps = {"23:59", "00:00"};
        assert(s.findMinDifference(tps) == 1);
    }
    {
        std::vector<std::string> tps = {"00:00", "23:59", "00:00"};
        assert(s.findMinDifference(tps) == 0);
    }
    {
        std::vector<std::string> tps = {"01:01", "13:01"};
        assert(s.findMinDifference(tps) == 12 * 60);
    }
    {
        std::vector<std::string> tps = {"00:00", "12:00", "23:00"};
        assert(s.findMinDifference(tps) == 60);
    }
}
