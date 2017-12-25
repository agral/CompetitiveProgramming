#!/usr/bin/env python3

"""
Solution to the challenge #23 of the "Advent of Code 2017" series.

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

import os

# Problem constants:
FIRST_REGISTER = "a"
LAST_REGISTER = "h"

class LameAssembler:
    def __init__(self):
        self.registers = dict()
        self.clear()

    def clear(self):
        for key in range(ord(FIRST_REGISTER), ord(LAST_REGISTER)+1):
            self.registers[chr(key)] = 0
        self.program = []
        self.IP = 0
        self.termination_flag = True
        self.mul_count = 0

    def value_of(self, x):
        if x >= FIRST_REGISTER and x <= LAST_REGISTER:
            return self.registers[x]
        else:
            try:
                return int(x)
            except:
                raise ValueError("Not a number nor valid register: {}".format(x))

    def interpret(self, line_of_code):
        if self.termination_flag:
            return

        tokens = line_of_code.split(" ")
        if   tokens[0] == "set":
            self.registers[tokens[1]] = self.value_of(tokens[2])
            self.IP += 1
        elif tokens[0] == "sub":
            self.registers[tokens[1]] -= self.value_of(tokens[2])
            self.IP += 1
        elif tokens[0] == "mul":
            self.registers[tokens[1]] *= self.value_of(tokens[2])
            self.IP += 1
            self.mul_count += 1
        elif tokens[0] == "jnz":
            if self.value_of(tokens[1]) == 0:
                self.IP += 1
            else:
                self.IP += self.value_of(tokens[2])
        else:
            raise RuntimeError("Unsupported instruction: {}".format(tokens[0]))

        self.step += 1

        if self.IP < 0 or self.IP >= len(self.program):
            print("The program has terminated successfully.")
            self.termination_flag = True

    def load_program(self, source_file):
        self.program = [line.strip() for line in open(source_file, "r")]
        self.IP = 0
        self.step = 0
        self.termination_flag = False

    def run_program(self):
        while not self.termination_flag:
            self.interpret(self.program[self.IP])

    def debug_program(self, breakpoints=[]):
        while not self.termination_flag:
            # debugflag = debugflag or (self.registers["g"] > -10 and self.registers["e"] > 108000)
            #if self.IP in breakpoints:
            if self.step >= 864790:
                print("Breakpoint: {}".format(self.IP))
                self.print()
                input()
            self.interpret(self.program[self.IP])

    def print_registers(self):
        print("Registers: ", end="")
        for key in range(ord(FIRST_REGISTER), ord(LAST_REGISTER)+1):
            print("[{}]={}\t".format(chr(key),
                    self.registers[chr(key)]), end="")
        print()

    def print(self):
        os.system("clear")
        self.print_registers()
        print("Next instruction: ({}) {}, step: {}".format(
                self.IP, self.program[self.IP], self.step))
        print("\nCurrent program:")
        for lnum in range(len(self.program)):
            print("{:2d}: {}\t\t\t{}".format(
                    lnum, self.program[lnum], "<" if lnum == self.IP else " "))


program_file = "input_23"
optimized_program_file = "optimized_program"

la = LameAssembler()
la.load_program(program_file)
la.run_program()
print("[Challenge A]: MUL instruction invoked {} times.".format(la.mul_count))

la = LameAssembler()
la.load_program(optimized_program_file)
la.registers["a"] = 1
la.debug_program([11])
print("[Challenge B]: registers[\"h\"] == {}".format(la.registers["h"]))
