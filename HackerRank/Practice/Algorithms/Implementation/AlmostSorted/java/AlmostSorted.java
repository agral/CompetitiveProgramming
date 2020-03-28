/*
 * Solution to the "Almost Sorted" challenge from HackerRank:
 * https://www.hackerrank.com/challenges/almost-sorted/problem
 * Created on: 25.03.2020
 * Last modified: 28.03.2020
 * Author: Adam Graliński (adam@gralin.ski)
 * License: MIT
 */

import java.io.*;
import java.util.*;
import java.util.stream.IntStream;

public class AlmostSorted
{
  public static boolean isSorted(int[] sequence) {
    return IntStream.range(0, sequence.length - 1).noneMatch(i -> sequence[i] > sequence[i + 1]);
  }
  public static void swap(int[] sequence, int lhs, int rhs) {
    int tmp = sequence[lhs];
    sequence[lhs] = sequence[rhs];
    sequence[rhs] = tmp;
  }
  public static void reverse(int[] sequence, int rangeStart, int rangeEnd) {
    while (rangeStart < rangeEnd) {
      swap(sequence, rangeStart, rangeEnd);
      rangeStart += 1;
      rangeEnd -= 1;
    }
  }
  public static void print(int[] sequence) {
    System.out.print("[");
    String separator = "";
    for (int i = 0; i < sequence.length; ++i) {
      System.out.print(separator + sequence[i]);
      separator = ", ";
    }
    System.out.println("]");
  }

  private static final Scanner scanner = new Scanner(System.in);

  public static void main(String[] args) {
    int N = scanner.nextInt();
    int[] sequence = new int[N];
    for (int n = 0; n < N; ++n) {
      sequence[n] = scanner.nextInt();
    }
    if (isSorted(sequence)) {
      System.out.println("yes");
    }
    else {
      int l = 0, r = N - 1;
      boolean foundL = false, foundR = false;

      // Finds out the index of first element of first strictly decreasing pair in the sequence (l):
      for (int i = 0; ((i < N - 2) && (!foundL)); ++i) {
        if (sequence[i] > sequence[i + 1]) {
          l = i;
          foundL = true;
        }
      }

      // Finds out the index of second element of last strictly decreasing pair in the sequence (r):
      for (int i = N - 1; ((i >= 1) && (!foundR)); --i) {
        if (sequence[i - 1] > sequence[i]) {
          r = i;
          foundR = true;
        }
      }

      // Checks whether the sequence with l, r swapped is sorted. If so, outputs "yes\nswap l r".
      // NOTE: this challenge requires outputting 1-indexed indices.
      swap(sequence, l, r);
      if (isSorted(sequence)) {
        System.out.format("yes\nswap %d %d\n", l + 1, r + 1);
      }
      else {
        // Swaps back, then checks whether the sequence with [l,r] range reversed is sorted.
        // If so, outputs "yes\nreverse l r".
        swap(sequence, l, r);
        reverse(sequence, l, r);
        if (isSorted(sequence)) {
          System.out.format("yes\nreverse %d %d\n", l + 1, r + 1);
        }
        else {
          System.out.println("no");
        }
      }
    }
  }
}
