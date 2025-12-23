#include <vector>
#include <cassert>
#include <cmath>

class Solution {
public: 
  int minSpeedOnTime(std::vector<int>& dist, double hour) {

    if (hour <= static_cast<double>(dist.size() - 1)) {
      // as each train takes at least 1 full hour to complete + embark on next one,
      // it is impossible to arrive on time if hour <= #trains - 1.
      return -1;
    }
    int low = 1; // lowest bound for optimal train speed.
    int high = 1e7; // highest bound
    while (low < high) {
      int mid = low + (high - low) / 2;
      if (this->canMakeOnTime(dist, hour, mid)) {
        // Midpoint can make it on time, making it an upper bound for true optimal speed.
        // (trains of midpoint speed reach target in desired time)
        high = mid;
      }
      else {
        // Midpoint can not make it on time, making it a lower bound for true optimal speed.
        // (trains of speed up to midpoint won't reach target in time)
        low = mid + 1;
      }
    }
    return low;
  }

private:
  bool canMakeOnTime(std::vector<int>& dist, double hour, double train_speed) {
    double travel_time = 0.0;
    for (auto it = dist.begin(); it != dist.end() - 1; it++) {
      travel_time += ceil(*it / train_speed);
    }
    travel_time += dist.back() / train_speed;
    bool ans = travel_time <= hour;
    return ans;
  }
};

int main() {
  Solution s{};
  std::vector<int> test1 = {1, 3, 2};
  assert(s.minSpeedOnTime(test1, 6.0) == 1);
  assert(s.minSpeedOnTime(test1, 2.7) == 3);
  assert(s.minSpeedOnTime(test1, 1.9) == -1);

  std::vector<int> input4 = {1, 1, 100'000};
  assert(s.minSpeedOnTime(input4, 2.01) == 100 * 100'000);

  std::vector<int> input5 = {1, 1};
  assert(s.minSpeedOnTime(input5, 1.0) == -1);

  std::vector<int> tc30 = {5, 3, 4, 6, 2, 2, 7};
  assert(s.minSpeedOnTime(tc30, 10.92) == 4);
} 
