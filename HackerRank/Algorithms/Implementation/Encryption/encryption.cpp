/**
 * Solution to the "Encryption" challenge from HackerRank:
 * https://www.hackerrank.com/challenges/encryption/problem
 * Created on: 01.09.2019
 * Last modified: 01.09.2019
 * Author: Adam Gralinski (adam@gralin.ski)
 * License: MIT
 */

#include <bits/stdc++.h>

std::string encrypt(const std::string& in)
{
  std::string ans = in;

  // Removes spaces from the input string:
  ans.erase(std::remove_if(ans.begin(), ans.end(), isspace), ans.end());

  int dimY = std::round(std::sqrt(ans.length()));
  int dimX = (dimY * dimY >= ans.length()) ? dimY : dimY + 1;

  std::stringstream encrypted;

  for (std::size_t x = 0; x < dimX; ++x)
  {
    if (x > 0)
    {
      encrypted << " ";
    }
    for (std::size_t y = 0; y < dimY; ++y)
    {
      std::size_t k = y * dimX + x;
      if (k < ans.length())
      {
        encrypted << ans[y * dimX + x];
      }
    }
  }

  return encrypted.str();
}

int main(int, char**)
{
  std::string input;
  std::getline(std::cin, input);
  std::cout << encrypt(input) << std::endl;
  return 0;
}
