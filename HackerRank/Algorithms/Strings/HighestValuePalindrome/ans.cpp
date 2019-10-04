/**
 * Solution to the "Highest Value Palindrome" challenge from HackerRank:
 * https://www.hackerrank.com/challenges/richie-rich/problem
 * Created on: 04.10.2019
 * Last modified: 04.10.2019
 * Author: Adam Gralinski (adam@gralin.ski)
 * License: MIT
 */

#include <iostream>
#include <vector>

const auto HIGHEST_CHAR = '9';

bool inplaceMakePalindromic(std::vector<char>& candidate, std::size_t noOfChangesAllowed)
{
  std::vector<bool> isChanged(candidate.size(), false);
  std::size_t changesDone = 0;

  for (int b = candidate.size() / 2 - 1, e = (candidate.size() + 1) / 2 ; b >= 0; --b, ++e)
  {
    if (candidate[b] != candidate[e])
    {
      if (changesDone == noOfChangesAllowed)
      {
        return false; // since we can not afford this change and a sequence won't be palindromic without it.
      }

      if (candidate[b] > candidate[e])
      {
        candidate[e] = candidate[b];
        isChanged[e] = true;
      }
      else // implies candidate[b] < candidate[e]
      {
        candidate[b] = candidate[e];
        isChanged[b] = true;
      }
      changesDone += 1;
    }
  }

  std::size_t b = 0, e = candidate.size() - 1;
  while ((b < e) && (changesDone < noOfChangesAllowed))
  {
    if ( (candidate[b] < HIGHEST_CHAR) && ((noOfChangesAllowed - changesDone >= 2) || (isChanged[b] || isChanged[e])) )
    {
      candidate[b] = HIGHEST_CHAR;
      if (!isChanged[b])
      {
        changesDone += 1;
        isChanged[b] = true;
      }

      candidate[e] = HIGHEST_CHAR;
      if (!isChanged[e])
      {
        changesDone += 1;
        isChanged[e] = true;
      }
    }
    b += 1;
    e -= 1;
  }

  // If the string is odd in size and there is still at least one change left, switches the middle char to '9':
  if ((changesDone < noOfChangesAllowed) && (b == e))
  {
    candidate[candidate.size() / 2] = HIGHEST_CHAR;
    changesDone += 1;
  }

  return true;
}

int main(int, char**)
{
  std::size_t noOfChars, noOfChangesAllowed;
  std::vector<char> input;

  std::cin >> noOfChars >> noOfChangesAllowed;
  input.resize(noOfChars);
  for (auto& c : input)
  {
    std::cin >> c;
  }

  if (!inplaceMakePalindromic(input, noOfChangesAllowed))
  {
    std::cout << "-1\n";
  }
  else
  {
    for (auto& c : input)
    {
      std::cout << c;
    }
    std::cout << '\n';
  }

  return 0;
}
