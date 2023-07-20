#include <cassert>
#include <vector>
#include <iostream>

using std::vector;

class Solution {
public:
  vector<int> asteroidCollision(vector<int>& asteroids) {
    vector<int> ans{};
    ans.push_back(asteroids[0]);
    for (int a = 1; a < asteroids.size(); a++) {
      if (asteroids[a] < 0) {
        bool is_new_one_destroyed = false;
        int idx = ans.size() - 1;
        while (idx >= 0 && !is_new_one_destroyed && ans[idx] > 0) {
          // a collision happens. Possible results:
          if (ans[idx] > -asteroids[a]) {
            // Case A: the new left-facing asteroid is smaller than this right-facing asteroid.
            // the new asteroid is destroyed. Nothing more happens.
            is_new_one_destroyed = true;
          }
          else if (ans[ans.size() - 1] < -asteroids[a]) {
            // Case B: the new left-facing asteroid is equal-or-bigger than right-facing asteroid.
            // the right-facing asteroid is destroyed.
            ans.pop_back();
            idx -= 1;
          }
          else {
            // Case C: the new left-facing asteroid is equal in mass to the right-facing asteroid.
            // Both asteroids are destroyed:
            ans.pop_back();
            is_new_one_destroyed = true;
            // we don't care about idx anymore.
          }
        }
        if (!is_new_one_destroyed) {
          ans.push_back(asteroids[a]);
        }
      }
      else { // this asteroid is a right-facing one. Just add it to the asteroid queue, no collision is possible.
        ans.push_back(asteroids[a]);
      }
    }
    return ans;
  }
};

int main() {
  Solution s{};
  vector<int> asteroids1 = {5, 10, -5};
  vector<int> expected1 = {5, 10};
  assert(s.asteroidCollision(asteroids1) == expected1);

  vector<int> asteroids2 = {8, -8};
  vector<int> expected2 = {};
  assert(s.asteroidCollision(asteroids2) == expected2);

  vector<int> asteroids3 = {10, 2, -5};
  vector<int> expected3 = {10};
  assert(s.asteroidCollision(asteroids3) == expected3);

  vector<int> asteroids4 = {-20, -10, 30, 20};
  vector<int> expected4 = {-20, -10, 30, 20};
  assert(s.asteroidCollision(asteroids4) == expected4);

  vector<int> asteroids5 = {-1, 1, 2, -3, 4, 2, 1, -2, 5};
  vector<int> expected5 = {-1, -3, 4, 5};
  assert(s.asteroidCollision(asteroids5) == expected5);

  vector<int> asteroids6 = {-1, 1, 2, -3, 4, 2, 1, -2, 5, -7};
  vector<int> expected6 = {-1, -3, -7};
  assert(s.asteroidCollision(asteroids6) == expected6);

  vector<int> asteroids7 = {1, -2, -2, -2};
  vector<int> expected7 = {-2, -2, -2};
  auto result = s.asteroidCollision(asteroids7);
  for (const auto& a: result) {
    std::cout << a << ", ";
  }
  std::cout << std::endl;

  assert(s.asteroidCollision(asteroids7) == expected7);
}
