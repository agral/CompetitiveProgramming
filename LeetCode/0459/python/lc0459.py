from typing import List

# Runtime complexity: O(n)
# Auxiliary space complexity: O(2*n) == O(n)
# Subjective level: easy++. Have to know the trick.
#   In addition, I'd consider one-letter strings as truthy - e.g. `a` can be constructed
#   using a substring (`a`) and repeating it 1 times. But OK, I can work around it.
# Solved on: 2026-02-16
class Solution:
    def repeatedSubstringPattern(self, s: str) -> bool:
        L = len(s)
        if L == 1: # An edge case - this problem considers one-letter strings as falsey.
            return False

        # If the input string `s` has a repeating pattern, then `s` repeated twice
        # contains the original `s` somewhere inside, in addition to the trivially found
        # `s`es at the beginning/ending of the combined repeated string.

        # combo = (s+s)[1:2*L-1] # note: should be 2*L-2, but Python requires end+1, so it's 2*L-1 here.
        # return combo.find(s) != -1
        return (s+s)[1:2*L-1].find(s) != -1
