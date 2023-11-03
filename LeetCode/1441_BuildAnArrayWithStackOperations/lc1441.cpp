#include <cassert>
#include <string>
#include <vector>

class Solution {
public:
  const std::string PUSH {"Push"};
  const std::string POP {"Pop"};
  std::vector<std::string> buildArray(std::vector<int>& target, int n) {
    std::vector<std::string> ans;
    int si = 1; // si: stream integer. Holds the next integer that will be "returned" from the stream.
    for (int ti = 0; ti < target.size(); ti++) { //ti: target's iterator
      while (si != target[ti]) {
        // Current stream value does not match the wanted value. It must be strictly less that target[ti].
        // Push and immediately pop unwanted values. This increases si until target value is met.
        ans.push_back(PUSH);
        ans.push_back(POP);
        si++;
      }

      // Current stream value matches the wanted one. Push it to the stack.
      ans.push_back(PUSH);
      si++;
    }

    return ans;
  }
};

int main() {
  Solution s;
  std::vector<int> target_example1 = {1, 3};
  std::vector<std::string> ea_example1 = {"Push", "Push", "Pop", "Push"};
  assert(s.buildArray(target_example1, 3) == ea_example1);

  std::vector<int> target_example2 = {1, 2, 3};
  std::vector<std::string> ea_example2 = {"Push", "Push", "Push"};
  assert(s.buildArray(target_example2, 3) == ea_example2);

  std::vector<int> target_example3 = {1, 2};
  std::vector<std::string> ea_example3 = {"Push", "Push"};
  assert(s.buildArray(target_example3, 4) == ea_example3);
}
