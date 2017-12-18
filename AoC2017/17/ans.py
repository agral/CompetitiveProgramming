#!/usr/bin/env python3

"""
Solution to the challenge #17 of the "Advent of Code 2017" series.

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

ITERATIONS_A = 2017
ITERATIONS_B = 50000000

class Spinlock:
    def __init__(self, step):
        self.buffer = [0]
        self.current = 0
        self.pos = 0
        self.step = step

    def insert(self):
        self.current += 1
        self.pos = (self.pos + self.step) % len(self.buffer)
        self.buffer.insert(self.pos + 1, self.current)
        self.pos += 1

    def print(self):
        print(self.buffer)

    def peek(self):
        print(self.buffer[self.pos-3:self.pos+4])

    def after_zero(self):
        zero_index = self.buffer.index(0)
        next_index = (zero_index + 1) % len(self.buffer)
        return self.buffer[next_index]

# Challenge A:
s = Spinlock(3)    # Sample config
s = Spinlock(380)  # Own puzzle input
for i in range(ITERATIONS_A):
    s.insert()
print("[Challenge A]: ", end="")
s.peek()


""" Super-fast spinlock emulator that only keeps track of what is the value
    stored immediately after 0.
    How does this work:
    in every Spinlock, the value at the very beginning (zeroth index)
    is always 0. So the value immediately after 0 is always at index 1.
    This emulator does not actually store items, it only remembers how many
    items were already put in, and what value resides at index 1."""
class SuperFastSpinlockForChallengeB:
    def __init__(self, step):
        self.size = 1 # at the beginning there is only one value - 0.
        self.current = 0
        self.pos = 0
        self.step = step
        self.first = None

    def insert(self):
        self.current += 1
        self.pos = (self.pos + self.step) % self.size
        if self.pos == 0:
            self.first = self.current
        self.pos += 1
        self.size += 1

# Challenge B:
s = SuperFastSpinlockForChallengeB(380)
for i in range(ITERATIONS_B):
    if i % 100000 == 0:
        print("Calculating... {}/{} ({:.1f}%) done so far\r".format(
            i, ITERATIONS_B, 100 * i / ITERATIONS_B), end="")
    s.insert()
print()
print("[Challenge B]: {}".format(s.first))
