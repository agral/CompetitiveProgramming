#!/usr/bin/env python3

"""
Solution to the challenge #24 of the "Advent of Code 2017" series.

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

class Component:
    def __init__(self, portA, portB):
        self.A = portA
        self.B = portB
        self.strength = self.A + self.B

    def __str__(self):
        return "{}/{}".format(self.A, self.B)

    def __repr__(self):
        return self.__str__()

    def match(self, number):
        return self.A == number or self.B == number

class Bridge:
    def __init__(self, list_of_links):
        self.length = 0
        self.terminator = 0
        self.links = []
        for link in list_of_links:
            self.add(link)

    def __str__(self):
        return "--".join([str(link) for link in self.links])

    def add(self, component):
        if self.terminator == component.A:
            self.terminator = component.B
        else:
            self.terminator = component.A
        self.links.append(component)

    def strength(self):
        result = 0
        for link in self.links:
            result += link.A + link.B
        return result


data_files = ["sample_input_24", "input_24"]
for data_file in data_files:
    data_lines = [line.strip() for line in open(data_file, "r")]

    components = set()
    for line in data_lines:
        a, b = line.split("/")
        components.add(Component(int(a), int(b)))

    print(len(components))

    bridge = Bridge([])
    new_bridges = [bridge]
    all_bridges = [bridge]
    strongest = 0

    counter = 1
    while len(new_bridges) > 0:
        print("Round #{}".format(counter))

        extended_bridges = []
        for bridge in new_bridges:
            matching = [link for link in components if link.match(
                    bridge.terminator) and not link in bridge.links]
            for link in matching:
                b = Bridge(bridge.links)
                b.add(link)
                extended_bridges.append(b)
        for bridge in extended_bridges:

            b = Bridge(bridge.links)
            s = b.strength()
            if s > strongest:
                strongest = s
            all_bridges.append(b)
        new_bridges = extended_bridges[:]
        counter += 1

    print("[Challenge A] Strongest bridge: {}".format(strongest))

    counter -= 2 # because the last round did not add any bridges.
    longest = [b for b in all_bridges if len(b.links) == counter]
    strongest = 0
    for b in longest:
        s = b.strength()
        if s > strongest:
            strongest = s
    print("[Challenge B] Strongest bridge of length={}: {}".format(
            counter, strongest))
