from typing import List

# Subjective level: medium; also too tedious and teaches you nothing. A worthless one, IMO.
# Solved on: 2026-04-07
class Robot:
    DIRS = ["North", "East", "South", "West"]

    # Runtime complexity: O(2*(height+width))
    # Auxiliary space complexity: O(2*(height+width))
    def __init__(self, width: int, height: int):
        self.IsFirstMove = True
        self.theta = 0

        # This is the circle around the grid's perimeter the robot will be following.
        # For fucks' sake, they want (x, y) and I'm used to using (y, x).
        #self.circle = (
        #    [((0, 0), self.DIRS[2])] +
        #    [((0, x), self.DIRS[1]) for x in range(1, width)] +
        #    [((y, width-1), self.DIRS[0]) for y in range(1, height)] +
        #    [((height-1, x), self.DIRS[3]) for x in range(width-2, -1, -1)] +
        #    [((y, 0), self.DIRS[2]) for y in range(height-2, 0, -1)]
        #)
        self.circle = (
            [((0, 0), self.DIRS[2])] +
            [((x, 0), self.DIRS[1]) for x in range(1, width)] +
            [((width-1, y), self.DIRS[0]) for y in range(1, height)] +
            [((x, height-1), self.DIRS[3]) for x in range(width-2, -1, -1)] +
            [((0, y), self.DIRS[2]) for y in range(height-2, 0, -1)]
        )

        self.circumference = len(self.circle)

    # Runtime complexity: O(1)
    # Auxiliary space complexity: O(1)
    def step(self, num: int) -> None:
        self.IsFirstMove = False
        self.theta += num
        self.theta %= self.circumference

    # Runtime complexity: O(1)
    # Auxiliary space complexity: O(1)
    def getPos(self) -> List[int]: # actually Tuple[int, int], but LC demands List[int]. That's OK.
        return self.circle[self.theta][0]

    # Runtime complexity: O(1)
    # Auxiliary space complexity: O(1)
    def getDir(self) -> str:
        if self.IsFirstMove:
            return self.DIRS[1] # before making its first move, the robot faces east.
        return self.circle[self.theta][1]
