DO_EXAMPLE = False

def is_fitting(lock, key):
    for pin_number in range(5):
        if lock[pin_number] + key[pin_number] > 5:
            return False
    return True

def main():
    A = 0
    B = 0
    with open("example.txt" if DO_EXAMPLE else "input.txt", "r") as file:
        lines = [line.strip() for line in file]
        l = 0
        locks = []
        keys = []
        while l < len(lines):
            isLock = (lines[l] == "#####")
            pins = []
            for x in range(5):
                height = 0
                for y in range(1, 5+1):
                    if lines[l+y][x] == "#":
                        height += 1
                pins.append(height)
            if isLock:
                locks.append(pins)
            else:
                keys.append(pins)
            l += 8 # 7 for the lock/key schematics, 1 for the empty line.

        # try every key with every lock:
        for lock in locks:
            for key in keys:
                if is_fitting(lock, key):
                    A += 1

        print(f"Answer A: {A}")

        print(f"Answer B: {B}")

if __name__ == "__main__":
    main()
