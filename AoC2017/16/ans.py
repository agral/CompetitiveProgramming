#!/usr/bin/env python3

"""
Solution to the challenge #16 of the "Advent of Code 2017" series.

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

#START = "abcde"
#input_file = "sample_input_16"
START = "abcdefghijklmnop"
input_file = "input_16"

order = list(START)

# Dance figures:
def spin(size):
    global order
    divider = len(order) - size
    order = order[divider:] + order[:divider]

def exchange(pos1, pos2):
    global order
    order[pos1], order[pos2] = order[pos2], order[pos1]

def partner(name1, name2):
    global order
    pos1 = order.index(name1)
    pos2 = order.index(name2)
    order[pos1], order[pos2] = order[pos2], order[pos1]

input_data = [line.strip() for line in open(input_file, "r")]
moves = input_data[0].split(",")

# Challenge A:
for move in moves:
    category = move[0]
    data = move[1:]
    if category == "s":
        spin(int(data))
    elif category == "x":
        indices = data.split("/")
        exchange(int(indices[0]), int(indices[1]))
    elif category == "p":
        partners = data.split("/")
        partner(partners[0], partners[1])

print("[Challenge A] Final order: ", end="")
print("".join(order))

# Challenge B: Check whether there is a cycle (position repeated after N iters):
START = "abcdefghijklmnop"
TOTAL_ITERATIONS = 1000 * 1000 * 1000
order = list(START)

previous_orders = set()
consecutive_iterations = []
cycle_length = None
for i in range(TOTAL_ITERATIONS):
    for move in moves:
        category = move[0]
        data = move[1:]
        if category == "s":
            spin(int(data))
        elif category == "x":
            indices = data.split("/")
            exchange(int(indices[0]), int(indices[1]))
        elif category == "p":
            partners = data.split("/")
            partner(partners[0], partners[1])
    strorder = "".join(order) # a list is unhashable, but a string is.
    if strorder in previous_orders:
        print("Repetition found after {} iterations.".format(i))
        cycle_length = i
        break
    else:
        previous_orders.add(strorder)
        consecutive_iterations.append(strorder)

rem = TOTAL_ITERATIONS % i
print("[Challenge B] Final order: ", end="")

# Iterations go from 1 to 1B, but python's arrays are 0-indexed:
print(consecutive_iterations[rem-1])

