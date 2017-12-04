#!/usr/bin/env python3

"""
Solution to a "Marc's Cakewalk" challenge by HackerRank:
https://www.hackerrank.com/challenges/marcs-cakewalk/problem
Written on 04-12-2017 in Python3
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


def main():
    cupcakes = int(input())
    calories = [int(num) for num in input().split()]
    ordered = sorted(calories, reverse=True)

    miles = 0
    penalty = 1 # == 2 ** 0
    for i in range(len(ordered)):
        miles += penalty * ordered[i]
        penalty *= 2

    print(miles)


if __name__ == "__main__":
    main()
