/*
 * Solution to the "Diagonal Difference" challenge from HackerRank:
 * https://www.hackerrank.com/challenges/diagonal-difference/problem
 * Created on: 21.03.2020
 * Last modified: 21.03.2020
 * Author: Adam Graliński (adam@gralin.ski)
 * License: MIT
 */

import java.io.*;
import java.util.*;

public class DiagonalDifference
{
  private static final Scanner scanner = new Scanner(System.in);

  public static void main(String[] args) {
    int N = scanner.nextInt();
    int[][] array = new int[N][N];
    for (int row = 0; row < N; ++row) {
      for (int col = 0; col < N; ++col) {
        array[row][col] = scanner.nextInt();
      }
    }

    int diag1 = 0, diag2 = 0;
    for (int k = 0; k < N; ++k) {
      diag1 += array[k][k];
      diag2 += array[k][N - k - 1];
    }

    System.out.println(Math.abs(diag1 - diag2));
  }
}
