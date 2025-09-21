import collections
from sortedcontainers import SortedList
from typing import List

# Runtime complexity:
# - __init__(): O(nlogn)
# - search(), rent(), drop(), report(): O(logn)
# Auxiliary space complexity: O(n + len(self.rented))
# Subjective level: medium/hard
class MovieRentingSystem:
    def __init__(self, n: int, entries: List[List[int]]) -> None:
        self.rented = SortedList() # list of (price, shop, movie)
        self.unrented = collections.defaultdict(SortedList) # map[movie](price, shop)
        self.shopMovieToPrice = {} # map[(shop, movie)]price

        for shop, movie, price in entries:
            self.unrented[movie].add( (price, shop) )
            self.shopMovieToPrice[(shop, movie)] = price

    def search(self, movie: int) -> List[int]:
        ans = []
        for _, shop in self.unrented[movie][:5]:
            ans.append(shop)
        return ans

    def rent(self, shop: int, movie: int) -> None:
        price = self.shopMovieToPrice[(shop, movie)]
        self.rented.add( (price, shop, movie) )
        self.unrented[movie].remove( (price, shop) )

    def drop(self, shop: int, movie: int) -> None:
        price = self.shopMovieToPrice[(shop, movie)]
        self.rented.remove( (price, shop, movie) )
        self.unrented[movie].add( (price, shop) )

    def report(self) -> List[List[int]]:
        return [[shop, movie] for _, shop, movie in self.rented[:5]]
