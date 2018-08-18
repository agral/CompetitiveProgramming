/**
 * Solution to the challenge #002 of Project Euler series on HackerRank:
 * https://www.hackerrank.com/contests/projecteuler/challenges/euler002
 * Created on: 18-08-2018
 * Last Modified: 18-08-2018
 * Author: Adam Gralinski (adam@gralin.ski)
 * License: MIT
 */

#include <iostream>
#include <cstdint>
#include <vector>

std::vector<uint64_t> fibSequence{1, 1, 2};
std::vector<uint64_t> answers{0, 0, 0}; // holds sum of even Fibonacci entries whose 1-index does not exceed k.
const uint64_t maxN = 4e16;

int main()
{
  // Precalculates the data until maxN is exceeded:
  uint64_t k = answers.size() - 1; // index of last calculated fibonacci number & answer
  while (fibSequence[k] < maxN)
  {
    fibSequence.push_back(fibSequence[k-1] + fibSequence[k]);
    answers.push_back(answers[k] + (fibSequence[k] % 2 == 0 ? fibSequence[k] : 0));
    k += 1;
  }

  int T;
  std::cin >> T;
  for (int t = 0; t < T; ++t)
  {
    uint64_t N;
    std::cin >> N;

    int n = 1; // This challenge assumes 1-indexing; I threw an extra "1" as zeroth entry.
    while (fibSequence[n] < N)
    {
      n += 1; // Linear search, could be optimized to binary if it timeouts.
    }
    std::cout << answers[n] << std::endl;
  }

  return 0;
}
