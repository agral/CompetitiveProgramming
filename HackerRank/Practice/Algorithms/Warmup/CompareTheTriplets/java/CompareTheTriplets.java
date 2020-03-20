/*
 * Solution to the "Compare the Triplets" challenge from HackerRank:
 * https://www.hackerrank.com/challenges/compare-the-triplets/problem
 * Created on: 20.03.2020
 * Last modified: 20.03.2020
 * Author: Adam Graliński (adam@gralin.ski)
 * License: MIT
 */

import java.io.*;
import java.util.*;

public class CompareTheTriplets
{
  private static final Scanner scanner = new Scanner(System.in);

  public static final int SIZE = 3;

  public static void main(String[] args) {
    int[] alice = new int[SIZE];
    int scoreAlice = 0, scoreBob = 0;
    for (int n = 0; n < SIZE; ++n) {
      alice[n] = scanner.nextInt();
    }
    for (int n = 0; n < SIZE; ++n) {
      int bob = scanner.nextInt();
      if (alice[n] > bob) {
        scoreAlice += 1;
      }
      else if (alice[n] < bob) {
        scoreBob += 1;
      }
    }

    System.out.println(scoreAlice + " " + scoreBob);
  }
}
