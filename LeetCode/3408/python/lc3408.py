from sortedcontainers import SortedDict, SortedSet
from typing import List

class Task:
    def __init__(self, userId: int, taskId: int, priority: int) -> None:
        self.userId = userId
        self.taskId = taskId
        self.priority = priority

    def __lt__(self, other):
        if self.priority == other.priority:
            return self.taskId > other.taskId # yes, greater-than in a less-than implementation.
        return self.priority > other.priority # this is since I want automatic sorting to put
                                              # high-prio / high-id items first.

# Runtime complexity: O(n * logn).
# - __init__(): O(n * logn).
# - add(), edit(), rmv(), execTop(): O(logn).
# Auxiliary space complexity: O(len(tasks)); that's O(n).
# Subjective level: medium/hard.
class TaskManager:
    def __init__(self, tasks: List[List[int]]):
        self.idToTask = SortedDict() # maps taskId to Task instance.
        self.tasks = SortedSet() # a set of tasks, auto-sorted by high-priority, high-id.
        for t in tasks:
            self.add(t[0], t[1], t[2])

    def add(self, userId: int, taskId: int, priority: int) -> None:
        """Adds a new task. It is guaranteed that this function will be called with
           taskId that does *not* exist in the system."""
        newTask = Task(userId, taskId, priority)
        self.idToTask[taskId] = newTask
        self.tasks.add(newTask)

    def edit(self, taskId: int, newPriority: int) -> None:
        """Sets the provided newPriority to a task identified by a taskId.
           It is guaranteed that this function will be called with taskId that exists in the system. """
        task = self.idToTask[taskId]
        self.tasks.remove(task)
        updated = Task(task.userId, taskId, newPriority)
        self.tasks.add(updated)
        self.idToTask[taskId] = updated

    def rmv(self, taskId: int) -> None:
        """Removes the task identified by taskId from the system. It is guaranteed that
           this function will only be called with taskId existing in the system."""
        task = self.idToTask[taskId]
        self.tasks.remove(task)
        del self.idToTask[taskId]

    def execTop(self) -> int:
        if not self.tasks:
            return -1
        task = self.tasks.pop(0)
        ans = task.userId
        del self.idToTask[task.taskId]
        return ans
