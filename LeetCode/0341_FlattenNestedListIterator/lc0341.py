class NestedIterator:
    def __init__(self, nestedList):
        # Convert nested list to a flat list.
        self.flat = []
        self.flatPos = 0;
        self.parse(nestedList)

    def parse(self, nestedList):
        for elem in nestedList:
            if (elem.isInteger()):
                self.flat.append(elem.getInteger())
            else:
                self.parse(elem.getList())

    def next(self):
        ans = self.flat[self.flatPos]
        self.flatPos += 1
        return ans

    def hasNext(self):
        return self.flatPos < len(self.flat)

