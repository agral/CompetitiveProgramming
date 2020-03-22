/*
 * Solution to the "Plus Minus" challenge from HackerRank:
 * https://www.hackerrank.com/challenges/plus-minus/problem
 * Created on: 22.03.2020
 * Last modified: 22.03.2020
 * Author: Adam Graliński (adam@gralin.ski)
 * License: MIT
 */

import java.io.*;
import java.util.*;

public class PlusMinus
{
  private static final Scanner scanner = new Scanner(System.in);

  public static void main(String[] args) {
    int N = scanner.nextInt();
    int negative = 0, zeroes = 0, positive = 0;
    for (int n = 0; n < N; ++n) {
      int number = scanner.nextInt();
      if (number > 0) {
        positive += 1;
      }
      else if (number == 0) {
        zeroes += 1;
      }
      else {
        negative += 1;
      }
    }

    System.out.format( "%.6f\n%.6f\n%.6f\n", Double.valueOf(positive) / N,
        Double.valueOf(negative) / N, Double.valueOf(zeroes) / N);
  }
}
