/**
 * Solution to the "Equal" challenge form HackerRank:
 * https://www.hackerrank.com/challenges/equal/problem
 * Created on: 24.10.2019
 * Last modified: 24.10.2019
 * Author: Adam Gralinski (adam@gralin.ski)
 * License: MIT
 */

#include <algorithm>
#include <iostream>
#include <vector>

long long getEqualizationStepsCount(const std::vector<int>& people)
{
  long long stepsCount = 0;
  for (auto& p : people)
  {
    int tmp = p;
    stepsCount += tmp / 5;
    tmp %= 5;
    stepsCount += tmp / 2;
    tmp %= 2;
    if (tmp == 1)
    {
      stepsCount += 1;
    }
  }
  return stepsCount;
}

int main(int, char**)
{
  int T;
  std::cin >> T;
  for (int t = 0; t < T; ++t)
  {
    std::size_t noOfPeople;
    std::cin >> noOfPeople;
    std::vector<int> people(noOfPeople);
    for (std::size_t n = 0; n < noOfPeople; ++n)
    {
      std::cin >> people[n];
    }
    std::sort(people.begin(), people.end());

    int minimum = people[0];
    std::transform(people.begin(), people.end(), people.begin(),
        [minimum](int p) -> int { return p - minimum; });

    std::vector<int> people1(people);
    std::transform(people1.begin(), people1.end(), people1.begin(), [](int p) -> int { return p + 1; });
    std::vector<int> people2(people);
    std::transform(people2.begin(), people2.end(), people2.begin(), [](int p) -> int { return p + 2; });

    long long operations0 = getEqualizationStepsCount(people);
    long long operations1 = getEqualizationStepsCount(people1);
    long long operations2 = getEqualizationStepsCount(people2);

    std::cout << std::min(operations0, std::min(operations1, operations2)) << "\n";
  }
  return 0;
}
