from typing import List
from collections import defaultdict

# Runtime complexity: amortized O(1) (for each query).
# Auxiliary space complexity: O(1).
# Subjective level: medium
class Spreadsheet:
    def __init__(self, rows: int) -> None:
        self.data = defaultdict(int)

    def setCell(self, cell: str, value: int) -> None:
        self.data[cell] = value

    def resetCell(self, cell: str) -> None:
        self.setCell(cell, 0)

    def getValue(self, formula: str) -> int:
        """Evaluates a formula in the form "=X+Y", where both X and Y are either
        cell references OR non-negative integers. Returns the computed sum."""
        lhs, rhs = formula[1:].split("+")
        return self._evaluate(lhs) + self._evaluate(rhs)

    def _evaluate(self, cell: str) -> int:
        if cell.isdigit():
            return int(cell)
        return self.data[cell]
