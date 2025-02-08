import collections
from sortedcontainers import SortedSet

class NumberContainers:
  def __init__(self):
    self.numbersToIndices = collections.defaultdict(SortedSet)
    self.indicesToNumbers = {}

  def change(self, index: int, number: int) -> None:
    if index in self.indicesToNumbers:
      originalNumber = self.indicesToNumbers[index]
      self.numbersToIndices[originalNumber].remove(index)
      if len(self.numbersToIndices[originalNumber]) == 0:
        del self.numbersToIndices[originalNumber]
    self.indicesToNumbers[index] = number
    self.numbersToIndices[number].add(index)

  def find(self, number: int) -> int:
    if number in self.numbersToIndices:
      return self.numbersToIndices[number][0]
    return -1
