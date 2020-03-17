/*
 * Solution to the "AVeryBigSum" challenge from HackerRank:
 * https://www.hackerrank.com/challenges/a-very-big-sum/problem
 * Created on: 17.03.2020
 * Last modified: 17.03.2020
 * Author: Adam Gralinski (adam@gralin.ski)
 * License: MIT
 */

import java.io.*;
import java.util.*;

public class AVeryBigSum
{
  static long bigSum(long[] array) {
    long bigsum = 0;
    for (int i = 0; i < array.length; ++i) {
      bigsum += array[i];
    }

    return bigsum;
  }

  private static final Scanner scanner = new Scanner(System.in);

  public static void main(String[] args) {
    int arraySize = Integer.parseInt(scanner.nextLine().trim());

    String[] strArray = scanner.nextLine().split(" ");
    if (strArray.length != arraySize) {
      System.err.println("Error: Declared and actual array sizes mismatch. Aborting.");
    }
    else {
      long[] longArray = new long[arraySize];
      for (int i = 0; i < arraySize; ++i) {
        longArray[i] = Integer.parseInt(strArray[i].trim());
      }

      System.out.println(String.valueOf(bigSum(longArray)));
    }
  }
}
