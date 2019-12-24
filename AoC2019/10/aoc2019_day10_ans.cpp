/**
 * Name: aoc2019_day10_ans
 * Description: Answers the AoC2019 day#10 challenge
 * Created on: 10.12.2019
 * Last modified: 11.12.2019
 * Author: Adam Gralinski (adam@gralin.ski)
 * License: CC0
 */

#include <algorithm>
#include <cmath> //atan2
#include <fstream>
#include <iostream>
#include <sstream>
#include <string>
#include <vector>
#include <numeric> //std::gcd
#include <utility>
#include <set>
#include <map>

typedef std::pair<int, int> Position;
const auto INPUT_FILENAME = std::string("aoc2019_day10_input.txt");
//const auto INPUT_FILENAME = std::string("very_simple_input.txt");
//const auto INPUT_FILENAME = std::string("simple_input.txt");
//const auto INPUT_FILENAME = std::string("large_simple_input.txt");
const auto ASTEROID = '#';
const auto EMPTY_SPACE = '.';
const auto VAPORIZE_LIMIT= 200;
const auto PI = 3.14159265359;
std::vector<std::vector<char>> map;

std::ostream& operator<< (std::ostream& os, const Position& p)
{
  os << "Position (" << p.second << ", " << p.first << ")";
  return os;
}

Position normalize(Position p)
{
  if (p.first == 0)
  {
    return {0, (p.second > 0 ? 1 : -1)};
  }
  else if (p.second == 0)
  {
    return {(p.first > 0 ? 1 : -1), 0};
  }

  auto gcd = std::gcd(p.first, p.second);
  return {p.first / gcd, p.second / gcd};
}

double vectorToAngle(const Position& v)
{
  //return std::atan2(v.first, v.second) * 180 / PI;

  // Normalized so that angle(NORTH) == 0 and grows clockwise.
  double rv = ((0.5 * PI) - std::atan2(v.first, v.second)) * (180.0 / PI);
  return (rv >= 0.0 ? rv : rv + 360.0);
}

int main(int, char**)
{
  std::ifstream inputFile;
  inputFile.open(INPUT_FILENAME);
  if (inputFile.is_open())
  {
    std::string line;
    while (std::getline(inputFile, line))
    {
      map.push_back(std::vector<char>(line.begin(), line.end()));
    }

    std::size_t maxAsteroidsVisible = 0;
    Position imsPos;
    for (std::size_t h = 0; h < map.size(); ++h)
    {
      for (std::size_t w = 0; w < map[h].size(); ++w)
      {
        if (map[h][w] == ASTEROID)
        {
          std::set<Position> directionsToAsteroids;
          for (std::size_t otherH = 0; otherH < map.size(); ++otherH)
          {
            for (std::size_t otherW = 0; otherW < map[otherH].size(); ++otherW)
            {
              if ((h == otherH) && (w == otherW))
              {
                continue;
              }
              if (map[otherH][otherW] == ASTEROID)
              {
                Position p = normalize({otherH - h, otherW - w});
                directionsToAsteroids.insert(p);
              }
            }
          }
          if (directionsToAsteroids.size() > maxAsteroidsVisible)
          {
            maxAsteroidsVisible = directionsToAsteroids.size();
            imsPos = {h, w};
          }
        }
      }
    }
    std::cout << "Max asteroids visible: " << maxAsteroidsVisible << "\n";
    std::cout << "Asteroid with IMS installed: " << imsPos << "\n";

    // Part B: which asteroid will be vaporized as 200th one?
    auto vaporizeCounter = 0;
    auto laserAngle = 0.0; // degrees, that is: pointing a bit left from the NORTH.
    while ((vaporizeCounter < VAPORIZE_LIMIT) && (!map.empty()))
    {
      std::map<double, Position> anglesToAsteroid;
      for (std::size_t otherH = 0; otherH < map.size(); ++otherH)
      {
        for (std::size_t otherW = 0; otherW < map.size(); ++otherW)
        {
          // Don't let the IMS target itself:
          if ((otherH == imsPos.first) && (otherW == imsPos.second))
          {
            continue;
          }

          if (map[otherH][otherW] == ASTEROID)
          {
            Position v = normalize({otherH - imsPos.first, otherW - imsPos.second});
            // This asteroid may *not* be the first in sight from the station.
            // Finds the one that is actually targeted by the laser:
            Position target = {imsPos.first, imsPos.second};
            do
            {
              target.first += v.first;
              target.second += v.second;
            }
            while (map.at(target.first).at(target.second) != ASTEROID);

            // When inserting, Y axis needs to be flipped (so that it faces upwards):
            anglesToAsteroid.insert(std::make_pair(vectorToAngle({-v.first, v.second}), target));
          }
        }
      }
      auto it = anglesToAsteroid.upper_bound(laserAngle);
      if (it == anglesToAsteroid.end())
      {
        laserAngle = -1.0;
      }
      else
      {
        Position attackedCoords {it->second.first, it->second.second};
        laserAngle = it->first;
        map[attackedCoords.first][attackedCoords.second] = EMPTY_SPACE;
        ++vaporizeCounter;
        if (vaporizeCounter == VAPORIZE_LIMIT)
        {
          std::cout << "Asteroid vaporized as #" << vaporizeCounter << ": " << attackedCoords << "\n";
        }
      }

    }
  }
  else
  {
    std::cerr << "Could not open file " << INPUT_FILENAME << " for reading.\n";
  }

  return 0;
}
