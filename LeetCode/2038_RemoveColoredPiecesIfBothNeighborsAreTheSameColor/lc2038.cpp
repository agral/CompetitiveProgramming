#include <cassert>
#include <string>

class Solution {
public:
  bool winnerOfGame(std::string colors) {
    // colors contains only letters: A or B.
    // Alice's and Bob's moves consist of finding a triplet of their respective letters and removing
    // the middle one. Thus, this task boils down to counting triplets (AAA and BBB),
    // Alice wins if count(AAA) > count(BBB); otherwise she loses.

    int score_alice = 0;
    int score_bob = 0;
    for (int i = 0; i + 2 < colors.size(); i++) {
      if (colors[i] == 'A' && colors[i+1] == 'A' && colors[i+2] == 'A') {
        score_alice++;
      }
      else if (colors[i] == 'B' && colors[i+1] == 'B' && colors[i+2] == 'B') {
        score_bob++;
      }
    }
    return score_alice > score_bob;
  }
};

int main() {
  Solution s;
  assert(s.winnerOfGame("AAABABB") == true);
  assert(s.winnerOfGame("AA") == false);
  assert(s.winnerOfGame("ABBBBBBAAA") == false);
}
