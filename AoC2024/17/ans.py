DO_EXAMPLE = True
DO_EXAMPLE = False
def main():

    with open("example_2024.txt" if DO_EXAMPLE else "input.txt", "r") as file:
        input = [line.strip() for line in file]

        # not anymore, see note below combo definition.
        #reg = { # the three registers
        #    "a": int(input[0].split(": ")[1]),
        #    "b": int(input[1].split(": ")[1]),
        #    "c": int(input[2].split(": ")[1])
        #}
        ip = 0
        is_halted = False
        program = [int(num) for num in input[4].split(": ")[1].split(",")]
        output = []

        # Combo operands:
        # 0 to 3 : literal 0 to 3
        # 4 - register A
        # 5 - register B
        # 6 - register C
        # 7 - disallowed.
        combo = [0, 1, 2, 3, int(input[0].split(": ")[1]),
                 int(input[1].split(": ")[1]), int(input[2].split(": ")[1])]
        # Don't know how to pass by reference in Python, so combo[4] to combo[6] are
        # now literally registers A to C.

        while not is_halted:
            if program[ip] == 0: # adv (division of A register)
                combo[4] //= 2 ** combo[program[ip + 1]]
                ip += 2
            elif program[ip] == 1: # bxl (bitwise XOR of B and literal operand)
                combo[5] ^= program[ip + 1]
                ip += 2
            elif program[ip] == 2: # bst (set B to value of combo operand modulo eight)
                combo[5] = combo[program[ip + 1]] % 8
                ip += 2
            elif program[ip] == 3: # jnz (jump-not-zero on A register)
                if combo[4] == 0:
                    ip += 2 # if this instruction does NOT jump, IP is increased normally.
                else:
                    ip = program[ip + 1]
                    # ip is not increased by two after a jump.
            elif program[ip] == 4: # bxc (bitwise-xor B and C, store in B. So B ^= C)
                combo[5] ^= combo[6]
                # this instruction "reads" an operand but "ignores" it.
                ip += 2
            elif program[ip] == 5: # out (append combo operand modulo eight to the output)
                output.append(combo[program[ip + 1]] % 8)
                ip += 2
            elif program[ip] == 6: # bdv (like adv, but store result in B register)
                combo[5] = combo[4] // 2 ** combo[program[ip + 1]]
                ip += 2
            elif program[ip] == 7: # cdv (like adv, but store result in C register)
                combo[6] = combo[4] // 2 ** combo[program[ip + 1]]
                ip += 2

            if ip >= len(program):
                is_halted = True

        A = ",".join(str(num) for num in output)
        print(f"Answer A: {A}")

        # This brute-force approach won't work - TLE.
        is_quine = False
        regA = 0
        while not is_quine:
            if (regA % 100000 == 0):
                print(f"Trying out regA={regA}")
            regA += 1
            ip = 0
            is_halted = False
            output = []
            combo = [0, 1, 2, 3, regA,
                 int(input[1].split(": ")[1]), int(input[2].split(": ")[1])]

            while not is_halted:
                if program[ip] == 0: # adv (division of A register)
                    combo[4] //= 2 ** combo[program[ip + 1]]
                    ip += 2
                elif program[ip] == 1: # bxl (bitwise XOR of B and literal operand)
                    combo[5] ^= program[ip + 1]
                    ip += 2
                elif program[ip] == 2: # bst (set B to value of combo operand modulo eight)
                    combo[5] = combo[program[ip + 1]] % 8
                    ip += 2
                elif program[ip] == 3: # jnz (jump-not-zero on A register)
                    if combo[4] == 0:
                        ip += 2 # if this instruction does NOT jump, IP is increased normally.
                    else:
                        ip = program[ip + 1]
                        # ip is not increased by two after a jump.
                elif program[ip] == 4: # bxc (bitwise-xor B and C, store in B. So B ^= C)
                    combo[5] ^= combo[6]
                    # this instruction "reads" an operand but "ignores" it.
                    ip += 2
                elif program[ip] == 5: # out (append combo operand modulo eight to the output)
                    output.append(combo[program[ip + 1]] % 8)
                    ip += 2
                elif program[ip] == 6: # bdv (like adv, but store result in B register)
                    combo[5] = combo[4] // 2 ** combo[program[ip + 1]]
                    ip += 2
                elif program[ip] == 7: # cdv (like adv, but store result in C register)
                    combo[6] = combo[4] // 2 ** combo[program[ip + 1]]
                    ip += 2

                if ip >= len(program):
                    is_halted = True

            is_quine = (output == program)
        print(f"Answer B: {regA}")

if __name__ == "__main__":
    main()
