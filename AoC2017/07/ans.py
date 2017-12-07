#!/usr/bin/env python3

"""
Solution to the task #07 of the "Advent of Code 2017" series.

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

class Program:
    def __init__(self, name, number, leaves):
        self.name = name
        self.number = number
        self.leaves = leaves
        self.parent = None
        self.totalWeight = None

data_file = "input_07"
#data_file = "sample_input_07"
lines = [line.strip() for line in open(data_file, "r")]

""" Each line consists of:
    NAME (number) [ -> list-of-leaves ] """
# Reads the data:
data = []
for line in lines:
    words = line.split(" ")
    name = words[0]
    number = int(words[1][1:-1])
    leaves = []
    if len(words) > 2:
        leaves_str = line.split("-> ", 1)[1]
        leaves = leaves_str.split(", ")
    data.append(Program(name, number, leaves))

# Creates a map: name -> Program instance:
byname = dict()
for program in data:
    byname[program.name] = program

# On the first pass, sets immediate parents of all the entries:
for program in data:
    for leaf in program.leaves:
        byname[leaf].parent = program
# Starting on the first program in the list, descends all the way to the root:
ans1 = data[0]
while ans1.parent is not None:
    ans1 = ans1.parent
print("The root of the tower has name: {}".format(ans1.name))
root = ans1

# Challenge B: finds the source of the unbalance.
def calculate_weight_of_tower(tower_root):
    for leaf in tower_root.leaves:
        calculate_weight_of_tower(byname[leaf])

    tower_root.totalWeight = tower_root.number + sum(byname[leaf].totalWeight for leaf in tower_root.leaves)

calculate_weight_of_tower(root)

analyzed = root
plate = dict()
weights = []
delta_weight = 0

while len(plate) != 1:
    for leaf in analyzed.leaves:
        p = byname[leaf]
        w = p.totalWeight
        if w in plate:
            plate[w].append(p)
        else:
            plate[w] = [p]
            weights.append(w)

    if len(plate) == 1:
        # The program currently analyzed is at fault!
        sane_weight = analyzed.number + delta_weight
        print("The plate {} should have weight {} instead of {}".format(
                analyzed.name, sane_weight, analyzed.number
        ))
    else:
        odd = None
        for k in plate:
            if len(plate[k]) == 1:
                odd = plate[k][0]
        normal_weight = weights[0]
        if normal_weight == odd.totalWeight:
            normal_weight = weights[1]
        delta_weight = normal_weight - odd.totalWeight
        analyzed = odd
        plate = dict()
        weights = []
        print("Descending to unbalanced node {}".format(odd.name))

