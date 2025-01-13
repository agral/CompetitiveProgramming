#include <string>
#include <fstream>
#include <iostream>

using namespace std;

int partA(const std::string& line) {
  int first = -1, second = -1;
  for (int i = 0; i < line.size() && first < 0; i++) {
    if (line[i] >= '0' && line[i] <= '9') {
      first = line[i] - '0';
    }
  }
  for (int i = line.size() - 1; i >= 0 && second < 0; i--) {
    if (line[i] >= '0' && line[i] <= '9') {
      second = line[i] - '0';
    }
  }

  int ans = 10 * first + second;
  std::cout << ans << '\n';
  return ans;
}

int partB(const std::string& line) {
  std::cout << "Considering: " << line << "\n";
  std::string digits[] = {"unused", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"};

  int first = -1, second = -1;
  int posFirst = -1, posSecond = -1;
  for (posFirst = 0; posFirst < line.size(); posFirst++) {
    if (line[posFirst] >= '0' && line[posFirst] <= '9') {
      first = line[posFirst] - '0';
      break;
    }
  }
  for (posSecond = line.size() - 1; posSecond >= 0; posSecond--) {
    if (line[posSecond] >= '0' && line[posSecond] <= '9') {
      second = line[posSecond] - '0';
      break;
    }
  }
  // Take care of spelled out digits too.
  cout << "posFirst: " << posFirst << ", posSecond: " << posSecond << "\n";
  for (int i = 0; i <= 5; i++) { // repeat a few times.
    // Note: a real hack. Should have gone for std::string::rfind.
  for (int i = 1; i <= 9; i++) {
    size_t pos_start = line.find(digits[i]);
    size_t pos_end = line.find(digits[i], posSecond);
    if (pos_start < posFirst) {
      posFirst = pos_start;
      first = i;
      cout << "Starting: " << first << " found.\n";
    }
    if (pos_end > posSecond && pos_end < line.size()) {
      posSecond = pos_end;
      second = i;
      cout << "Ending: " << second << " found.\n";
    }
  }
  }

  int ans = 10 * first + second;
  std::cout << ans << '\n';
  return ans;
}

int main() {
  ifstream input_file("input.txt");
  ifstream example_file("example.txt");

  string line;
  int sum = 0;
  //while (getline(example_file, line)) {
  while (getline(input_file, line)) {
    sum += partB(line); // 54443 - too high
  }
  std::cout << "\n" << sum << "\n";
}
