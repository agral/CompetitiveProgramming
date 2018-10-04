/**
 * Solution to the challenge #022 of Project Euler series on HackerRank:
 * https://www.hackerrank.com/contests/projecteuler/challenges/euler022
 * Created on: 20-09-2018
 * Last Modified: 20-09-2018
 * Author: Adam Gralinski (adam@gralin.ski)
 * License: MIT
 */

#include <algorithm>
#include <iostream>
#include <cstdint>
#include <string>
#include <vector>

std::vector<std::string> data;

uint64_t GetScore(const std::string& entry, uint64_t multiplier)
{
  uint64_t baseScore = 0;
  for (auto& c : entry)
  {
    baseScore += (c - 'A' + 1);
  }

  return baseScore * multiplier;
}

void ReadData()
{
  std::size_t N;
  std::cin >> N;
  data.resize(N);
  for (std::size_t n = 0; n < N; ++n)
  {
    std::cin >> data[n];
  }

  std::sort(data.begin(), data.end());
}

int main()
{
  ReadData();
  int Q;
  std::cin >> Q;
  for (int q = 0; q < Q; ++q)
  {
    std::string query;
    std::cin >> query;

    uint64_t index;
    for (index = 0; (index < data.size() && (data[index] != query)); ++index)
    {
    }

    std::cout << GetScore(query, index + 1) << std::endl;
  }
  return 0;
}
