/*
 * Solution to the "Mini-Max Sum" challenge from HackerRank:
 * https://www.hackerrank.com/challenges/mini-max-sum/problem
 * Created on: 23.03.2020
 * Last modified: 23.03.2020
 * Author: Adam Graliński (adam@gralin.ski)
 * License: MIT
 */

import java.io.*;
import java.util.*;

public class MinimaxSum
{
  public static final int SIZE = 5;
  private static final Scanner scanner = new Scanner(System.in);

  public static void main(String[] args) {
    long input[] = new long[SIZE];
    for (int i = 0; i < SIZE; ++i) {
      input[i] = scanner.nextInt();
    }
    Arrays.sort(input);
    long min = 0, max = 0;
    for (int i = 0; i < SIZE - 1; ++i) {
      min += input[i];
      max += input[i + 1];
    }

    System.out.format("%d %d\n", min, max);
  }
}
