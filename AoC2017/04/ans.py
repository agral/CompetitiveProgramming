#!/usr/bin/env python3

"""
Solution to the task #04 of the "Advent of Code 2017" series.

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

data = open("input_04", "r").readlines()
data = [line.strip() for line in data]

valid = 0
invalid = 0

for line in data:
    words = line.split()
    wset = set(words)
    if len(words) == len(wset):
        valid += 1
    else:
        invalid += 1

print("Valid: {}, invalid: {}, total: {}".format(valid, invalid, valid+invalid))

# challenge #02:
valid = 0
invalid = 0
for line in data:
    words = line.split()
    # Creates a list that has the words with same letters,
    # but the letters are sorted alphabetically (e.g. "cat" --> "act")
    sorted_words = ["".join(sorted(word)) for word in words]
    swset = set(sorted_words)
    if len(sorted_words) == len(swset):
        valid += 1
    else:
        invalid += 1

print("Under the new policy,")
print("Valid: {}, invalid: {}, total: {}".format(valid, invalid, valid+invalid))
