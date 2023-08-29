#include <cassert>
#include <string>
#include <vector>
#include <iostream>

class Solution {
public:
  int bestClosingTime(std::string customers) {
    std::vector<int> prefixSumN;
    prefixSumN.resize(customers.size() + 1);
    prefixSumN[0] = 0; // a corner case. There are no prior Ns if a shop is closed immediately.
    for (std::size_t c = 0; c < customers.size(); c++) {
      prefixSumN[c+1] = prefixSumN[c] + (customers[c] == 'N' ? 1 : 0);
    }

    std::vector<int> suffixSumY;
    suffixSumY.resize(customers.size() + 1);
    suffixSumY[customers.size()] = 0; // a corner case. No customers can come after last customers are allowed to come.
    for (int c = customers.size() - 1; c >= 0; c--) {
      suffixSumY[c] = suffixSumY[c+1] + (customers[c] == 'Y' ? 1 : 0);
    }

    int bestIdx = 0;
    int bestIdxVal = prefixSumN[0] + suffixSumY[0];
    for (int i = 1; i < prefixSumN.size(); i++) {
      int penalty = prefixSumN[i] + suffixSumY[i];
      if (penalty < bestIdxVal) {
        bestIdx = i;
        bestIdxVal = penalty;
      }
    }

    // For visual confirmation that values are calculated correctly:
    /*
    std::cout << customers << "\n";
    for (const auto& n: prefixSumN) {
      std::cout << n;
    }
    std::cout << "\n";
    for (const auto& y: suffixSumY) {
      std::cout << y;
    }
    std::cout << "\n";
    for (int i = 0; i < prefixSumN.size(); i++) {
      std::cout << (i == bestIdx ? '^' : ' ');
    }
    std::cout << "\n";
    */

    return bestIdx;
  }
};

int main() {
  Solution s{};

  assert(s.bestClosingTime("YYNY") == 2);
  assert(s.bestClosingTime("NNNNN") == 0);
  assert(s.bestClosingTime("YYYY") == 4);
  assert(s.bestClosingTime("NYYYNNNYNN") == 4);
}
