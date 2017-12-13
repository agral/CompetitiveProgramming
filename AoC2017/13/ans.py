#!/usr/bin/env python3

"""
Solution to the challenge #13 of the "Advent of Code 2017" series.

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
import copy


data_files = ["sample_input_13", "input_13" ]

# Prints out the firewall!
def visualize(firewall, firewall_length, player_x):
    for line in range(firewall_length+1):
        if line in firewall:
            level = firewall[line]
            for field in range(level["length"]):
                field_ch1, field_ch2, interior = '[', ']', ' '
                if field == 0 and player_x == line:
                    field_ch1, field_ch2 = '(', ')'
                if field == level["position"]:
                    interior = '>' if level["direction"] == 1 else '<'
                print("{}{}{} ".format(field_ch1, interior, field_ch2), end="")
            print()
        else:
            empty_level = "(X)" if line == player_x else "..."
            print(empty_level)
    print("----------")


def reset_firewall(firewall):
    for level in firewall:
        firewall[level]["position"] = 0
        firewall[level]["direction"] = 1


def scanners_step(firewall):
    for level in firewall:
        firewall[level]["position"] += firewall[level]["direction"]
        if firewall[level]["position"] == 0:
            firewall[level]["direction"] = 1
        elif firewall[level]["position"] == firewall[level]["length"] - 1:
            firewall[level]["direction"] = -1


def calculate_severity(firewall, firewall_length):
    time = 0 # picoseconds
    reset_firewall(firewall)
    severity = 0

    x_position = 0 # X just entered the firewall! :-)
    while x_position <= firewall_length:
        ###visualize(firewall, firewall_length, x_position)
        # Does X stay exactly under the scanner's beam?
        if x_position in firewall:
            if firewall[x_position]["position"] == 0:
                severity += firewall[x_position]["length"] * x_position
        # Moves X, then moves all the scanners in their appropriate directions:
        x_position += 1
        scanners_step(firewall)
    return severity

def is_caught(firewall, firewall_length, delay_time):
    time = 0 # picoseconds
    reset_firewall(firewall)
    for t in range(delay_time):
        scanners_step(firewall)
        time += 1

    x_position = 0 # X just entered the firewall! :-)
    while x_position <= firewall_length:
        ###visualize(firewall, firewall_length, x_position)
        # Does X stay exactly under the scanner's beam?
        if x_position in firewall:
            if firewall[x_position]["position"] == 0:
                return True
        # Moves X, then moves all the scanners in their appropriate directions:
        x_position += 1
        scanners_step(firewall)
    return False

def is_caught2(firewall, firewall_length, delay_time):
    x_position = 0
    while x_position <= firewall_length:
        if x_position in firewall:
            if firewall[x_position]["position"] == 0:
                return True
        x_position += 1
        scanners_step(firewall)
    return False

for data_file in data_files:
    print("Processing file: {} ...".format(data_file))
    lines = [line.strip() for line in open(data_file, "r")]

    firewall = dict()
    firewall_length = 0

    for line in lines:
        data = line.split(": ")
        row = int(data[0])
        length = int(data[1])
        firewall[row] = {
                "length": length,
                "position": 0,
                "direction": 1
        }
        if row > firewall_length:
            firewall_length = row

    # Challenge A: if X were to enter immediately, what is a total severity?
    severity = calculate_severity(firewall, firewall_length)
    print("Challenge A: Total severity: {}.".format(severity))

    # Challenge B: how many picoseconds to wait in order not to get caught?
    reset_firewall(firewall)
    delay = 0
    for i in range(delay):
        scanners_step(firewall)
    f = copy.deepcopy(firewall)
    while is_caught2(f, firewall_length, delay):
        delay += 1
        scanners_step(firewall)
        f = copy.deepcopy(firewall)
        if delay % 10000 == 0:
            print("is_caught: checking state after {} picoseconds...".format(
                    delay))
    print("X should go after {} seconds.".format(delay))
