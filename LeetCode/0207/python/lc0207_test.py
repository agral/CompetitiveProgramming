import unittest

from lc0207 import Solution

class Test_canFinish(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.canFinish(2, [[1,0]]), True)
        self.assertEqual(s.canFinish(2, [[1,0], [0,1]]), False)

if __name__ == "__main__":
    unittest.main()
