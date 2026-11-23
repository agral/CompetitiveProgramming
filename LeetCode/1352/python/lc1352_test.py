import unittest

from lc1352 import ProductOfNumbers

class Test_productOfNumbers(unittest.TestCase):
    def test(self):
        p = ProductOfNumbers()
        p.add(3)
        p.add(0)
        p.add(2)
        p.add(5)
        p.add(4)

        self.assertEqual(p.getProduct(2), 4*5)
        self.assertEqual(p.getProduct(3), 4*5*2)
        self.assertEqual(p.getProduct(4), 4*5*2*0)

        p.add(8)
        self.assertEqual(p.getProduct(2), 8*4)

if __name__ == "__main__":
    unittest.main()
