#!/usr/bin/env python3

"""
Solution to the challenge #18 of the "Advent of Code 2017" series.

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

import time

class LameAssemblerA:
    def __init__(self):
        self.registers = dict()
        self.clear_all()

    def clear_all(self):
        for key in range(ord("a"), ord("z")+1):
            self.registers[chr(key)] = 0
            self.program = []
            self.IP = 0
            self.last_sound = None
            self.termination_flag = False

    def set(self, register, value):
        self.registers[register] = value

    def snd(self, frequency):
        self.last_sound = frequency

    def add(self, register, delta):
        self.registers[register] += delta

    def mul(self, registerX, factor ):
        self.registers[registerX] *= factor

    def mod(self, registerX, divisor):
        self.registers[registerX] %= divisor

    def rcv(self, conditional):
        if conditional != 0:
            self.termination_flag = True
            print("rcv: returning {}".format(self.last_sound))
            return self.last_sound

    def jgz(self, conditional, offset):
        if conditional > 0:
            self.IP += offset

            # Negates the increment of IP after every instruction -
            # jump instruction should not modify the stack:
            self.IP -= 1

    def value_of(self, number_or_register):
        if number_or_register >= "a" and number_or_register <= "z":
            return self.registers[number_or_register]
        else:
            try:
                return int(number_or_register)
            except ValueError:
                raise RuntimeError("Not a number nor register: {}".format(
                        number_or_register))
            except:
                print("Unexpected error!")
                raise

    def interpret(self, line_of_code):
        #print("INTERPRET: {}, IP={}".format(line_of_code, self.IP))
        chunks = line_of_code.split(" ")
        if   chunks[0] == "set":
            self.set(chunks[1], self.value_of(chunks[2]))
        elif chunks[0] == "snd":
            self.snd(self.value_of(chunks[1]))
        elif chunks[0] == "add":
            self.add(chunks[1], self.value_of(chunks[2]))
        elif chunks[0] == "mul":
            self.mul(chunks[1], self.value_of(chunks[2]))
        elif chunks[0] == "mod":
            self.mod(chunks[1], self.value_of(chunks[2]))
        elif chunks[0] == "rcv":
            self.rcv(self.value_of(chunks[1]))
        elif chunks[0] == "jgz":
            self.jgz(self.value_of(chunks[1]), self.value_of(chunks[2]))

        self.IP += 1

    def import_program_from_file(self, filename):
        self.program = [line.strip() for line in open(filename, "r")]
        self.IP = 0

    def run_program(self):
        while (not self.termination_flag) \
                and self.IP < len(self.program) and self.IP >= 0:
            self.interpret(self.program[self.IP])


# FFFFFUUUUUUUUUUuuuu-, you've changed too much of the requirements,
# and I don't particularly want to break the existing & working implementation
# of LameAssembler for part A. Duplicating everything:
programs = []

class LameAssemblerB:
    def __init__(self, program_id, snd_target_id):
        self.id = program_id
        self.target = snd_target_id
        self.registers = dict()
        self.clear_all()
        self.registers["p"] = program_id


    def clear_all(self):
        print("CLEAR_ALL for #{}".format(self.id))
        for key in range(ord("a"), ord("z")+1):
            self.registers[chr(key)] = 0
            self.program = []
            self.IP = 0
            self.termination_flag = False
            self.recv_queue = []
            self.waiting_for_input = False
            self.snd_count = 0


    def set(self, register, value):
        self.registers[register] = value
        self.IP += 1


    def snd(self, value):
        print("Program #{} sends data: {} to program #{}.".format(
            self.id, value, self.target))
        programs[self.target].recv_queue.append(value)
        self.snd_count += 1
        self.IP += 1


    def add(self, register, delta):
        self.registers[register] += delta
        self.IP += 1


    def mul(self, registerX, factor ):
        self.registers[registerX] *= factor
        self.IP += 1


    def mod(self, registerX, divisor):
        self.registers[registerX] %= divisor
        self.IP += 1


    def rcv(self, register):
        print(self.recv_queue)
        if len(self.recv_queue) > 0:
            self.registers[register] = self.recv_queue.pop(0)
            self.IP += 1
            print("RECV: register {} has been set to {}".format(
                register, self.registers[register]))
        else:
            print("Program #{} pauses, waiting to receive input.".format(
                self.id))
            self.waiting_for_input = True


    def jgz(self, conditional, offset):
        if conditional > 0:
            self.IP += offset
            # Successful JUMP instructions do not increment the IP.
        else:
            # Failed conditional, however, increments the IP:
            self.IP += 1

    def value_of(self, number_or_register):
        if number_or_register >= "a" and number_or_register <= "z":
            return self.registers[number_or_register]
        else:
            try:
                return int(number_or_register)
            except ValueError:
                raise RuntimeError("Not a number nor register: {}".format(
                        number_or_register))
            except:
                print("Unexpected error!")
                raise


    def interpret(self, line_of_code):
        # print("Interpret(#{}, {})".format(self.id, line_of_code))
        chunks = line_of_code.split(" ")
        if   chunks[0] == "set":
            self.set(chunks[1], self.value_of(chunks[2]))
        elif chunks[0] == "snd":
            self.snd(self.value_of(chunks[1]))
        elif chunks[0] == "add":
            self.add(chunks[1], self.value_of(chunks[2]))
        elif chunks[0] == "mul":
            self.mul(chunks[1], self.value_of(chunks[2]))
        elif chunks[0] == "mod":
            self.mod(chunks[1], self.value_of(chunks[2]))
        elif chunks[0] == "rcv":
            self.rcv(chunks[1])
        elif chunks[0] == "jgz":
            self.jgz(self.value_of(chunks[1]), self.value_of(chunks[2]))

        # time.sleep(0.2)

    def import_program_from_file(self, filename):
        self.program = [line.strip() for line in open(filename, "r")]
        self.IP = 0

    def run_program(self):
        while (not self.termination_flag) \
              and ((not self.waiting_for_input) or len(self.recv_queue) > 0) \
              and (self.IP >= 0 and self.IP < len(self.program)):

            self.interpret(self.program[self.IP])

        if self.IP < 0 or self.IP >= len(self.program):
            print("Program #{} has finished its execution successfully.".format(
                    self.id))
            self.termination_flag = True



data_files = ["sample_input_18", "input_18"]
for data_file in data_files:
    print("Parsing data file: {}".format(data_file))

    print("[Challenge 18A]: ", end="")
    la = LameAssemblerA()
    la.import_program_from_file(data_file)
    la.run_program()

    print("[Challenge 18B]: ")
    programs = []
    programs.append(LameAssemblerB(0, 1))
    programs[0].import_program_from_file(data_file)
    programs.append(LameAssemblerB(1, 0))
    programs[1].import_program_from_file(data_file)

    while True:
        programs[0].run_program()
        programs[1].run_program()

        if programs[0].waiting_for_input and len(programs[0].recv_queue) > 0:
                print("program #0 resumes execution.")
                programs[0].waiting_for_input = False
        if programs[1].waiting_for_input and len(programs[1].recv_queue) > 0:
                print("program #1 resumes execution.")
                programs[1].waiting_for_input = False

        if programs[0].waiting_for_input and programs[1].waiting_for_input:
            print("[Scheduler]: Execution of programs resulted in DEADLOCK.")
            break
        elif programs[0].termination_flag and programs[1].termination_flag:
            print("[Scheduler]: Both programs terminated successfully.")
            break
    print("[Scheduler]: Program #1 called snd() {} times.".format(
            programs[1].snd_count))
