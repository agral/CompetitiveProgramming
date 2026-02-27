import unittest

from lc0056 import Solution

class Test_merge(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.merge([[1, 3], [2, 6], [8, 10], [15, 18]]), [[1, 6], [8, 10], [15, 18]])
        self.assertEqual(s.merge([[1, 4], [4, 5]]), [[1, 5]])
        self.assertEqual(s.merge([[4, 7], [1, 4]]), [[1, 7]])

if __name__ == "__main__":
    unittest.main()
