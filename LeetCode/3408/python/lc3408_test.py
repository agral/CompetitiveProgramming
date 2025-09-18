import unittest

from lc3408 import TaskManager

class Test_TaskManager(unittest.TestCase):
    def test(self):
        t = TaskManager([
            [1, 101, 10], # userId=1, taskId=101, priority=10
            [2, 102, 20], # userId=2, taskId=102, priority=20
            [3, 103, 15], # userId=3, taskId=103, priority=15
        ])
        t.add(4, 104, 5)  # userId=4, taskId=104, priority=5
        t.edit(102, 8)    # edit taskId=102, new priority=8
        executed = t.execTop()
        self.assertEqual(executed, 3)
        t.rmv(101)
        t.add(5, 105, 15) # userId=5, taskId=105, priority=15
        executed = t.execTop()
        self.assertEqual(executed, 5)

if __name__ == "__main__":
    unittest.main()
