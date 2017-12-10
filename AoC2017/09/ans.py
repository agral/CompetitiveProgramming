#!/usr/bin/env python3

"""
Solution to the challenge #09 of the "Advent of Code 2017" series.

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

def without_canceled_chars(text):
    next_bang = text.find("!")
    while next_bang != -1:
        without_this_bang = text[:next_bang] + text[next_bang+2:]
        text = without_this_bang
        next_bang = text.find("!")
    return text


def without_garbage(text):
    next_garbage_start = text.find("<")
    garbage_counter = 0
    while next_garbage_start != -1:
        next_garbage_end = text.find(">")
        if next_garbage_end == -1:
            raise RuntimeError("Unterminated garbage")
        without_this_garbage = \
                text[:next_garbage_start] + text[next_garbage_end+1:]
        text = without_this_garbage
        garbage_counter += (next_garbage_end - next_garbage_start - 1)
        next_garbage_start = text.find("<")
    print("Removed {} garbage characters.".format(garbage_counter))
    return text


def calculate_score(text):
    current_score = 0
    current_level = 0
    for char in text:
        if char == '{':
            current_level += 1
        elif char == '}':
            current_score += current_level
            current_level -= 1

    if current_level != 0:
        raise RuntimeError("Unterminated groups")

    return current_score


data_file = "sample_input_09"
data_file = "input_09"

lines = [line.strip() for line in open(data_file, "r")]

for line in lines:
    #print(line)
    text = without_garbage(without_canceled_chars(line))
    #print(text)
    print(calculate_score(text))
    print()
