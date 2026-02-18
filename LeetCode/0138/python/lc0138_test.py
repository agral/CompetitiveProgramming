import unittest

from random import randint
from typing import List

from lc0138 import Solution, Node

def makeList(values: List[int]) -> Node | None:
    if not values:
        return
    # 1st pass: create all the nodes, store these in a list:
    head = Node(values[0])
    prev = head
    nodes = [head]
    for i in range(1, len(values)):
        curr = Node(values[i])
        nodes.append(curr)
        prev.next = curr
        prev = curr

    # 2nd pass: set the random pointers of each individual Node
    # to a random Node from the entire list - or a null pointer (None).
    L = len(nodes)
    curr = head
    while curr != None:
        random_idx = randint(0, L) # inclusive on both ends.
        if random_idx == L:
            curr.random = None
        else:
            curr.random = nodes[random_idx]
        curr = curr.next
    return head

def is_valid_copy(src: Node | None, cpy: Node | None) -> bool:
    if src is None:
        return cpy is None

    currSrc = src
    currCpy = cpy
    while currSrc is not None:
        # For each non-null source's node, check that the analogous copy's node exists
        # with both `val` and `random` values matching.
        if not currCpy:
            print(f"currCpy is null for currSrc={currSrc}")
            return False
        if currSrc.random is None and currCpy.random is not None:
            print(f"{currSrc} has null random, but {currCpy} has non-null random pointer")
            return False
        if currSrc.random is not None and currCpy.random is None:
            print(f"{currSrc} has non-null random, but {currCpy} has null random pointer")
            return False
        if currSrc.random is not None and currCpy.random is not None:
            # have to compare the random pointers
            if currSrc.random.val != currCpy.random.val:
                print(f"Randoms' values mismatch: {currSrc.random.val} != {currCpy.random.val}")
                return False
        # the only other possible case is both currSrc.random && currCpy.random being null.
        # Now for the values:
        if currSrc.val != currCpy.val:
            print(f"Values mismatch: {currSrc.val} != {currCpy.val}")
            return False

        # Advance both pointers:
        currSrc = currSrc.next
        currCpy = currCpy.next

    # currSrc is now null; a copy is valid if currCpy is also null
    # (for if a currCpy still continues, it is not a copy of src).
    return currCpy is None


def toString(head: Node | None) -> str:
    values = []
    curr = head
    while curr != None:
        values.append(f"(v:{curr.val}, r:{'None' if curr.random is None else curr.random.val})")
        curr = curr.next
    return f"[{' -> '.join(values)}]"


class Test_solve(unittest.TestCase):
    def test(self):
        s = Solution()
        ex1 = makeList([7, 13, 11, 10, 1])
        self.assertTrue(is_valid_copy(ex1, s.copyRandomList(ex1)))

        ex2 = makeList([1, 2])
        self.assertTrue(is_valid_copy(ex2, s.copyRandomList(ex2)))

        print("Valid copies are OK, now trying for invalid copies, expect errors logged below:")
        self.assertFalse(is_valid_copy(ex1, ex2))
        self.assertFalse(is_valid_copy(ex1, ex1.next))

if __name__ == "__main__":
    unittest.main()
