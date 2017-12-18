#!/usr/bin/env python3

"""
Solution to the challenge #15 of the "Advent of Code 2017" series.

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


class Generator:
    def __init__(self, seed, factor, discriminator=1):
        self.seed = seed
        self.previous = seed
        self.factor = factor
        self.discriminator = discriminator
        self.MAX = 2147483647

    def generate(self):
        val = self.previous * self.factor % self.MAX
        self.previous = val
        return val

    def generate2(self):
        while(True):
            val = self.previous * self.factor % self.MAX
            self.previous = val
            if val % self.discriminator == 0:
                return val


def judge(val1, val2):
    v1 = val1 % 65536 # clips the value to the least 16 bits.
    v2 = val2 % 65536
    return 1 if v1 == v2 else 0

# Challenge A:
ROUNDS = 40 * 1000 * 1000
g1 = Generator(679, 16807)
g2 = Generator(771, 48271)

score = 0
for i in range(ROUNDS):
    score += judge(g1.generate(), g2.generate())
print(score)

# Challenge B:
ROUNDS = 5 * 1000 * 1000
g1 = Generator(679, 16807, 4)
g2 = Generator(771, 48271, 8)

score = 0
for i in range(ROUNDS):
    score += judge(g1.generate2(), g2.generate2())
print(score)
