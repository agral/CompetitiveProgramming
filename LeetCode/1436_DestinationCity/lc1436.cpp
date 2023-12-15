#include <cassert>
#include <string>
#include <vector>
#include <unordered_set>

class Solution {
public:
  std::string destCity(std::vector<std::vector<std::string>>& paths) {
    std::unordered_set<std::string> cities;
    for (const auto& v: paths) {
      // Store all departure cities in a set
      cities.insert(v[0]);
    }
    for (const auto& v: paths) {
      // Check all arrival cities. One of them will not be in a departure cities' set.
      // This is the final destination - return it.
      if (cities.find(v[1]) == cities.end()) {
        return v[1];
      }
    }
    throw;
  }
};

int main() {
  Solution s;
  {
    std::vector<std::vector<std::string>> paths = {
      {"London", "New York"},
      {"New York", "Lima"},
      {"Lima", "Sao Paulo"}
    };
    assert(s.destCity(paths) == "Sao Paulo");
  }
  {
    std::vector<std::vector<std::string>> paths = {{"B", "C"}, {"D", "B"}, {"C", "A"}};
    assert(s.destCity(paths) == "A");
  }
  {
    std::vector<std::vector<std::string>> paths = {{"A", "Z"}};
    assert(s.destCity(paths) == "Z");
  }
}
