import unittest

from lc2515 import Solution

class Test_closestTarget(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.closestTarget(["hello", "this", "is", "a", "test", "hello"], "hello", 1), 1)
        self.assertEqual(s.closestTarget(["a", "b", "c"], "c", 0), 1)
        self.assertEqual(s.closestTarget(["a", "b", "c"], "g", 0), -1)
        #self.assertEqual(s.closestTarget([], "", 0), 0)

if __name__ == "__main__":
    unittest.main()
