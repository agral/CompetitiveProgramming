/**
 * Solution to the "Time Conversion" challenge from HackerRank:
 * https://www.hackerrank.com/challenges/time-conversion/problem
 * Created on: 23.03.2020
 * Last modified: 24.03.2020
 * Author: Adam Graliński (adam@gralin.ski)
 * License: MIT
 */

#include <iomanip>
#include <iostream>
#include <string>

int main(int, char**)
{
  std::string twelve_hour_time;
  std::cin >> twelve_hour_time;

  // Twelve-hour time example: 07:05:45PM
  auto hour = (std::stoi(twelve_hour_time.substr(0, 2)) % 12);
  if (twelve_hour_time.substr(8, 2) == "PM") {
    hour += 12;
  }
  std::cout << std::setfill('0') << std::setw(2) << hour << twelve_hour_time.substr(2, 6) << "\n";

  return 0;
}
