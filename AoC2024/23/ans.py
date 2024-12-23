from collections import defaultdict

DO_EXAMPLE = True
DO_EXAMPLE = False

def main():
    with open("example.txt" if DO_EXAMPLE else "input.txt", "r") as file:
        # maps computer name to a set of computers directly connected to it.
        direct_connections = defaultdict(set[str])

        for i, line in enumerate(file):
            c1, c2 = line.strip().split("-")
            direct_connections[c1].add(c2)
            direct_connections[c2].add(c1)

        connected_triplets = set()
        connected_chief_triplets = set()
        for i, first in enumerate(direct_connections.keys()):
            print(f"\rAnalyzing node #{i+1} [{first}]...", end="")

            if len(direct_connections[first]) >= 2:
                for second in direct_connections[first]:
                    if second > first:
                        for third in direct_connections[first]:
                            if third > second and second in direct_connections[third]:
                                connected_triplets.add( (first, second, third) )
                                if first[0] == 't' or second[0] == 't' or third[0] == 't':
                                    connected_chief_triplets.add( (first, second, third) )

        print(f"\nAnswer A: {len(connected_chief_triplets)}")

        lans = []
        for computer in direct_connections:
            lan_set = set()
            lan_set.add(computer)
            list_conns = list(direct_connections[computer])
            for i in range(len(list_conns) - 1):
                for j in range(len(list_conns)):
                    if list_conns[i] in direct_connections[list_conns[j]]:
                        lan_set.add(list_conns[j])
            lans.append(tuple(sorted(lan_set)))

        biggest_clique = ()
        biggest_clique_size = -1
        for lan in lans:
            if len(lan) == lans.count(lan):
                if len(lan) > biggest_clique_size:
                    biggest_clique = lan
                    biggest_clique_size = len(lan)

        # Note: biggest_clique is a set, so items are automatically sorted alphabetically.
        password = ",".join(biggest_clique) 
        print(f"Answer B: {password}")

if __name__ == "__main__":
    main()
