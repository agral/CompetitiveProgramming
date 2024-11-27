#include <iostream>
#include <string>

int main(int argc, char *argv[]) {
  int P;
  std::cin >> P;
  for(int p = 0; p < P; p++) {
    int ttt = 0, tth = 0, tht = 0, thh = 0,
        htt = 0, hth = 0, hht = 0, hhh = 0;
    int N;
    std::string s;
    std::cin >> N >> s;
    for(int ch = 2; ch < 40; ch++) {
      if((s[ch - 2] == 'T') && (s[ch - 1] == 'T') && (s[ch] == 'T')) {
        ttt++;
      }
      if((s[ch - 2] == 'T') && (s[ch - 1] == 'T') && (s[ch] == 'H')) {
        tth++;
      }
      if((s[ch - 2] == 'T') && (s[ch - 1] == 'H') && (s[ch] == 'T')) {
        tht++;
      }
      if((s[ch - 2] == 'T') && (s[ch - 1] == 'H') && (s[ch] == 'H')) {
        thh++;
      }
      if((s[ch - 2] == 'H') && (s[ch - 1] == 'T') && (s[ch] == 'T')) {
        htt++;
      }
      if((s[ch - 2] == 'H') && (s[ch - 1] == 'T') && (s[ch] == 'H')) {
        hth++;
      }
      if((s[ch - 2] == 'H') && (s[ch - 1] == 'H') && (s[ch] == 'T')) {
        hht++;
      }
      if((s[ch - 2] == 'H') && (s[ch - 1] == 'H') && (s[ch] == 'H')) {
        hhh++;
      }
    }
    std::cout << N << " " << ttt << " " << tth << " " << tht << " " << thh
      << " " << htt << " " << hth << " " << hht << " " << hhh << std::endl;
  }
}
