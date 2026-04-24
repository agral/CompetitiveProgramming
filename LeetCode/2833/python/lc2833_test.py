import unittest

from lc2833 import Solution

class Test_furthestDistanceFromOrigin(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.furthestDistanceFromOrigin("L_RL__R"), 3)
        self.assertEqual(s.furthestDistanceFromOrigin("_R__LL_"), 5)
        self.assertEqual(s.furthestDistanceFromOrigin("_______"), 7)

if __name__ == "__main__":
    unittest.main()
