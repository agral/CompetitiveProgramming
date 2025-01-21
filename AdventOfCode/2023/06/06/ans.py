import re

def main():
    A = 0
    B = 0
    #with open("example.txt", "r") as file:
    with open("input.txt", "r") as file:
        line = re.sub(" +", " ", file.readline())
        (_, stimes) = [s.strip() for s in line.split(":")]
        line = re.sub(" +", " ", file.readline())
        (_, sdistances) = [s.strip() for s in line.split(":")]

        times = [int(x) for x in stimes.split(" ")]
        distances = [int(x) for x in sdistances.split(" ")]

        ways = [0] * len(times)
        for n in range(len(times)):
            for k in range(1, times[n]): # from 1 to nth_time-1
                distance = k * (times[n] - k)
                if distance > distances[n]:
                    ways[n] += 1
        A = 1
        for num in ways:
            A *= num
        print(A)

        time = int(stimes.replace(" ", ""))
        best_distance = int(sdistances.replace(" ", ""))
        for k in range(1, time):
            distance = k * (time - k)
            if distance > best_distance:
                B += 1
        print(B)

if __name__ == "__main__":
    main()
