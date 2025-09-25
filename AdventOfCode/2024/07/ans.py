DO_EXAMPLE = True
def main():
    A = 0
    B = 0
    equations = []
    good_results = set()
    good_results_b = set()
    with open("example.txt" if DO_EXAMPLE else "input.txt", "r") as file:
        for line in file:
            (s_target, s_numbers) = line.split(":")
            target = int(s_target.strip())
            numbers = [int(x) for x in s_numbers.strip().split(" ")]
            equations.append({"target": target, "numbers": numbers})

            for i in range(pow(2, len(numbers)-1)):
                result = numbers[0]
                op = i
                for slot in range(len(numbers) - 1):
                    r = (result * numbers[slot+1]) if (op % 2 == 1) else (result + numbers[slot+1])
                    op //= 2
                    result = r
                if result == target:
                    good_results.add(result)
                    continue


            for i in range(pow(3, len(numbers) - 1)):
                result = numbers[0]
                op = i
                for slot in range(len(numbers) - 1):
                    if (op % 3 == 0):
                        result += numbers[slot + 1]
                    elif (op % 3 == 1):
                        result *= numbers[slot + 1]
                    else:
                        result = int(str(result) + str(numbers[slot + 1]))
                    op //= 3
                if result == target:
                    good_results_b.add(result)
                    continue

        A = sum(good_results)
        print(f"Answer A: {A}")
        B = sum(good_results_b)
        print(f"Answer B: {B}")

if __name__ == "__main__":
    main()
