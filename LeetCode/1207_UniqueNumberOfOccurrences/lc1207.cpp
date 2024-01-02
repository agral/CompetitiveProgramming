#include <cassert>
#include <unordered_map>
#include <unordered_set>
#include <vector>

class Solution {
public:
  bool uniqueOccurrences(std::vector<int>& arr) {
    // Make a frequency map: number -> count-of-number-in-arr
    std::unordered_map<int, int> freq;
    for (const int num: arr) {
      ++freq[num];
    }

    // Push each entry from frequency map into uniques set.
    // If the entry can not be pushed, it is a duplicate.
    std::unordered_set<int> uniques;
    for (const auto& [_, frequency]: freq) {
      if (!uniques.insert(frequency).second) {
        return false;
      }
    }
    return true;
  }
};

int main() {
  Solution s;
  {
    std::vector<int> arr = {1, 2, 2, 1, 1, 3};
    assert(s.uniqueOccurrences(arr) == true);
  }
  {
    std::vector<int> arr = {1, 2};
    assert(s.uniqueOccurrences(arr) == false);
  }
  {
    std::vector<int> arr = {-3, 0, 1, -3, 1, 1, 1, -3, 10, 0};
    assert(s.uniqueOccurrences(arr) == true);
  }
}
