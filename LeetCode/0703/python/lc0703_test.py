import unittest

from lc0703 import KthLargest

class Test_KthLargest(unittest.TestCase):
    def test1(self):
        kl = KthLargest(3, [4, 5, 8, 2]) # keep: 4, 5, 8
        self.assertEqual(kl.add(3), 4) # keep: 4, 5, 8
        self.assertEqual(kl.add(5), 5) # keep: 5, 5, 8
        self.assertEqual(kl.add(10), 5) # keep: 5, 8, 10
        self.assertEqual(kl.add(9), 8) # keep: 8, 9, 10
        self.assertEqual(kl.add(4), 8) # keep: 8, 9, 10

    def test2(self):
        kl = KthLargest(4, [7, 7, 7, 7, 8, 3]) # keep: 7, 7, 7, 8
        self.assertEqual(kl.add(2), 7) # keep: 7, 7, 7, 8
        self.assertEqual(kl.add(10), 7) # keep: 7, 7, 8, 10
        self.assertEqual(kl.add(9), 7) # keep: 7, 8, 9, 10
        self.assertEqual(kl.add(9), 8) # keep: 8, 9, 9, 10


if __name__ == "__main__":
    unittest.main()
