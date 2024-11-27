#include <iostream>
#include <string>
#include <algorithm>
#include <cstdlib>

const int BUFSIZE = 256;

std::string first, second, result;
unsigned long long f, s, r;

void addrev();

int main(int argc, char *argv[]) {
  unsigned int n;
  std::cin >> n;

  for(unsigned int i = 0; i < n; i++) {
    addrev();
  }

  return 0;
}

void addrev() {
  std::cin >> first >> second;
  std::reverse(first.begin(), first.end());
  std::reverse(second.begin(), second.end());
  f = strtoull(first.c_str(), NULL, 10);
  s = strtoull(second.c_str(), NULL, 10);
  r = f + s;
  result = std::to_string(r);
  std::reverse(result.begin(), result.end());
  r = strtoull(result.c_str(), NULL, 10);
  std::cout << r << std::endl;
}
