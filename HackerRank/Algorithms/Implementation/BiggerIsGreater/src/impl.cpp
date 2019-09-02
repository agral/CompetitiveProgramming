#include "impl.hpp"

#include <iostream>
#include <algorithm>
#include <vector>

const std::string NO_ANSWER = "no answer";

std::string getNextLexicographicPermutation(const std::string& input)
{
  if (input.length() <= 1)
  {
    return NO_ANSWER;
  }

  // Finds the position of the pivot (the first decreasing element in reversed sequence):
  int pivot = input.length() - 2;
  while ((pivot >= 0) && (input[pivot] >= input[pivot + 1]))
  {
    pivot -= 1;
  }

  if (pivot == -1)
  {
    return NO_ANSWER;
  }

  // Finds the successor to the pivot:
  std::size_t successor = input.length() - 1;
  while (input[successor] <= input[pivot])
  {
    successor -= 1;
  }

  std::vector<char> ans(input.begin(), input.end());
  std::swap(ans[pivot], ans[successor]);
  std::reverse(ans.begin() + pivot + 1, ans.end());

  return std::string(ans.begin(), ans.end());
}

