/**
 * Solution to the challenge #004 of Project Euler series on HackerRank:
 * https://www.hackerrank.com/contests/projecteuler/challenges/euler004
 * Created on: 19-08-2018
 * Last Modified: 19-08-2018
 * Author: Adam Gralinski (adam@gralin.ski)
 * License: MIT
 */

/**
 * A palindromic number reads the same both ways.
 * The smallest 6-digit palindrome that is also a product of two 3-digit numbers
 * is 101101 = 143 * 707.
 *
 * Find the largest palindrome made from the product of two 3-digit numbers
 * strictly less than N.
 */

#include <iostream>
#include <algorithm>
#include <cstdint>
#include <vector>
#include <limits>

std::vector<int> results;
std::vector<int> uniqueResults;

bool IsPalindrome(int n)
{
  int originalN = n;
  int reverse = 0;
  while (n > 0)
  {
    int lastDigit = n % 10;
    n /= 10;
    reverse *= 10;
    reverse += lastDigit;
  }

  return (reverse == originalN);
}

void FindPalindromicResults()
{
  for (int a = 100; a < 1000; ++a)
  {
    for (int b = a; b < 1000; ++b)
    {
      int result = a * b;
      if (IsPalindrome(result))
      {
        results.push_back(result);
      }
    }
  }
}

int main()
{
  FindPalindromicResults();
  std::sort(results.begin(), results.end());
  int lastResult = 0;
  for (auto& result : results)
  {
    if (result > lastResult)
    {
      uniqueResults.push_back(result);
      lastResult = result;
    }
  }
  uniqueResults.push_back(std::numeric_limits<int>::max());

  int T;
  std::cin >> T;
  for (int t = 0; t < T; ++t)
  {
    int N;
    std::cin >> N;
    int k = 0;
    while (uniqueResults[k] < N)
    {
      k++;
    }
    std::cout << uniqueResults[k-1] << std::endl;
  }
  return 0;
}
