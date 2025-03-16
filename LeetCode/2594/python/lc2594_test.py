import unittest

from lc2594 import Solution

class Test_repairCars(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.repairCars([4, 2, 3, 1], 10), 16)
        self.assertEqual(s.repairCars([5, 1, 8], 6), 16)

if __name__ == "__main__":
    unittest.main()
