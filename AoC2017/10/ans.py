#!/usr/bin/env python3

"""
Solution to the challenge #10 of the "Advent of Code 2017" series.

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

def fancyprint(sequence, current_position=0, start=None, end=None):
    text = ""
    for i in range(len(sequence)):
        if i == start:
            text += "("
        if i == current_position:
            text += "["
        text += str(sequence[i])
        if i == current_position:
            text += "]"
        if i == end:
            text += ")"
        text += " "
    print(text)


data_file = "input_10"
list_length = 256

#data_file = "sample_input_10"
#list_length = 5

#data_file = "aoc_input"
#list_length = 256

# "list" is a reserved keyword. "sequence", however, is not.
sequence = [num for num in range(list_length)]
print("Sequence is: ", end='')
fancyprint(sequence)

data = [line.strip() for line in open(data_file, "r")]
#lengths = [int(datum) for datum in data[0].split(",")] # valid for challenge A
lengths = [ord(char) for char in data[0]] # valid for challenge B

POSTFIX_LENGTHS = [17, 31, 73, 47, 23]
lengths += POSTFIX_LENGTHS
print("LENGTHS:")
print(lengths)

current_position = 0
current_skip = 0

no_of_rounds = 1 # valid for challenge A
no_of_rounds = 64 # valid for challenge B
for r in range(no_of_rounds):
    print("Round #{} begins.".format(r+1))
    for length in lengths:
        print("Step #{}, length={}, current_position={}".format(
                current_skip, length, current_position))
        start_index = current_position
        end_index = (current_position + length) % list_length
        #fancyprint(sequence, current_position, start_index, end_index-1)

        if length == 0:
            afterwards = sequence
            print("Zero-length datum processed.")
        else:
            if start_index < end_index:
                prefix = sequence[:start_index]
                sublist = sequence[start_index:end_index]
                postfix = sequence[end_index:]
                sublist = list(reversed(sublist))
                afterwards = prefix + sublist + postfix
            elif start_index >= end_index:
                prefix = sequence[:end_index]
                constant = sequence[end_index:start_index]
                postfix = sequence[start_index:]
                sublist = postfix + prefix
                sublist = list(reversed(postfix+prefix))

                postfix = sublist[:len(postfix)]
                prefix = sublist[len(postfix):]
                afterwards = prefix + constant + postfix

        current_position = (end_index + current_skip) % list_length
        current_skip += 1
        sequence = afterwards

    print("After round #{} : ".format(r+1), end='')
    fancyprint(sequence, current_position)
    print()

# Challenge A:
checksum = sequence[0] * sequence[1]
print("-------------\nChecksum: {}\n-------------".format(checksum))

# Challenge B:
print("Sparse hash:")
print(sequence)

dense_hash = [0] * 16
for part in range(256):
    dense_hash[part // 16] ^= sequence[part]

dense_str = ["{:02x}".format(chunk) for chunk in dense_hash]
print("Dense hash: ", end='')
print("".join(dense_str))
