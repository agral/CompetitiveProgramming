from typing import List

# Runtime complexity: O(n)
# Auxiliary space complexity: O(2) == O(1).
# Subjective level: easy+
# Solved on: 2026-02-09
class Solution:
    def countStudents(self, students: List[int], sandwiches: List[int]) -> int:
        sandwich_preference = [0, 0]
        for s in students:
            sandwich_preference[s] += 1

        for i in range(len(sandwiches)):
            if sandwich_preference[sandwiches[i]] == 0:
                # This sandwich won't be taken by the remaining students, who all
                # strongly prefer a sandwich of other type.
                # Note: the sandwich of other type may be present in the sandwich queue,
                # but no one of the students remaining in the queue will ever take
                # the unpreferred front one.
                return len(sandwiches) - i
            # otherwise there's still *someone* in the students' queue who will take
            # this sandwich.
            sandwich_preference[sandwiches[i]] -= 1

        # if reached here, all the sandwiches have been taken by the students.
        return 0
