import unittest

from lc0831 import Solution

class Test_maskPII(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.maskPII("LeetCode@LeetCode.com"), "l*****e@leetcode.com")
        self.assertEqual(s.maskPII("AB@qq.com"), "a*****b@qq.com")
        self.assertEqual(s.maskPII("1(234)567-890"), "***-***-7890")

if __name__ == "__main__":
    unittest.main()
