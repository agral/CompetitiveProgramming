import unittest

from lc1887 import Solution

class Test_reductionOperations(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.reductionOperations([5, 1, 3]), 3)
        self.assertEqual(s.reductionOperations([1, 1, 1]), 0)
        self.assertEqual(s.reductionOperations([1, 1, 2, 2, 3]), 4)
        self.assertEqual(s.reductionOperations([1, 1, 1, 1, 2]), 1)
        self.assertEqual(s.reductionOperations([1, 1, 1, 2, 2]), 2)
        self.assertEqual(s.reductionOperations([1, 2, 2, 2, 2]), 4)
        self.assertEqual(s.reductionOperations([1, 2, 3, 4, 5]), 10)
        self.assertEqual(s.reductionOperations([1, 2, 3, 4, 4]), 9)

if __name__ == "__main__":
    unittest.main()
