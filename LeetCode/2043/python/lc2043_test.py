import unittest

from lc2043 import Bank

class Test_bank(unittest.TestCase):
    def test(self):
        b = Bank([10, 100, 20, 50, 30])
        self.assertEqual(b.accounts[1], 10)
        self.assertEqual(b.accounts[5], 30)

        isSuccessful = b.withdraw(3, 10)
        self.assertEqual(isSuccessful, True)

        isSuccessful = b.transfer(5, 1, 20)
        self.assertEqual(isSuccessful, True)

        isSuccessful = b.transfer(3, 4, 15)
        self.assertEqual(isSuccessful, False)

        isSuccessful = b.withdraw(10, 50)
        self.assertEqual(isSuccessful, False)

if __name__ == "__main__":
    unittest.main()
