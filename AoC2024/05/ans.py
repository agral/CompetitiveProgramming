from typing import List

DO_EXAMPLE = not True

graph = {}
updates = []
num_outgoing = {}
all_reachable = {}

seen_nodes = set()

def is_correct(update: List[int]) -> bool:
    for i in range(len(update)-1):
        for j in range(i+1, len(update)):
            if not update[i] in graph.keys():
                return False
            if not update[j] in graph[update[i]]:
                return False
    return True

def bubble_sort(alist):
    for i in range(len(alist) - 1, 0, -1):
        for j in range(i):
            # if there's an edge from j+1 to j, swap:
            if alist[j] in graph[alist[j+1]]:
                tmp = alist[j+1]
                alist[j+1] = alist[j]
                alist[j] = tmp
    return alist

def main():
    A = 0
    B = 0
    with open("example.txt" if DO_EXAMPLE else "input.txt", "r") as file:
        # Read input data:
        for line in file:
            if "|" in line:
                # Construct a directed graph of pages' priority:
                (src, dest) = [int(s) for s in line.split("|")]
                seen_nodes.add(src)
                seen_nodes.add(dest)
                if src in graph:
                    graph[src].append(dest)
                else:
                    graph[src] = [dest]


            if "," in line:
                update = [int(s) for s in line.split(",")]
                updates.append(update)

        for n in seen_nodes:
            if n not in graph.keys():
                graph[n] = []

        # Note: there's a cycle in the input data. There is no cycle in the example data.
        # Having a cycle in the input data makes the entire data invalid.
        # But OK, let's just not consider the graph in its entirety - only the immediate edges
        # and see if it works. What a confusing description of the problem!

        # Used the following code to determine if a cycle is present:
        # Make another graph that represents *all* nodes reachable from given node.
        # for given nodes a, b, c; if a -> b and b -> c, the edge a -> c may not be explicitly given
        # in the input data. 
        for key in graph.keys():
            all_reachable[key] = set()
            for node in graph[key]:
                all_reachable[key].add(node)
        is_updated = True # for the first run.
        while (is_updated):
            is_updated = False
            for a in all_reachable.keys():
                for b in all_reachable.keys():
                    for c in all_reachable.keys():
                        if (b in all_reachable[a] and c in all_reachable[b] and not c in all_reachable[a]):
                            all_reachable[a].add(c)
                            is_updated = True
        for node in all_reachable:
            if len(all_reachable[node]) == len(seen_nodes):
                # all nodes, including this one, is reachable from here. There's a cycle. Go home.
                print(f"Cycle found with node v={node}.")
                break
        # end of cycle checking code.

        for update in updates:
            if is_correct(update):
                A += update[len(update) // 2]
            else:
                # that didn't work for the reason described in the previous note regarding graph cycles.
                # given that there's a cycle, *all* vertices have the same number of reachable vertices,
                # which equals the total count of graph vertices.
                #corrected = sorted(update, key=lambda x: num_outgoing[x] if x in num_outgoing.keys() else -1,
                #                   reverse=True)
                corrected = bubble_sort(update[:])
                B += corrected[len(update) // 2]

        print(f"Answer A: {A}")
        print(f"Answer B: {B}")

if __name__ == "__main__":
    main()
