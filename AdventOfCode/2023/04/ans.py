import re
class Card:
    def __init__(self, number, winners, owned):
        self.number = number
        self.count = 1
        self.winners = winners
        self.owned = owned

    def get_score(self):
        count = self.get_num_wins()
        if count == 0:
            return 0
        return 2 ** (count - 1)

    def get_num_wins(self):
        return len(set(self.winners).intersection(self.owned))


    def __str__(self):
        return (f"Card(num={self.number}, count={self.count},\nwinners=[" +
        ", ".join([str(w) for w in self.winners]) + "]\nowned=[" +
        ", ".join([str(o) for o in self.owned]) + "]\n" +
        f"num_wins: {self.get_num_wins()}\n"
        )
    def __repr__(self):
        return self.__str__()

def main():
    A = 0
    B = 0
    last_card_num = 0
    #with open("example.txt", "r") as file:
    with open("input.txt", "r") as file:
        cards = dict()
        for line in file:
            line = re.sub(' +', ' ', line)
            (s_num, s_card_data) = [s.strip() for s in line.split(":")]
            card_number = int(s_num.split(" ")[1].strip())
            last_card_num = max(last_card_num, card_number)
            (s_winners, s_owned) = [s.strip() for s in s_card_data.split("|")]
            winners = [int(x) for x in s_winners.split(" ") if len(x) > 0]
            print(winners)
            owned = [int(x) for x in s_owned.split(" ") if len(x) > 0]
            print(owned)
            cards[card_number] = Card(card_number, winners, owned)
            A += cards[card_number].get_score()
        print(A)

        for i in range(1, last_card_num + 1):
            for c in range(i+1, i + 1 + cards[i].get_num_wins()):
                cards[c].count += cards[i].count
            B += cards[i].count
        print(B)

if __name__ == "__main__":
    main()
