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
print(available_coins)

def count_ways(target_value, available_coins, currently_considered=0):
    # If target == 0, then there is only one valid solution
    # (use 0 of each of the available coins):
    if target_value == 0:
        return 1

    # There is no way to obtain a negative sum, either:
    elif target_value < 0:
        return 0

    # If there are no more coins to use, but the target value is not zero,
    # there is no way to obtain the target value:
    elif currently_considered >= len(available_coins) and target_value != 0:
        return 0

    # The number of the ways is a sum of two simpler cases:
    # (1): solutions that do not use the first coin at all
    # (2): solutions that use the first coin at least once.
    return count_ways(target_value, available_coins, currently_considered+1) + \
            count_ways(target_value-available_coins[currently_considered],
                    available_coins, currently_considered)

print(count_ways(target_value, available_coins))
