DO_EXAMPLE = True
DO_EXAMPLE = False

signals = {}
gates = []

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

    def __repr__(self):
        return f"{self.gate_type} ({self.input1}, {self.input2}) -> {self.output}; result={self.result}"

def get_z_value():
    bits = []
    z_bit = 0
    while True:
        signal_name = f"z{z_bit:02d}"
        if not signal_name in signals:
            break
        bits.append(signals[signal_name])
        z_bit += 1
    ans = 0
    for bit in reversed(bits):
        ans *= 2
        ans += (1 if bit else 0)
    return ans

def main():
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

        processed_gates = 0
        while processed_gates != len(gates):
            for gate in gates:
                if gate.result == None and gate.is_ready():
                    gate.logic()
                    signals[gate.output] = gate.result
                    processed_gates += 1

        print(f"Answer A: {get_z_value()}")

        print(f"Answer B: {B}")

if __name__ == "__main__":
    main()
