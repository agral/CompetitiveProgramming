#include <iostream>
#include <bitset>

int in[] = {137, 1315, 73, 136, 255, 1384, 16385};

void printNumberEncrypted(int n) {  
  if(n == 1) {
    std::cout << "2(0)";
  }
  else if(n == 2) {
    std::cout << "2";
  }
  else {
    if(n % 2 == 1) {
      std::cout << "2(";
      printNumberEncrypted(n/2);
      std::cout << ")";
    }
    else {
      printNumberEncrypted(n/2);
    }
  }
}

void printNumber(int n) {
  int mask = 1 << 16;
  bool firstExpression = true;
  while(mask > 0) {
    if(n & mask) {
      if(!firstExpression) {

      }
    }
    mask >>= 1;
  }
}

int main(int argc, char* argv[]) {
  for(int i = 0; i < 7; i++) {
    printNumberEncrypted(in[i]);
  }
  return 0;
}
