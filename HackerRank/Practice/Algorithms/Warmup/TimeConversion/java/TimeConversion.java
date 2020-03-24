/*
 * Solution to the "Time Conversion" challenge from HackerRank:
 * https://www.hackerrank.com/challenges/time-conversion/problem
 * Created on: 23.03.2020
 * Last modified: 24.03.2020
 * Author: Adam Graliński (adam@gralin.ski)
 * License: MIT
 */

import java.io.*;
import java.util.*;

public class TimeConversion
{
  private static final Scanner scanner = new Scanner(System.in);

  public static void main(String[] args) {
    String twelve_hour_time = scanner.nextLine();

    // Twelve-hour time example: 07:05:45PM
    int hour = Integer.parseInt(twelve_hour_time.substring(0, 2)) % 12;
    if (twelve_hour_time.substring(8, 10).equals("PM")) {
      hour += 12;
    }

    System.out.format("%02d%s\n", hour, twelve_hour_time.substring(2, 8));
  }
}
