/**
 * Solution to the "SherlockAndAnagrams" challenge from HackerRank:
 * https://www.hackerrank.com/challenges/sherlock-and-anagrams/problem
 * Created on: 05.10.2019
 * Last modified: 22.10.2019
 * Author: Adam Gralinski (adam@gralin.ski)
 * License: MIT
 */

#include <array>
#include <algorithm>
#include <iostream>
#include <string>

bool isAnagram(std::string const& first, std::string const& second)
{
  std::array<int, 256> histogram = {0};
  for (auto const &f : first)
  {
    histogram[static_cast<unsigned char>(f)] += 1;
  }
  for (auto const &s : second)
  {
    histogram[static_cast<unsigned char>(s)] -= 1;
  }

  return std::all_of(histogram.begin(), histogram.end(), [](int num){return num == 0;});
}

int main(int, char**)
{
  int T;
  std::cin >> T;
  for (int t = 0; t < T; ++t)
  {
    std::string s;
    std::cin >> s;
    std::size_t count = 0;

    // Index of first substring's beginning - from 0 to semi-last character of the input string:
    for (std::size_t beginning = 0; beginning < s.size() - 1; ++beginning)
    {
      // Length of the substrings considered in this iteration - from 1 to such a length, that the first substring
      // will terminate at most at the semi-last character of the input string
      // (so that the second substring can still be created and be different from the first one):
      for (std::size_t substring_length = 1; beginning + substring_length < s.size(); ++substring_length)
      {
        auto first_substring = s.substr(beginning, substring_length);

        // The offset, in chars, between the beginnings of the first and the second substrings:
        for (std::size_t offset = 1; beginning + substring_length + offset <= s.size(); ++offset)
        {
          // The substrings are of the same length (otherwise they would trivially not be anagrams of each other):
          auto second_substring = s.substr(beginning + offset, substring_length);
          if (isAnagram(first_substring, second_substring))
          {
            ++count;
          }
        }
      }
    }
    std::cout << count << std::endl;
  }
  return 0;
}
