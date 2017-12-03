#!/usr/bin/env python3

"""
Solution to the task #03 of the "Advent of Code 2017" series.

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

requested_index = int(input("Requested index: "))
# Finds the ring on which the requested index lies.
# In the bottom-right "corner" of the ring resides the number:
# (2*i + 1) ^ 2.
# This number is the last one found on this ring. I call it a "border" number.
# e.g. ring #0 has "border": 1 == (2*0+1)^2 == 1^2
#      ring #1 has "border": 9 == (2*1+1)^2 == 3^2
#      ring #2 has "border": 25 == (2*2+1)^2 == 5^2
#      etc.
current_ring = 0
current_border = (2 * current_ring + 1) ** 2

while current_border < requested_index:
    current_ring += 1
    current_border = (2 * current_ring + 1) ** 2

print("Number {} is on the ring #{} (border number is {}).".format(
        requested_index, current_ring, current_border))

# Goes back from the "current_border" number, searching for the requested_index.
# It is guaranteed that the requested_index is on current_ring, so at most
# (4*current_ring) elements have to be checked.
# The indicator will be moved according to the following pattern:
# LEFT: current_ring times,
# UP: current_ring times,
# RIGHT: current_ring times,
# DOWN: current_ring times. This walk leaves the cursor in the place of
# where the walk has started.
# It is guaranteed that the requested_index exists on this path, and the walk
# will be terminated as soon as the element is found.
# The offset from the data port (index 1) to current_border is:
# DOWN current_ring times and RIGHT current_ring times.
# The offset to the element pointed to by the indicator will be updated
# simultaneously.
OFFSET = {
    "L": 0,
    "U": 0,
    "R": current_ring,
    "D": current_ring
}
ADJACENT = {
    "L": "R",
    "U": "D",
    "R": "L",
    "D": "U"
}

walk_pattern = []
for i in range(2 * current_ring):
    walk_pattern.append("L")
for i in range(2 * current_ring):
    walk_pattern.append("U")
for i in range(2 * current_ring):
    walk_pattern.append("R")
for i in range(2 * current_ring):
    walk_pattern.append("D")

indicator = current_border
while indicator != requested_index:
    direction = walk_pattern.pop(0)
    adjacent = ADJACENT[direction]
    if OFFSET[adjacent] > 0:
        OFFSET[adjacent] -= 1
    else:
        OFFSET[direction] += 1
    indicator -= 1

directions = []
total_steps = 0
if OFFSET["L"] > 0:
    total_steps += OFFSET["L"]
    directions.append("{} steps to the LEFT".format(OFFSET["L"]))
if OFFSET["R"] > 0:
    total_steps += OFFSET["R"]
    directions.append("{} steps to the RIGHT".format(OFFSET["R"]))
if OFFSET["U"] > 0:
    total_steps += OFFSET["U"]
    directions.append("{} steps to the TOP".format(OFFSET["U"]))
if OFFSET["D"] > 0:
    total_steps += OFFSET["D"]
    directions.append("{} steps to the BOTTOM".format(OFFSET["D"]))

reply = "The number is " + " and ".join(directions) + " from the entry port."

print(reply)
print("The total number of steps necessary to access this datum is {}".format(
        total_steps))


# Challenge #02: find first entry that exceeds N on a spiral:
size = 1001 # this should be an odd number.
mid = size // 2
board = [[0 for col in range(size)] for row in range(size)]
current_pos = [mid, mid]
board[mid][mid] = 1
last_entry = 1
current_length = 1 # the length of current arm of a spiral
found = False

def calculate():
    global found
    if found:     # lame optimization: disallows calculating results
        return    # after the target has been found.
    x, y = current_pos[0], current_pos[1]
    board[x][y] = board[x-1][y-1] + board[x][y-1] + board[x+1][y-1] +\
                  board[x-1][y]                   + board[x+1][y]   +\
                  board[x-1][y+1] + board[x][y+1] + board[x+1][y+1]
    #print("calculating: x={}, y={} == {}".format(x, y, board[x][y]))
    if board[x][y] > requested_index:
        print("First entry exceeding {} is {}".format(
                requested_index, board[x][y]))
        found = True

while current_length < size and not found:
    for i in range(current_length):
        current_pos[0] += 1 # takes one step to the RIGHT
        calculate()
    for i in range(current_length):
        current_pos[1] -= 1 # takes one step to the TOP
        calculate()
    current_length += 1
    for i in range(current_length):
        current_pos[0] -= 1 # takes one step to the LEFT
        calculate()
    for i in range(current_length):
        current_pos[1] += 1 # takes one step to the BOTTOM
        calculate()
    current_length += 1

