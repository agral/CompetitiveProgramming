#include <iostream>
#include <string>

void printNextPalindrome(std::string number);

int main(int argc, char *argv[]) {
  unsigned int t;
  std::cin >> t;
  std::string dummy;
  std::getline(std::cin, dummy);
  for(unsigned int testCase = 0; testCase < t; testCase++) {
    std::string nextNumber;
    std::getline(std::cin, nextNumber);
    printNextPalindrome(nextNumber);
  }
}

void printNextPalindrome(std::string number) {
  int currentDigit;
  const int BUFFER_SIZE = 1500000;
  char answer[BUFFER_SIZE];
  int length = number.length();
  int halfLength = length / 2;
  bool isLengthEven = (length % 2 == 0);

  // a special case: for numbers consisting of solely 9's:
  bool isNinesOnly = true;
  for(int i = 0; i < length; i++) {
    if(number[i] != '9') {
      isNinesOnly = false;
      break;
    }
  }
  if(isNinesOnly) {
    answer[0] = '1';
    for(int d = 1; d < length; d++) {
      answer[d] = '0';
    }
    answer[length] = '1';
    answer[length + 1] = '\0';
  }
  else {
    int isFirstHalfGreater = 0; // 0: equal, -1: no, 1: yes.
    currentDigit = 0;
    int currentDigitB = halfLength;
    if(!isLengthEven) {
      currentDigitB++;
    }
    while(
        (number[currentDigit] == number[currentDigitB]) &&
        (currentDigit < halfLength)
    ) {
      currentDigit++;
      currentDigitB++;
    }
    if(number[currentDigit] > number[currentDigitB]) {
      isFirstHalfGreater = 1;
    } else if(number[currentDigit] < number[currentDigitB]) {
      isFirstHalfGreater = -1;
    }

    number.copy(answer, halfLength, 0);
    currentDigit = halfLength - 1;

    if(!isLengthEven) {
      answer[halfLength] = number[halfLength];
      currentDigit = halfLength;
    }
    
    if(isFirstHalfGreater == 1) {
      // just re-write first half to form a palindrome:
      for(currentDigit = 0; currentDigit < halfLength; currentDigit++) {
        answer[length - currentDigit - 1] = answer[currentDigit];
      }
    } else {
      answer[currentDigit] = answer[currentDigit] + 1;
      while((answer[currentDigit] > '9') && (currentDigit > 0)) {
        answer[currentDigit] = '0';
        currentDigit--;
        answer[currentDigit] = answer[currentDigit] + 1;
      }
      // now re-write first half to form a palindrome:
      for(currentDigit = 0; currentDigit < halfLength; currentDigit++) {
        answer[length - currentDigit - 1] = answer[currentDigit];
      }
    }

    answer[length] = '\0';
  }

  std::string output(answer);
  //std::cout << answer << std::endl;
  std::cout << output << std::endl;
}

