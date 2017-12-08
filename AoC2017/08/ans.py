#!/usr/bin/env python3

"""
Solution to the task #08 of the "Advent of Code 2017" series.

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

import math


data_file = "input_08"
#data_file = "sample_input_08"
lines = [line.strip() for line in open(data_file, "r")]

# Each line consists of:
# register_name inc|dec value if register_name2 ==|!=|<|<=|>|>= value2

valid_operations = ["inc", "dec"]
valid_conditionals = ["==", "!=", "<", "<=", ">", ">="]
registers = dict()

# for part B of the challenge:
max_value_ever = 0
max_value_ever_name = None

for line in lines:
    parts = line.split(" ")
    register_name = parts[0]
    operation = parts[1]
    value = int(parts[2])
    # Third element is always "if", this is not checked.
    register_name2 = parts[4]
    conditional = parts[5]
    if conditional not in valid_conditionals:
        raise RuntimeError("Unsupported conditional: {}".format(conditional))
    value2 = int(parts[6])

    if register_name not in registers:
        registers[register_name] = 0
    if register_name2 not in registers:
        registers[register_name2] = 0

    conditionMet = False
    if ((conditional == "==" and registers[register_name2] == value2)
    or (conditional == "!=" and registers[register_name2] != value2)
    or (conditional == "<" and registers[register_name2] < value2)
    or (conditional == "<=" and registers[register_name2] <= value2)
    or (conditional == ">" and registers[register_name2] > value2)
    or (conditional == ">=" and registers[register_name2] >= value2)
    ):
        conditionMet = True

    if conditionMet:
        if operation == "inc":
            registers[register_name] += value
            print("Adding {} to register {}, it now holds {}.".format(
                value, register_name, registers[register_name]))
        elif operation == "dec":
            registers[register_name] -= value
            print("Substracting {} from register {}, it now holds {}.".format(
                value, register_name, registers[register_name]))
        else:
            raise RuntimeError("Unsupported operation: {}".format(operation))
        if registers[register_name] > max_value_ever:
            max_value_ever = registers[register_name]
            max_value_ever_name = register_name

# Challenge #08A:
total_max = float("-inf")
total_max_name = None

for key in registers:
    if registers[key] > total_max:
        total_max = registers[key]
        total_max_name = key
print("The total maximum value is {} in register {}".format(
        total_max, total_max_name))

# Challenge #08B:
print("The maximum temporary value was {} in register {} (it now holds {})".format(
        max_value_ever, max_value_ever_name, registers[max_value_ever_name]))
