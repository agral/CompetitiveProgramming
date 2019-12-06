/**
 * Name: aoc2019_day06_ans
 * Description: Answers the AoC2019 day#06 challenge
 * Created on: 06.12.2019
 * Last modified: 06.12.2019
 * Author: Adam Gralinski (adam@gralin.ski)
 * License: CC0
 */

#include <algorithm>
#include <fstream>
#include <iostream>
#include <sstream>
#include <string>
#include <vector>
#include <map>
#include <set>

//const std::string inputFileName = "sample_input.txt";
//const std::string inputFileName = "sample_input_partB.txt";
const std::string inputFileName = "aoc2019_day06_input.txt";
const char CHAR_ORBIT_DELIMITER = ')';
const auto CENTER_OF_MASS_NAME = std::string("COM");
const auto YOU_NAME = std::string("YOU");
const auto SANTA_NAME = std::string("SAN");
std::map<std::string, std::string> satellitesToBodies;

int main(int, char**)
{
  std::ifstream inputHandle;
  inputHandle.open(inputFileName);
  if (inputHandle.is_open())
  {
    std::string line;
    while (std::getline(inputHandle, line))
    {
      std::stringstream ss(line);
      std::string body, satellite;
      std::getline(ss, body, CHAR_ORBIT_DELIMITER);
      std::getline(ss, satellite);
      std::cout << "[" << satellite << "] is orbiting [" << body << "].\n";
      satellitesToBodies[satellite] = body;
    }

    auto totalOrbitsCount = 0LL;
    for (auto [satellite, body] : satellitesToBodies)
    {
      auto distanceToCenter = 1;
      while (body != CENTER_OF_MASS_NAME)
      {
        body = satellitesToBodies[body];
        distanceToCenter += 1;
      }
      totalOrbitsCount += distanceToCenter;
      std::cout << "Satellite: " << satellite << " is " << distanceToCenter << " orbits away from center.\n";
    }
    std::cout << "Total orbits count: " << totalOrbitsCount << "\n";

    // Part B: how many orbital transfers between YOU and SAN?
    std::vector<std::string> youToCom;
    auto body = satellitesToBodies[YOU_NAME];
    while (body != CENTER_OF_MASS_NAME)
    {
      youToCom.push_back(body);
      body = satellitesToBodies[body];
    }
    auto santaDistanceToCommonBody = 0;
    body = satellitesToBodies[SANTA_NAME];
    while (std::find(youToCom.begin(), youToCom.end(), body) == youToCom.end())
    {
      body = satellitesToBodies[body];
      santaDistanceToCommonBody += 1;
    }
    auto it = std::find(youToCom.begin(), youToCom.end(), body);
    auto youDistanceToCommonBody = it - youToCom.begin();
    auto totalOrbitsChange = youDistanceToCommonBody + santaDistanceToCommonBody;
    std::cout << "You need to change " << totalOrbitsChange << " orbits to get to Santa.\n";
  }
  else
  {
    std::cerr << "Could not open file " << inputFileName << " for reading.\n";
    return 1;
  }
  return 0;
}
