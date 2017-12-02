#!/usr/bin/env python3

"""
Solution to the task #02 of the "Advent of Code 2017" series.

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

#data_file = "sample_input01"
#separator = " "

data_file = "input"
separator = "\t"

contents = [line.strip() for line in open(data_file)]
rows = [line.split(separator) for line in contents]

idata = []
for row in rows:
    irow = [int(datum) for datum in row]
    idata.append(irow)

sdata = [sorted(row) for row in idata]

# Challenge 01: calculate the checksum:
delta_checksums = [row[-1] - row[0] for row in sdata]
checksum = sum(delta_checksums)
print("---------------\nChecksum: {:d}\n---------------\n".format(checksum))

# Challenge 02: find sum of even results of division:
divsum = 0
for row in sdata:
    for d in range(0, len(sdata)):
        for D in range(d+1, len(sdata)):
            if row[D] % row[d] == 0:
                res = row[D] // row[d]
                print("{} / {} == {}".format(row[D], row[d], res))
                divsum += res

print("---------------\nDiv-sum: {:d}\n---------------".format(divsum))
