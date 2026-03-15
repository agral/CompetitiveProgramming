import unittest

from lc1622 import Fancy

class Test_Fancy(unittest.TestCase):
    def test(self):
        f = Fancy()
        f.append(2) # [2]
        f.addAll(3) # [2+3] == [5]
        f.append(7) # [5, 7]
        f.multAll(2) # [10, 14]
        self.assertEqual(f.getIndex(0), 10)
        f.addAll(3) # [10+3, 14+3] == [13, 17]
        f.append(10) # [13, 17, 10]
        f.multAll(2) # [26, 34, 20]
        self.assertEqual(f.getIndex(0), 26)
        self.assertEqual(f.getIndex(1), 34)
        self.assertEqual(f.getIndex(2), 20)

if __name__ == "__main__":
    unittest.main()
