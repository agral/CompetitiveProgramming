#!/usr/bin/env python3

"""
Solution to the task #01 of the "Advent of Code 2017" series.

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

data = open("input", "r").read().strip()
#data = input("Please provide a captcha-string: ")

captcha = 0
# Checks the first N-1 digits, where N is the length of a captcha string:
for i in range(len(data) - 1):
    if data[i] == data[i+1]:
        captcha += int(data[i])

# checks the last digit (the list is assumed to be circular):
if data[len(data) - 1] == data[0]:
    captcha += int(data[0])

second_captcha = 0
offset = len(data) // 2

for i in range(len(data)):
    adjacent = (i + offset) % len(data)
    if data[i] == data[adjacent]:
        second_captcha += int(data[i])

print(str(captcha))
print(str(second_captcha))
