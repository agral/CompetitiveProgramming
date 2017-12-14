#!/usr/bin/env python3

"""
Solution to the challenge #14 of the "Advent of Code 2017" series.
(Based on the solution to the challenge #10)

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

keys = ["flqrgnkx", "ffayrhll"]
DISK_SIZE = 128

def knothash(key):
    POSTFIX = [17, 31, 73, 47, 23]
    ROUNDS = 64
    LIST_LENGTH = 256
    sequence = list(range(LIST_LENGTH))
    bytekey = [ord(character) for character in key] + POSTFIX

    position = 0
    skip = 0
    for r in range(ROUNDS):
        for b in bytekey:
            start_index = position
            end_index = (position + b) % LIST_LENGTH
            if b == 0:
                afterwards = sequence
            else:
                if start_index < end_index:
                    prefix = sequence[:start_index]
                    sublist = sequence[start_index:end_index]
                    postfix = sequence[end_index:]
                    afterwards = prefix + list(reversed(sublist)) + postfix
                else: # The list wraps around
                    prefix = sequence[:end_index]
                    constant = sequence[end_index:start_index]
                    postfix = sequence[start_index:]
                    sublist = list(reversed(postfix + prefix))
                    prefix = sublist[len(postfix):]
                    postfix = sublist[:len(postfix)]
                    afterwards = prefix + constant + postfix

            position = (end_index + skip) % LIST_LENGTH
            skip += 1
            sequence = afterwards

    return sequence


def to_dense(sparse_hash):
    dense = [0] * 16
    for b in range(256):
        dense[b // 16] ^= sparse_hash[b]
    return dense


def to_hex_string(byte_list):
    return "".join(["{:02x}".format(byte) for byte in byte_list])


def to_bin_string(byte_list):
    return "".join(["{:08b}".format(byte) for byte in byte_list])


for key in keys:
    disk = []
    for i in range(DISK_SIZE):
        merged_key = "{:s}-{:d}".format(key,i)
        bin_hash = to_bin_string(to_dense(knothash(merged_key)))
        disk.append(bin_hash)

    # Challenge A: how many 1's are used?:
    ansA = sum([row.count('1') for row in disk])
    print("Part A: {} bits are set in total.".format(ansA))

    # Challenge B: how many side-connected regions are present?:
    diskmap = []
    diskmap.append(['0'] * (DISK_SIZE + 2))
    for row in disk:
        _row = []
        _row.append('0')
        for char in row:
            _row.append(char)
        _row.append('0')
        diskmap.append(_row)
    diskmap.append(['0'] * (DISK_SIZE + 2))

    def clear_region(x, y):
        diskmap[y][x] = '0'
        if diskmap[y][x-1] == '1':
            clear_region(x-1, y)
        if diskmap[y][x+1] == '1':
            clear_region(x+1, y)
        if diskmap[y-1][x] == '1':
            clear_region(x, y-1)
        if diskmap[y+1][x] == '1':
            clear_region(x, y+1)


    regions_count = 0
    for y in range(1, DISK_SIZE+1):
        for x in range(1, DISK_SIZE+1):
            if diskmap[y][x] == '1':
                clear_region(x, y)
                regions_count += 1
    print("Part B: {} distinct regions detected.".format(regions_count))

