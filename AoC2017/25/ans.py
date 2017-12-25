#!/usr/bin/env python3

"""
Solution to the challenge #25 of the "Advent of Code 2017" series.

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

class Tape:
    def __init__(self):
        self.t = dict()
        self.position = 0

    def move(self, direction):
        if direction == "left":
            self.position -= 1
        elif direction == "right":
            self.position += 1
        else:
            raise ValueError("Invalid direction: {}".format(direction))

    def get(self):
        result = 0
        try:
            result = self.t[self.position]
        except KeyError:
            pass
        return result

    def write(self, value):
        self.t[self.position] = value

    def count_value_of(self, target_value):
        result = 0
        for _, v in self.t.items():
            if v == target_value:
                result += 1
        return result


class TuringMachine:
    def __init__(self, rules_file):
        self.state = None
        self.logic = dict()
        self.load_rules(rules_file)
        self.tape = Tape()
        self.step_counter = 0

    def load_rules(self, rules_file):
        lines = [line.strip() for line in open(rules_file, "r")]
        self.state = lines[0].split(" ")[-1][0]
        self.target_steps = int(lines[1].split(" ")[5])

        l = 3
        while l < len(lines):
            state = lines[l].split(" ")[-1][0]
            cond0_write_val = int(lines[l+2].split(" ")[-1][0])
            cond0_move_dir = lines[l+3].split(" ")[-1][:-1]
            cond0_next_state = lines[l+4].split(" ")[-1][0]

            cond1_write_val = int(lines[l+6].split(" ")[-1][0])
            cond1_move_dir = lines[l+7].split(" ")[-1][:-1]
            cond1_next_state = lines[l+8].split(" ")[-1][0]

            self.logic[state] = {
                0: {
                    "write_value": cond0_write_val,
                    "move_direction": cond0_move_dir,
                    "next_state": cond0_next_state
                },
                1: {
                    "write_value": cond1_write_val,
                    "move_direction": cond1_move_dir,
                    "next_state": cond1_next_state
                }
            }

            l += 10

    def step(self):
        val = self.tape.get()
        self.tape.write(self.logic[self.state][val]["write_value"])
        self.tape.move(self.logic[self.state][val]["move_direction"])
        self.state = self.logic[self.state][val]["next_state"]
        self.step_counter += 1

    def run(self):
        while self.step_counter != self.target_steps:
            self.step()
        print("After {} steps, {} fields contain 1.".format(
                self.step_counter, self.tape.count_value_of(1)))

data_files = ["sample_input_25", "input_25"]
for data_file in data_files:
    print("Working with file: {}".format(data_file))
    t = TuringMachine(data_file)
    t.run()
