#!/usr/bin/env python3

"""
Solution to the challenge #12 from the "Advent of Code 2017" series.

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

data_files = ["sample_input_12", "input_12"]

for data_file in data_files:
    lines = [line.strip() for line in open(data_file, "r")]

    pipes = []
    counter = 0

    for line in lines:
        raw = line.split(" <-> ")
        current = int(raw[0])
        if not current == counter:
            raise RuntimeError("Current is not counter in line #{}".format(
                    current+1))
        connections_str = raw[1].split(", ")
        connections = [int(connection) for connection in connections_str]
        pipes.append({
                "value": current,
                "connections": connections,
                "visited": False,
                "group": None
        })
        counter += 1

    # Challenge A:
    connected = set()
    def bfs(node):
        if not node["visited"]:
            node["visited"] = True
            connected.add(node["value"])
            for number in node["connections"]:
                bfs(pipes[number])
    bfs(pipes[0])
    #print(connected)
    print("A: Program #0 is connected with {} other programs.".format(
            len(connected)))

    # Challenge B:
    for node in pipes:
        node["visited"] = False

    groups = set()
    def bfs2(node, group):
        if not node["visited"]:
            node["visited"] = True
            node["group"] = group
            groups.add(group)
            for number in node["connections"]:
                bfs2(pipes[number], group)

    for node in pipes:
        if not node["visited"]:
            bfs2(node, node["value"])

    #print(groups)
    print("B: There are {} mutually exclusive groups.".format(len(groups)))

