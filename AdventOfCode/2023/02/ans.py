def sum_A(line):
    game_info, gamelist = line.split(":")
    game_num = int(game_info.split(" ")[1])
    games = [game.strip() for game in gamelist.split(";")]
    for game in games:
        sballs = [kind.strip() for kind in game.split(",")]
        for sball in sballs:
            count, kind = (item.strip() for item in sball.split(" "))
            count = int(count)
            if (
                (count > 14) or
                ((count > 13) and (kind == "green")) or
                ((count > 12) and (kind == "red"))
                ):
                print(f"[FAIL] Game #{game_num} not possible")
                return 0

    print(f"[OK] Game #{game_num} possible")
    return game_num

def calc_B(line):
    _, gamelist = line.split(":")
    games = [game.strip() for game in gamelist.split(";")]
    mins = {"red": 0, "green": 0, "blue": 0}
    for game in games:
        sballs = [kind.strip() for kind in game.split(",")]
        for sball in sballs:
            count, kind = (item.strip() for item in sball.split(" "))
            count = int(count)
            mins[kind] = max(mins[kind], count)

    return mins["red"] * mins["green"] * mins["blue"]

def main():
    A = 0
    B = 0
    #with open("example.txt", "r") as file:
    with open("input.txt", "r") as file:
        for line in file:
            A += sum_A(line)
            B += calc_B(line)
    print(A)
    print(B)

if __name__ == "__main__":
    main()
