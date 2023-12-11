#include <cassert>
#include <vector>

class Solution {
public:
  int findSpecialInteger(std::vector<int>& arr) {
    const std::size_t sz = arr.size();
    const std::size_t q = sz / 4;
    for (std::size_t i = 0; i < sz - q; i++) {
      if (arr[i] == arr[i+q]) {
        return arr[i];
      }
    }
    return -1;
  }
};

int main() {
  Solution s;
  {
    std::vector<int> arr = {1, 2, 2, 6, 6, 6, 6, 7, 10};
    assert(s.findSpecialInteger(arr) == 6);
  }
  {
    std::vector<int> arr = {1, 1};
    assert(s.findSpecialInteger(arr) == 1);
  }
}
