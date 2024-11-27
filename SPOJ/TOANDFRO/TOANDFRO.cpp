#include <iostream>
#include <string>

char board[20][100];
int columns, rows;

void printBoard();

int main(int argc, char* argv[]) {
  int nextChar;
  std::string ct;

  std::cin >> columns;
  while(columns > 0) {
    std::cin >> ct;
    rows = ct.length() / columns;

    // std::cout << "Matrix: " << columns << "x" << rows << "." << std::endl;
    
    nextChar = 0;
    for(int r = 0; r < rows; r++) {
      for(int c = 0; c < columns; c++) {
        if(r % 2 == 0) {
          board[c][r] = ct[nextChar];
        }
        else {
          board[columns - c - 1][r] = ct[nextChar];
        }
        nextChar++;
      }
    }
    
    // printBoard();

    for(int c = 0; c < columns; c++) {
      for(int r = 0; r < rows; r++) {
        std::cout << board[c][r];
      }
    }

    std::cout << std::endl;
    std::cin >> columns;
  }
}

void printBoard() {
  for(int r = 0; r < rows; r++) {
    for(int c = 0; c < columns; c++) {
      std::cout << board[c][r];
      // std::cout << "(" << r << "," << c << ")\t";
    }
    std::cout << std::endl;
  }
}
