import unittest

from lc0242 import Solution

class Test_isAnagram(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.isAnagram("anagram", "nagaram"), True)
        self.assertEqual(s.isAnagram("rat", "car"), False)

if __name__ == "__main__":
    unittest.main()
