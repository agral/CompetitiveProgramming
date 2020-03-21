/**
 * Solution to the "Diagonal Difference" challenge from HackerRank:
 * https://www.hackerrank.com/challenges/diagonal-difference/problem
 * Created on: 21.03.2020
 * Last modified: 21.03.2020
 * Author: Adam Graliński (adam@gralin.ski)
 * License: MIT
 */

#include <iostream>
#include <vector>

int main(int, char**)
{
  std::size_t N;
  std::cin >> N;

  std::vector<std::vector<int>> v;
  v.resize(N);
  for (std::size_t row = 0; row < N; ++row) {
    v[row].resize(N);
    for (std::size_t col = 0; col < N; ++col) {
      std::cin >> v[row][col];
    }
  }

  int diag1{}, diag2{};
  for (std::size_t k = 0; k < N; ++k) {
    diag1 += v[k][k];
    diag2 += v[k][N-k-1];
  }

  std::cout << std::abs(diag1 - diag2) << "\n";
  return 0;
}
