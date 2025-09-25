import copy

DO_EXAMPLE = True
DO_EXAMPLE = False

signals = {}
gates = []
suspect_gates = []

class Gate:
    def __init__(self, gate_type, input1, input2, output):
        self.gate_type = gate_type
        self.input1 = input1
        self.input2 = input2
        self.output = output
        self.result = None

    def is_ready(self):
        return True if (self.input1 in signals and self.input2 in signals) else False

    def logic(self):
        if self.gate_type == "AND":
            self.result = (signals[self.input1] and signals[self.input2])
        elif self.gate_type == "OR":
            self.result = (signals[self.input1] or signals[self.input2])
        elif self.gate_type == "XOR":
            self.result = (signals[self.input1] != signals[self.input2])
        else:
            raise ValueError(f"Unknown gate type: {self.gate_type}")

        if self.output[0] == 'z':
            num = int(self.output[1:])
            if self.result != signals[f"exp{num:02d}"]:
                suspect_gates.append(self)
                print(f"Incorrect value found for z={num}; {self}")

    def __repr__(self):
        return f"{self.gate_type} ({self.input1}, {self.input2}) -> {self.output}; result={self.result}"

def get_value(wire_name):
    bits = []
    bit_num = 0
    while True:
        signal_name = f"{wire_name}{bit_num:02d}"
        if not signal_name in signals:
            break
        bits.append(signals[signal_name])
        bit_num += 1
    ans = 0
    for bit in reversed(bits):
        ans *= 2
        ans += (1 if bit else 0)
    return ans

def to_binary(num):
    target_bits = []
    while num > 0:
        target_bits.append(num % 2)
        num >>= 1
    target_bits.reverse()
    return target_bits

def swap_outputs(gate_lhs, gate_rhs):
    gate_lhs.output, gate_rhs.output = gate_rhs.output, gate_lhs.output

def main():
    global signals
    global gates

    B = 0
    with open("example.txt" if DO_EXAMPLE else "input.txt", "r") as file:
        is_reading_inputs = True
        for line in file:
            if is_reading_inputs:
                if line == "\n":
                    is_reading_inputs = False
                else:
                    signal, s_value = line.strip().split(": ")
                    signals[signal] = True if s_value == "1" else False
            else:
                s_gate, output = line.strip().split(" -> ")
                inp1, gate_type, inp2 = s_gate.split(" ")
                gates.append(Gate(gate_type, inp1, inp2, output))

        # This is for part B. Needs to be set up before processing the gates.
        # Maybe start checking which gate set the first incorrect output byte?
        # -> but this may be caused indirectly. Let's try this.
        expected_output = [(True if val == 1 else False) for val in to_binary(get_value('x')+get_value('y'))]
        for i in range(len(expected_output)):
            sig = f"exp{i:02d}" # manually checked, no "exp" in input signals.
            signals[sig] = expected_output[len(expected_output)-1-i] # bit hacky but works

        original_signals = copy.deepcopy(signals)
        original_gates = gates[:]

        processed_gates = 0
        while processed_gates != len(gates):
            old_pg = processed_gates
            for gate in gates:
                if gate.result == None and gate.is_ready():
                    gate.logic()
                    signals[gate.output] = gate.result
                    processed_gates += 1
            if old_pg == processed_gates:
                print("Gates will not settle in this configuration.")
                processed_gates = len(gates) # to break the loop

        print(f"Answer A: {get_value('z')}")
        print(to_binary(get_value('z')))
        print(to_binary(get_value('x')+get_value('y')))

        # Part B: Surely it can be done smarter than by bruteforce?
        print("Part B:")
        print("".join([str(c) for c in to_binary(get_value('z'))]))
        print("".join([str(c) for c in to_binary(get_value('x')+get_value('y'))]))
        # The byte patterns of actual and expected outputs are as follows:
        # bitnum:      40|     32|     24|     16|       8       0
        # actual:   1101011111110111000100110001011010101101001110
        # expected: 1101100000010111000100110010011010001011001110
        # invalid?:     XXXXXXX               XX      X  XX

        if False: # Tried this approach, wouldn't work - for reason why, see next comment.
            signals = copy.deepcopy(original_signals)
            gates = original_gates[:]
            swap_outputs(suspect_gates[0], suspect_gates[5])

            processed_gates = 0
            while processed_gates != len(gates):
                old_pg = processed_gates
                for gate in gates:
                    if gate.result == None and gate.is_ready():
                        gate.logic()
                        signals[gate.output] = gate.result
                        processed_gates += 1
                if old_pg == processed_gates:
                    print("Gates will not settle in this configuration.")
                    break

        # No, gates won't settle because this circuit must have been realized with full adders!
        # (so higher output bytes will only settle after lower output bytes).
        # Let's see if there's a half adder realized in the input:
        # x00 XOR y00 -> z00 \
        # x00 AND y00 -> gmk / this forms a correct half adder with carry_out_0 outputted to gmk
        #
        # Bits #0 and #1 are correct. Expect to see a full adder for z01.
        # vgr XOR gmk -> z01 \ correct, double-xored x01 ^ y01 ^ carry_out_0
        # y01 XOR x01 -> vgr | correct, XOR comparison of raw input bits
        # vgr AND gmk -> fsm / correct, fsm is carry_out_1, and that's a full adder.
        #
        # From correctly-outputting bit #6:
        # btp XOR hsm -> z06
        # x06 XOR y06 -> btp
        # btp AND hsm -> jtg  # jtg is carry_out_6 / carry_in_7.
        #
        # Let's analyze the first incorrect one - bit #7:
        # y07 AND x07 -> z07 \ obviously wrong, should be XOR of x07, y07 and carry_in_7.
        # x07 XOR y07 -> mvw | maybe right, so now z07 should be set to mvw xor hsm
        # It's not, but there's this:
        # x06 AND y06 -> hgj
        # jtg OR hgj -> pmc
        # (side note: for carry_out pin, OR gate can be used instead of XOR)
        # To be continued, probably after christmas.

        print(f"Answer B: {B}")
        return


if __name__ == "__main__":
    main()
