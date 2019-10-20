/**
 * Solution to the "Queen's Attack 2" challenge from HackerRank:
 * https://www.hackerrank.com/challenges/queens-attack-2/problem
 * Created on: 13.09.2019
 * Last modified: 14.09.2019
 * Author: Adam Gralinski (adam@gralin.ski)
 * License: MIT
 */

/*
Note: The board is indexed somewhat counterintuitively in this challenge. The first index is row index,
rising "upwards" instead of "downwards". The second index is the column index, rising "rightwards".
Minimum index is 1, maximum is N.
For example, on the board below queen is placed at (5, 3) and the obstacle at (2, 1). N is 6:
 +------+
6|......|
5|..Q...|
4|......|
3|......|
2|X.....|
1|......|
 +------+
  123456
*/
#include <cstdlib>
#include <iostream>
#include <algorithm>

struct Pos
{
  Pos(int newx = 0, int newy = 0) : x(newx), y(newy) {}
  int x;
  int y;
};

int main(int, char**)
{
  int boardSize, obstaclesCount;
  std::cin >> boardSize >> obstaclesCount;

  // Holds the position of the queen (at index 5) and eight obstacles closest to the queen
  // (at indices 1-4, 6-9).
  //                           7 8 9                        NW  N  NE
  //                  index:   4 5 6     -->   meaning:     W   Q   E
  //                           1 2 3                        SW  S  SE
  // Position 0 is unused.
  Pos board[10];
  std::cin >> board[5].y >> board[5].x;

  // Before processing real obstacles, places 8 "virtual" obstacles at indices 0 / N+1:
  board[0] = {-1, -1}; // nonexistent obstacle.
  board[2] = {board[5].x, 0};
  board[4] = {0, board[5].y};
  board[6] = {boardSize + 1, board[5].y};
  board[8] = {board[5].x, boardSize + 1};

  int dist1 = std::min(board[5].x, board[5].y);
  board[1] = {board[5].x - dist1, board[5].y - dist1};
  int dist3 = std::min(boardSize + 1 - board[5].x, board[5].y);
  board[3] = {board[5].x + dist3, board[5].y - dist3};
  int dist7 = std::min(board[5].x, boardSize + 1 - board[5].y);
  board[7] = {board[5].x - dist7, board[5].y + dist7};
  int dist9 = std::min(boardSize + 1 - board[5].x, boardSize + 1 - board[5].y);
  board[9] = {board[5].x + dist9, board[5].y + dist9};

  Pos obstaclePos;
  for (std::size_t i = 0; i < obstaclesCount; ++i)
  {
    std::cin >> obstaclePos.y >> obstaclePos.x;

    if (obstaclePos.y == board[5].y) // Obstacle lies on the same row as the queen
    {
      // case: obstacle to the right of the queen and to the left of the so-far closest eastern obstacle:
      if ((obstaclePos.x > board[5].x) && (obstaclePos.x < board[6].x))
      {
        board[6].x = obstaclePos.x;
      }
      else if ((obstaclePos.x < board[5].x) && (obstaclePos.x > board[4].x))
      {
        board[4].x = obstaclePos.x;
      }
    }
    else if (obstaclePos.x == board[5].x)
    {
      // case: obstacale to the right of the queen and to the left of the so-far closest western obstacle:
      if ((obstaclePos.y > board[5].y) && (obstaclePos.y < board[8].y))
      {
        board[8].y = obstaclePos.y;
      }
      else if ((obstaclePos.y < board[5].y) && (obstaclePos.y > board[2].y))
      {
        board[2].y = obstaclePos.y;
      }
    }
    else
    {
      int diffX = board[5].x - obstaclePos.x;
      int diffY = board[5].y - obstaclePos.y;

      // case: obstacle lies on the backslash diagonal:
      if (diffX == -diffY)
      {
        if ((obstaclePos.x > board[5].x) && (obstaclePos.x < board[3].x))
        {
          board[3].x = obstaclePos.x;
          board[3].y = obstaclePos.y;
        }
        else if ((obstaclePos.x < board[5].x) && (obstaclePos.x > board[7].x))
        {
          board[7].x = obstaclePos.x;
          board[7].y = obstaclePos.y;
        }
      }
      // case: obstacle lies on the forward slash diagonal:
      else if (diffX == diffY)
      {
        if ((obstaclePos.x > board[5].x) && (obstaclePos.x < board[9].x))
        {
          board[9].x = obstaclePos.x;
          board[9].y = obstaclePos.y;
        }
        else if ((obstaclePos.x < board[5].x) && (obstaclePos.x > board[1].x))
        {
          board[1].x = obstaclePos.x;
          board[1].y = obstaclePos.y;
        }
      }
    }

  }

  int attackedSquaresCount =
      (board[6].x - board[4].x - 2) + (board[8].y - board[2].y - 2) +
      (board[9].x - board[1].x - 2) + (board[7].y - board[3].y - 2);
  std::cout << attackedSquaresCount << std::endl;

  return 0;
}
