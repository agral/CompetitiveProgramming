#include <cassert>
#include <vector>

class MountainArray {
public:
  MountainArray(std::vector<int> values): m_values(values), m_counter(0) {}

  int get(int index) {
    m_counter++;
    if (m_counter >= 100) {
      throw;
    }
    return m_values.at(index);
  }

  int length() const { return m_values.size(); }

private:
  std::vector<int> m_values;
  int m_counter;
};

class Solution {
public:
  int findInMountainArray(int target, MountainArray& mountainArr) {
    int length = mountainArr.length();
    // Note: peak is guaranteed to never be at first nor last index in the mountainArr by definition.
    // Searching between 1 and length-2 indices. When tried 0 and length-1, need to take extra care
    // with corner cases of peaks close to left or right border of the array.
    int left_index = 1;
    int right_index = length - 2;
    int peak_index = -1; // -1 meaning peak not found yet. Expect this to be any value >= 1.
    while ((peak_index == -1) && (left_index <= right_index)) {
      int mid_index = left_index + (right_index - left_index) / 2;
      int previous_val = mountainArr.get(mid_index - 1);
      int mid_val = mountainArr.get(mid_index);
      int next_val = mountainArr.get(mid_index + 1);

      if ((previous_val < mid_val) && (mid_val < next_val)) {
        // we're in the left (ascending) slope of the mountain array.
        left_index = mid_index + 1;
      }
      else if ((previous_val > mid_val) && (mid_val > next_val)) {
        // we're in the right (descending) slope of the mountain array.
        right_index = mid_index - 1;
      }
      else {
        // bingo, mid-point (the highest value of the mountain array) is found.
        peak_index = mid_index;
      }
    }

    // Search the left slope, where numbers are guaranteed to be in ascending order:
    left_index = 0;
    right_index = peak_index;
    while (left_index <= right_index) {
      int mid_index = left_index + (right_index - left_index) / 2;
      int mid_val = mountainArr.get(mid_index);
      if (mid_val < target) {
        left_index = mid_index + 1;
      }
      else if (mid_val > target) {
        right_index = mid_index - 1;
      }
      else {
        return mid_index;
      }
    }

    // Left slope did not contain target value. Search the right slope,
    // where the numbers are guaranteed to be in descending order:
    left_index = peak_index;
    right_index = length - 1;
    while (left_index <= right_index) {
      int mid_index = left_index + (right_index - left_index) / 2;
      int mid_val = mountainArr.get(mid_index);
      if (mid_val > target) {
        left_index = mid_index + 1;
      }
      else if (mid_val < target) {
        right_index = mid_index - 1;
      }
      else {
        return mid_index;
      }
    }

    // Both slopes checked, did not contain target value.
    // So target does not exist in the mountainArray at all.
    return -1;
  }
};

int main() {
  Solution s;

  MountainArray ma1{std::vector<int>({1, 2, 3, 4, 5, 3, 1})};
  assert(s.findInMountainArray(3, ma1) == 2);

  MountainArray ma2{std::vector<int>({0, 1, 2, 4, 2, 1})};
  assert(s.findInMountainArray(3, ma2) == -1);

  MountainArray ma_wa1{std::vector<int>({1, 5, 2})};
  assert(s.findInMountainArray(3, ma_wa1) == -1);
  assert(s.findInMountainArray(1, ma_wa1) == 0);
  assert(s.findInMountainArray(5, ma_wa1) == 1);
  assert(s.findInMountainArray(2, ma_wa1) == 2);

  MountainArray ma_wa2{std::vector<int>({3, 5, 3, 2, 0})};
  assert(s.findInMountainArray(3, ma_wa2) == 0);
  assert(s.findInMountainArray(5, ma_wa2) == 1);
  assert(s.findInMountainArray(2, ma_wa2) == 3);
  assert(s.findInMountainArray(0, ma_wa2) == 4);
  assert(s.findInMountainArray(1, ma_wa2) == -1);
}
