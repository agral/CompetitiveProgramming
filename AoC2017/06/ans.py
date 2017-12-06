#!/usr/bin/env python3

"""
Solution to the task #06 of the "Advent of Code 2017" series.

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

data_file = "input_06"
#data_file = "sample_input_06"
separator = "\t"

contents = [line.strip() for line in open(data_file)]
banks = [int(entry) for entry in contents[0].split(separator)]
past_distributions = set()
counter = 0

def rebalance():
    global counter
    counter += 1

    # Finds the index of a first bank with most load:
    most_loaded = 0
    for bank in range(1, len(banks)):
        if banks[bank] > banks[most_loaded]:
            most_loaded = bank

    # Performs the rebalancing:
    storage = banks[most_loaded]
    banks[most_loaded] = 0
    ptr = most_loaded
    while storage > 0:
        ptr = (ptr + 1) % len(banks)
        banks[ptr] += 1
        storage -= 1


# Emulates a do-while loop in python:
while True:
    rebalance()
    current_state_serialized = ",".join(str(entry) for entry in banks)
    if current_state_serialized in past_distributions:
        print("First repetition found in {} steps".format(counter))
        break
    else:
        past_distributions.add(current_state_serialized)
        print("After rebalance #{}: {}".format(counter, current_state_serialized))

# Challenge #02:
root_state = ",".join(str(entry) for entry in banks)
counter = 0
while True:
    rebalance()
    current_state_serialized = ",".join(str(entry) for entry in banks)
    if current_state_serialized == root_state:
        print("The cycle has length of {}".format(counter))
        break
