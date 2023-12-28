#include <cassert>
#include <string>
#include <vector>

class Solution {
private:
  static constexpr int kMAX = 101;

    // dp[i][k]: length of optimal compression of s[i...n)
    // with at most k deletions
  std::vector<std::vector<int>> dp;

public:
  int getLengthOfOptimalCompression(std::string s, int k) {
    dp.clear();
    dp.resize(s.size(), std::vector<int>(k+1, kMAX));
    return calc(s, 0, k);
  }

private:
  int calc(const std::string& s, int i, int k) {
    if (k < 0) {
      return kMAX; // impossible to do it with k<0, so return "BIG_INT".
    }
    if ((s.size() == i) || (s.size() - i <= k)) {
      return 0;
    }
    if (dp[i][k] != kMAX) {
      return dp[i][k]; // if already calculated, return it. Otherwise, finally, do calculate:
    }

    int maxFreq = 0;
    std::vector<int> count(2 * kMAX); // probably overshooting, but kMAX is not enough. kMAX+20 looks OK,
                                      // but let's not take any chances ;-) - one WA already obtained.
    for (int a = i; a < s.size(); ++a) {
      maxFreq = std::max(maxFreq, ++count[s[a]]);
      dp[i][k] = std::min(dp[i][k], getEncodedLength(maxFreq) +
                          calc(s, a + 1, k + i + maxFreq - a - 1));
    }

    return dp[i][k];
  }

  int getEncodedLength(int freq) {
    if (freq == 1) {
      return 1; // e.g. "a"
    }
    if (freq < 10) {
      return 2; // e.g. "o2", "a3", ..., "a9"
    }
    if (freq < 100) {
      return 3; // e.g. "a10", "a11", ..., "a99"
    }
    return 4; // "a100". As max string size is 100, this is the only case with encoded length of 4.
  }
};

int main() {
  Solution s;
  assert(s.getLengthOfOptimalCompression("aaabcccd", 2) == 4);
  assert(s.getLengthOfOptimalCompression("aabbaa", 2) == 2);
  assert(s.getLengthOfOptimalCompression("aaaaaaaaaaa", 0) == 3);
  assert(s.getLengthOfOptimalCompression("aaaaabaaaaafffwfff", 2) == 5);
}
