/*
 * Solution to the "Staircase" challenge from HackerRank:
 * https://www.hackerrank.com/challenges/staircase/problem
 * Created on: 23.03.2020
 * Last modified: 23.03.2020
 * Author: Adam Graliński (adam@gralin.ski)
 * License: MIT
 */

import java.io.*;
import java.util.*;

public class Staircase
{
  private static final Scanner scanner = new Scanner(System.in);

  public static void main(String[] args) {
    int height = scanner.nextInt();
    for (int h = height - 1; h >= 0; --h) {
      for (int s = 0; s < h; ++s) {
        System.out.print(" ");
      }
      for (int b = 0; b < height - h; ++b) {
        System.out.print("#");
      }
      System.out.println();
    }
  }
}
