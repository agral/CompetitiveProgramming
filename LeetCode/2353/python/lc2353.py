from collections import defaultdict
from typing import List
from sortedcontainers import SortedSet

# Runtime complexity: O(n).
#  - __init__(): O(n)
#  - changeRating: O(log(n)) on average; O(n) in worst case.
#  - highestRated: O(log(n)).
# Auxiliary space complexity: O(n).
# Subjective level: medium.
class FoodRatings:
    def __init__(self, foods: List[str], cuisines: List[str], ratings: List[int]):
        self.foodToRating = dict()
        self.foodToCuisine = dict()
        self.cuisineToCombo = defaultdict(
            lambda: SortedSet(key=lambda k: (-k[0], k[1]))
        )
        for i in range(len(foods)):
            self.foodToRating[foods[i]] = ratings[i]
            self.foodToCuisine[foods[i]] = cuisines[i]
            self.cuisineToCombo[cuisines[i]].add((ratings[i], foods[i]))

    def changeRating(self, food: str, newRating: int) -> None:
        rating = self.foodToRating[food]
        cuisine = self.foodToCuisine[food]
        combo = self.cuisineToCombo[cuisine]
        combo.remove((rating, food))
        combo.add((newRating, food))
        self.foodToRating[food] = newRating

    def highestRated(self, cuisine: str) -> str:
        combo = self.cuisineToCombo[cuisine]
        return combo[0][1]
