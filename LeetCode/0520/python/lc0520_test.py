import unittest

from lc0520 import Solution

class Test_detectCapitalUse(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertEqual(s.detectCapitalUse("USA"), True) # Example 01
        self.assertEqual(s.detectCapitalUse("FlaG"), False) # Example 02

        self.assertEqual(s.detectCapitalUse("ABCDE"), True) # Case 01, "all letters are capitals"
        self.assertEqual(s.detectCapitalUse("XXXXXXX"), True)
        self.assertEqual(s.detectCapitalUse("agral"), True) # Case 02, "no capitals whatsoever"
        self.assertEqual(s.detectCapitalUse("boringcase"), True)
        self.assertEqual(s.detectCapitalUse("Charles"), True) # Case 03, "only the first letter is capital"
        self.assertEqual(s.detectCapitalUse("Korea"), True) # Case 03, "only the first letter is capital"
        self.assertEqual(s.detectCapitalUse("PapuaNewGuinea"), False) # none of cases 01-03 hold
        self.assertEqual(s.detectCapitalUse("detectCapitalUse"), False) # none of cases 01-03 hold
        self.assertEqual(s.detectCapitalUse("DETECTcapitaluse"), False) # none of cases 01-03 hold

        self.assertEqual(s.detectCapitalUse("a"), True) # Case 02, "no capitals whatsoever"
        self.assertEqual(s.detectCapitalUse("A"), True) # Case 01 and Case 03 both hold"

if __name__ == "__main__":
    unittest.main()
