#!/usr/bin/env python3

"""
Solution to the "Coin Change" challenge by HackerRank:
https://www.hackerrank.com/challenges/coin-change/problem
Written on 27-12-2017 in Python3
by simba (szczerbiakadam@gmail.com).

MIT License

Copyright (c) 2017 Adam Szczerbiak (szczerbiakadam@gmail.com)

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
"""

target_value, coin_count = [int(k) for k in input().split(" ")]
available_coins = [int(v) for v in input().split(" ")]

def count_ways(target_value, coins):
    dp_table = [[0 for column in range(len(coins))] \
            for row in range(target_value+1)]

    # Fills the entries for the base case [target_value == 0]:
    for column in range(len(coins)):
        dp_table[0][column] = 1

    # Fills the rest of the table:
    for row in range(1, target_value+1):
        for c in range(len(coins)):
            # Counts the number of distinct solutions including c'th coin:
            inclusive = dp_table[row-coins[c]][c] if row-coins[c] >= 0 else 0

            # Counts the number of distinct solutions excluding c'th coin:
            exclusive = dp_table[row][c-1] if c > 0 else 0

            dp_table[row][c] = inclusive + exclusive

    """
    for row in range(target_value+1):
        for c in range(len(coins)):
            print("{}\t".format(dp_table[row][c]), end="")
        print()

    """
    return dp_table[target_value][len(coins)-1]

print(count_ways(target_value, available_coins))
