import unittest

from lc3484 import Spreadsheet

class Test_spreadsheet(unittest.TestCase):
    def test(self):
        s = Spreadsheet(3)
        self.assertEqual(s.getValue("=5+7"), 12)

        s.setCell("A1", 10)
        self.assertEqual(s.getValue("=A1+6"), 16)
        self.assertEqual(s.getValue("=6+A1"), 16)

        s.setCell("B2", 15)
        self.assertEqual(s.getValue("=A1+B2"), 25)

        s.resetCell("A1") # same as setCell("A1", 0). A1 is now 0.
        self.assertEqual(s.getValue("=A1+B2"), 15)


        s = Spreadsheet(458)
        self.assertEqual(s.getValue("=O126+10272"), 10272)

        s = Spreadsheet(24)
        s.setCell("B24", 66688)
        s.resetCell("O15")
        self.assertEqual(s.getValue("=B24+O15"), 66688)

if __name__ == "__main__":
    unittest.main()
