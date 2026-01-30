import unittest

from lc0149 import Solution

class Test_maxPoints(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.maxPoints([[1,1],[2,2],[3,3]]), 3)
        self.assertEqual(s.maxPoints([[1,1],[3,2],[5,3],[4,1],[2,3],[1,4]]), 4)
        self.assertEqual(s.maxPoints([[3,3],[1,4],[1,1],[2,1],[2,2]]), 3)
        self.assertEqual(s.maxPoints([[0,0],[4,5],[7,8],[8,9],[5,6],[3,4],[1,1]]), 5)

if __name__ == "__main__":
    unittest.main()
