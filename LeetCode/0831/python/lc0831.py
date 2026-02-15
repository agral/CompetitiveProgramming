# Runtime complexity: O(1)
# Auxiliary space complexity: O(1)
# Subjective level: medium, but it's only due to some bullshit rules regarding numbers' masking.
# Solved on: 2026-02-15
class Solution:
    def maskPII(self, s: str) -> str:
        at_pos = s.find('@')
        if at_pos != -1: # so s is an e-mail address:
            s = s.lower()
            return s[0] + "*****" + s[at_pos-1:]

        # the rules for the phone number matching are something else entirely:
        nums = "".join(d for d in s if d.isdigit())
        if len(nums) == 10:
            return "***-***-" + nums[-4:] # case of country code having 0 digits - handled separately

        num_stars = len(nums) - 10
        return "+" + "*" * num_stars + "-***-***-" + nums[-4:]
