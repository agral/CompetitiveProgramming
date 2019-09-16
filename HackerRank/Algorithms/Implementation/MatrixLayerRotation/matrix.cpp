/**
 * Solution to the "Matrix Layer Rotation" challenge from HackerRank:
 * https://www.hackerrank.com/challenges/matrix-rotation-algo/problem
 * Created on: 16.09.2019
 * Last modified: 16.09.2019
 * Author: Adam Gralinski (adam@gralin.ski)
 * License: MIT
**/

#include <algorithm>
#include <iostream>
#include <vector>

int main(int, char**)
{
  std::size_t width, height; // <2, 300>, also min(width, height) % 2 == 0.
  std::size_t rotation; // <1, 1e9)
  std::vector<std::vector<int>> matrix;

  // Reads the width, height, rotation and initial contents of the matrix:
  std::cin >> height >> width >> rotation;
  matrix.resize(height);
  for (auto& row : matrix)
  {
    row.resize(width);
    for (auto& elem : row)
    {
      std::cin >> elem;
    }
  }

  // Performs the rotation:
  for (std::size_t ring = 0; ring < std::min(width, height) / 2; ++ring)
  {
    // Unravels the ring:
    std::size_t ring_len = 2 * (width - 1 - (2 * ring) + height - 1 - (2 * ring));
    std::size_t half_ring_len = ring_len / 2;
    std::vector<int> buf(ring_len, -1);
    std::vector<int>::iterator ib = buf.begin();
    for (std::size_t y = ring; y < height - ring - 1; ++y)
    {
      *ib = matrix[y][ring];
      *(ib + half_ring_len) = matrix[height - 1 - y][width - 1 - ring];
      ib++;
    }
    for (std::size_t x = ring; x < width - ring - 1; ++x)
    {
      *ib = matrix[height - 1 - ring][x];
      *(ib + half_ring_len) = matrix[ring][width - 1 - x];
      ib++;
    }

    // Creates a rotated copy of the unraveled ring:
    std::vector<int> rotatedBuf(ring_len, -1);
    std::vector<int>::iterator irb = rotatedBuf.begin() + (rotation % ring_len);
    for (ib = buf.begin(); ib != buf.end();)
    {
      *irb = *ib;
      ++ib;
      ++irb;
      if (irb == rotatedBuf.end())
      {
        irb = rotatedBuf.begin();
      }
    }

    // Rewrites the contents of the matrix using a rotated unraveled ring:
    irb = rotatedBuf.begin();
    for (std::size_t y = ring; y < height - ring - 1; ++y)
    {
      matrix[y][ring] = *irb;
      matrix[height - 1 - y][width - 1 - ring] = *(irb + half_ring_len);
      irb++;
    }
    for (std::size_t x = ring; x < width - ring - 1; ++x)
    {
      matrix[height - 1 - ring][x] = *irb;
      matrix[ring][width - 1 - x] = *(irb + half_ring_len);
      irb++;
    }
  }

  // Finally, prints out the contents of the array, rotated now:
  for (auto &row : matrix)
  {
    for (auto &c : row)
    {
      std::cout << c << " ";
    }
    std::cout << std::endl;
  }

  return 0;
}
