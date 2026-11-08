import unittest

from lc1267 import Solution

class Test_countServers(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.countServers([[1, 0], [0, 1]]), 0)
        self.assertEqual(s.countServers([[1, 0], [1, 1]]), 3)
        self.assertEqual(s.countServers([
            [1, 1, 0, 0],
            [0, 0, 1, 0],
            [0, 0, 1, 0],
            [0, 0, 0, 1]]), 4)

if __name__ == "__main__":
    unittest.main()
