/**
 * Solution to the challenge #011 of Project Euler series on HackerRank:
 * https://www.hackerrank.com/contests/projecteuler/challenges/euler011
 * Created on: 24-08-2018
 * Last Modified: 24-08-2018
 * Author: Adam Gralinski (adam@gralin.ski)
 * License: MIT
 */

#include <algorithm>
#include <cstdint>
#include <iostream>
#include <vector>
#include <string>

const int GRID_SIZE = 20;
const int ADJACENCY_LENGTH = 4;

// Invariants: 0 <= Every number in a grid <= 100

int main()
{
  std::vector<std::vector<int>> grid;

  // Reads the input:
  grid.resize(GRID_SIZE);
  for (int y = 0; y < GRID_SIZE; ++y)
  {
    grid[y].resize(GRID_SIZE);
    for (int x = 0; x < GRID_SIZE; ++x)
    {
      std::cin >> grid[y][x];
    }
  }

  // Calculates the answer:
  int maxProduct = 0;
  // Checks horizontal lines:
  for (int y = 0; y < GRID_SIZE; ++y)
  {
    for (int x = 0; x <= GRID_SIZE - ADJACENCY_LENGTH; ++x)
    {
      int product = grid[y][x] * grid[y][x+1] * grid[y][x+2] * grid[y][x+3];
      maxProduct = std::max(maxProduct, product);

    }
  }

  // Checks vertical lines:
  for (int y = 0; y <= GRID_SIZE - ADJACENCY_LENGTH; ++y)
  {
    for (int x = 0; x < GRID_SIZE; ++x)
    {
      int product = grid[y][x] * grid[y+1][x] * grid[y+2][x] * grid[y+3][x];
      maxProduct = std::max(maxProduct, product);
    }
  }

  // Checks diagonal lines:
  for (int y = 0; y <= GRID_SIZE - ADJACENCY_LENGTH; ++y)
  {
    for (int x = 0; x <= GRID_SIZE - ADJACENCY_LENGTH; ++x)
    {
      int product = grid[y][x] * grid[y+1][x+1] * grid[y+2][x+2] * grid[y+3][x+3];
      maxProduct = std::max(maxProduct, product);
      product = grid[y][x+3] * grid[y+1][x+2] * grid[y+2][x+1] * grid[y+3][x];
      maxProduct = std::max(maxProduct, product);
    }
  }

  std::cout << maxProduct << std::endl;

  return 0;
}
