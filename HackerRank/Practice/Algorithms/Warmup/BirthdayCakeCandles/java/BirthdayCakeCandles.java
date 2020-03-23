/*
 * Solution to the "Birthday Cake Candles" challenge from HackerRank:
 * https://www.hackerrank.com/challenges/birthday-cake-candles/problem
 * Created on: 23.03.2020
 * Last modified: 23.03.2020
 * Author: Adam Graliński (adam@gralin.ski)
 * License: MIT
 */

import java.io.*;
import java.util.*;

public class BirthdayCakeCandles
{
  private static final Scanner scanner = new Scanner(System.in);

  public static void main(String[] args) {
    int N = scanner.nextInt();
    int tallestHeight = -1;
    int count = 0;
    for (int n = 0; n < N; ++n) {
      int candleHeight = scanner.nextInt();
      if (candleHeight == tallestHeight) {
        count += 1;
      }
      else if (candleHeight > tallestHeight) {
        tallestHeight = candleHeight;
        count = 1;
      }
    }
    System.out.println(count);
  }
}
