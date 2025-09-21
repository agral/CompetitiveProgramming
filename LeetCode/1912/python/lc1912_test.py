import unittest

from lc1912 import MovieRentingSystem

class Test_mrs(unittest.TestCase):
    def test(self):
        m = MovieRentingSystem(3,
            [[0, 1, 5], [0, 2, 6], [0, 3, 7], [1, 1, 4], [1, 2, 7], [2, 1, 5]])
        self.assertEqual(m.search(1), [1, 0, 2])
        m.rent(0, 1) # rent movie 1 from shop 0.
        m.rent(1, 2) # rent movie 2 from shop 1.
        self.assertEqual(m.report(), [[0, 1], [1, 2]])
        m.drop(1, 2) # return movie 2 to shop 1.
        self.assertEqual(m.search(2), [0, 1])

if __name__ == "__main__":
    unittest.main()
