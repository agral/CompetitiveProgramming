class Solution:
    def canChange(self, start: str, target: str) -> bool:
        SZ = len(start)
        start_idx = 0
        target_idx = 0
        while start_idx <= SZ and target_idx <= SZ:
            while start_idx < SZ and start[start_idx] == "_":
                start_idx += 1
            while target_idx < SZ and target[target_idx] == "_":
                target_idx += 1
            if start_idx == SZ or target_idx == SZ:
                return start_idx == SZ and target_idx == SZ
            if start[start_idx] != target[target_idx]:
                return False
            if start[start_idx] == "L" and start_idx < target_idx:
                return False
            if start[start_idx] == "R" and start_idx > target_idx:
                return False
            start_idx += 1
            target_idx += 1
        return True


def main():
    s = Solution()
    assert(s.canChange("_L__R__R_", "L______RR") == True)
    assert(s.canChange("R_L_", "__LR") == False)
    assert(s.canChange("_R", "R_") == False)

if __name__ == "__main__":
    main()
