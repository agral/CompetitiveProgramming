/**
 * Name: aoc2019_day01_ans.cpp
 * Description: Answers the AoC2019 day#01 challenge
 * Created on: 02.12.2019
 * Last modified: 02.12.2019
 * Author: Adam Graliński (adam@gralin.ski)
 * License: CC0
 */

#include <algorithm>
#include <fstream>
#include <iostream>
#include <numeric>
#include <vector>

long long GetFuelRequired(long long mass)
{
  return (mass / 3LL) - 2LL;
}

int main()
{
  std::vector<long long> moduleMasses;
  std::ifstream myfile;
  myfile.open("aoc2019_day01_input.txt");
  if (myfile.is_open())
  {
    long long moduleMass;
    while (true)
    {
      myfile >> moduleMass;
      if (myfile.eof())
      {
        break;
      }
      moduleMasses.push_back(moduleMass);
    }
    myfile.close();
  }
  else
  {
    std::cerr << "Could not open file.\n";
  }

  long long totalMass = std::accumulate(moduleMasses.begin(), moduleMasses.end(), 0LL);
  std::cout << "modules: " << moduleMasses.size() << "\n";
  std::cout << "total mass: " << totalMass << "\n";

  std::vector<long long> fuelMasses;
  fuelMasses.resize(moduleMasses.size());

  for (std::size_t index = 0; index < moduleMasses.size(); ++index)
  {
    fuelMasses[index] = GetFuelRequired(moduleMasses[index]);
  }

  long long initialFuel = std::accumulate(fuelMasses.begin(), fuelMasses.end(), 0LL);
  std::cout << "Initial fuel requirement: " << initialFuel << "\n";

  int iteration = 1;
  long long totalFuel = initialFuel;

  while (std::any_of(fuelMasses.begin(), fuelMasses.end(), [](long long fm){return fm > 0LL;}))
  {
    for (auto& fm : fuelMasses)
    {
      if (fm > 0LL)
      {
        fm = std::max(GetFuelRequired(fm), 0LL);
        totalFuel += fm;
      }
    }
    iteration += 1;
    std::cout << "Iteration: " << iteration << ", total fuel: " << totalFuel << "\n";
  }

  return 0;
}
