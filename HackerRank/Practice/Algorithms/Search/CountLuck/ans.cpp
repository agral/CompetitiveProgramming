/**
 * Solution to the "Count Luck" challenge form HackerRank:
 * https://www.hackerrank.com/challenges/count-luck/problem
 * Created on: 17.10.2019
 * Last modified: 20.10.2019
 * Author: Adam Gralinski (adam@gralin.ski)
 * License: MIT
 */

#include <iostream>
#include <vector>

// The following chars may appear on the challenge's map:
auto const TREE = 'X';
//auto const PATH = '.'; (unused)
auto const HERO = 'M';
auto const GOAL = '*';

auto const SUCCESS = "Impressed";
auto const FAILURE = "Oops!";

struct IntermediatePosition
{
  IntermediatePosition(std::size_t newY, std::size_t newX, int newHeadingFrom, std::size_t newCount)
  : y(newY), x(newX), headingFrom(newHeadingFrom), crossroadsEncounteredSoFar(newCount) {}
  std::size_t y;
  std::size_t x;
  int headingFrom; // 2, 4, 5, 6 or 8.
  std::size_t crossroadsEncounteredSoFar;
};

int GetCrossroadsCount(const std::vector<std::vector<char>> &map)
{
  // Pinpoints the heroes' location:
  std::size_t heroX, heroY;
  bool isHeroFound = false;
  for (std::size_t y = 0; ((y < map.size()) && (!isHeroFound)); ++y)
  {
    for (std::size_t x = 0; ((x < map[y].size()) && (!isHeroFound)); ++x)
    {
      if (map[y][x] == HERO)
      {
        heroX = x;
        heroY = y;
        isHeroFound = true;
      }
    }
  }

  if (!isHeroFound)
  {
    std::cerr << "Malformed map: HERO not found.\n";
    return 0;
  }

  IntermediatePosition startPosition{heroY, heroX, 5, 0};
  std::vector<IntermediatePosition> positions;
  positions.push_back(startPosition);

  while(!positions.empty())
  {
    IntermediatePosition currentPosition = positions.back();
    positions.pop_back();

    if (map[currentPosition.y][currentPosition.x] == GOAL)
    {
      return currentPosition.crossroadsEncounteredSoFar;
    }

    bool canHeadDown = false, canHeadLeft = false, canHeadRight = false, canHeadUp = false;
    std::size_t availableHeadingsCount = 0;

    if ((currentPosition.headingFrom != 2) && (map[currentPosition.y + 1][currentPosition.x] != TREE))
    {
      canHeadDown = true;
      availableHeadingsCount += 1;
    }
    if ((currentPosition.headingFrom != 4) && (map[currentPosition.y][currentPosition.x - 1] != TREE))
    {
      canHeadLeft = true;
      availableHeadingsCount += 1;
    }
    if ((currentPosition.headingFrom != 6) && (map[currentPosition.y][currentPosition.x + 1] != TREE))
    {
      canHeadRight = true;
      availableHeadingsCount += 1;
    }
    if ((currentPosition.headingFrom != 8) && (map[currentPosition.y - 1][currentPosition.x] != TREE))
    {
      canHeadUp = true;
      availableHeadingsCount += 1;
    }

    std::size_t crossroadsEncounteredSoFar = currentPosition.crossroadsEncounteredSoFar;
    if (availableHeadingsCount >= 2)
    {
      crossroadsEncounteredSoFar += 1;
    }

    if (canHeadDown)
    {
      positions.push_back(IntermediatePosition(
          currentPosition.y + 1, currentPosition.x, 8, crossroadsEncounteredSoFar));
    }
    if (canHeadLeft)
    {
      positions.push_back(IntermediatePosition(
          currentPosition.y, currentPosition.x - 1, 6, crossroadsEncounteredSoFar));
    }
    if (canHeadRight)
    {
      positions.push_back(IntermediatePosition(
          currentPosition.y, currentPosition.x + 1, 4, crossroadsEncounteredSoFar));
    }
    if (canHeadUp)
    {
      positions.push_back(IntermediatePosition(
          currentPosition.y - 1, currentPosition.x, 2, crossroadsEncounteredSoFar));
    }
  }

  std::cerr << "Error: Entire forest traversed and no GOAL has been found.\n";
  return 0;
}

int main(int, char**)
{
  int T;
  std::cin >> T;
  for (int t = 0; t < T; ++t)
  {
    std::size_t height, width;
    std::cin >> height >> width;
    std::vector<std::vector<char>> map;
    map.resize(height + 2);
    map[0].resize(width + 2);
    map[height + 1].resize(width + 2);
    for (std::size_t h = 1; h <= height; ++h)
    {
      map[h].resize(width + 2);
      for (std::size_t w = 1; w <= width; ++w)
      {
        std::cin >> map[h][w];
      }
    }
    for (std::size_t h = 0; h <= height; ++h)
    {
      map[h][0] = TREE;
      map[height + 1 - h][width + 1] = TREE;
    }
    for (std::size_t w = 0; w <= width; ++w)
    {
      map[0][width + 1 - w] = TREE;
      map[height+1][w] = TREE;
    }
    int ronGuess;
    std::cin >> ronGuess;

    std::cout << ( GetCrossroadsCount(map) == ronGuess ? SUCCESS : FAILURE ) << "\n";
  }
  return 0;
}
