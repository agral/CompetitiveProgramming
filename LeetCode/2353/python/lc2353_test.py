import unittest

from lc2353 import FoodRatings

class Test_foodRatings(unittest.TestCase):
    def test(self):
        fr = FoodRatings(
            ["kimchi", "miso",     "sushi",    "moussaka", "ramen",    "bulgogi"],
            ["korean", "japanese", "japanese", "greek",    "japanese", "korean"],
            [9,        12,         8,          15,         14,         7],
        )

        self.assertEqual(fr.highestRated("korean"), "kimchi")
        self.assertEqual(fr.highestRated("japanese"), "ramen")
        fr.changeRating("sushi", 16)
        self.assertEqual(fr.highestRated("japanese"), "sushi")
        fr.changeRating("ramen", 16)
        self.assertEqual(fr.highestRated("japanese"), "ramen")


if __name__ == "__main__":
    unittest.main()
