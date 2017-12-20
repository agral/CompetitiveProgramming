#!/usr/bin/env python3

"""
Solution to the challenge #05 of the "Advent of Code 2017" series.
http://adventofcode.com/2017/day/5
Written on 05-12-2017 in Python3
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

#data_file = "testcase0"
data_file = "input_05"

jumps = [int(line.strip()) for line in open(data_file, "r")]

IP = 0 # Instruction Pointer
steps = 0 # Steps counter
while IP >= 0 and IP < len(jumps):
    newIP = IP + jumps[IP]
    jumps[IP] += 1
    IP = newIP
    steps += 1

print("Part A completed in {} steps.".format(steps))


# Challenge 02: steps DECREASE by 1 if jump offset >= 3
jumps = [int(line.strip()) for line in open(data_file, "r")]

IP = 0 # Instruction Pointer
steps = 0 # Steps counter
while IP >= 0 and IP < len(jumps):
    newIP = IP + jumps[IP]
    if jumps[IP] >= 3:
        jumps[IP] -= 1
    else:
        jumps[IP] += 1
    IP = newIP
    steps += 1

print("Part B completed in {} steps.".format(steps))
