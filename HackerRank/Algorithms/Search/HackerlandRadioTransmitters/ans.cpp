/**
 * Solution to the "Hackerland Radio Transmitters" challenge from HackerRank:
 * https://www.hackerrank.com/challenges/hackerland-radio-transmitters/problem
 * Created on: 22.06.2019
 * Last modified: 04.09.2019
 * Author: Adam Gralinski (adam@gralin.ski)
 * License: MIT
 */

#include <algorithm>
#include <iostream>
#include <vector>

int main(int, char**)
{
  // Reads in the testcase's data:
  //   N: number of houses in Hackerland,            1 <= N <= 10e5
  //   K: range of each individual transmitter,      1 <= K <= 10e5
  //   houses: array of houses' addresses (N times), 1 <= house(n) <= 10e5
  int N, K;
  std::cin >> N >> K;
  std::vector<int> houses(N);
  for (int n = 0; n < N; ++n)
  {
    std::cin >> houses[n];
  }
  // Sorts the array of houses' addresess in increasing order.
  std::sort(houses.begin(), houses.end());

  // Greedily assigns a minimum number of transmitters to cover every house in Hackerland.
  // Does so by:
  int totalTransmittersCount = 0;
  int nextUncovered = 0;
  while (nextUncovered < N)
  {
    int installTransmitterHere = nextUncovered;
    while ((houses[installTransmitterHere] - houses[nextUncovered] <= K) && (installTransmitterHere < N))
    {
      installTransmitterHere += 1;
    }

    // Goes one step back since the loop above "overshoots":
    installTransmitterHere -= 1;
    totalTransmittersCount += 1;

    // This transmitter will also cover other houses residing on this position,
    // and potentially other houses to the right (in its range limit):
    nextUncovered = installTransmitterHere + 1;
    while ((houses[installTransmitterHere] + K >= houses[nextUncovered]) && (nextUncovered < N))
    {
      nextUncovered++;
    }
  }

  std::cout << totalTransmittersCount << std::endl;
  return 0;
}
