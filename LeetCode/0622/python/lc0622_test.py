import unittest

from lc0622 import MyCircularQueue

class Test_MyCircularQueue(unittest.TestCase):
    def test(self):
        q = MyCircularQueue(3)                # Queue is: -, -, -
        self.assertEqual(q.enQueue(1), True)  # Queue is: 1, -, -
        self.assertEqual(q.enQueue(2), True)  # Queue is: 1, 2, -
        self.assertEqual(q.enQueue(3), True)  # Queue is: 1, 2, 3
        self.assertEqual(q.enQueue(4), False) # Queue is: still 1, 2, 3 (this op simply fails)
        self.assertEqual(q.Rear(), 3)         # Queue is: 1, 2, 3
        self.assertEqual(q.isFull(), True)
        self.assertEqual(q.deQueue(), True)   # Queue is: -, 2, 3
        self.assertEqual(q.enQueue(4), True)  # Queue is: 4, 2, 3 (front: 1, rear: 0)
        self.assertEqual(q.Rear(), 4)         # Queue is: 4, 2, 3 (front: 1, rear: 0), so 4 should be returned.

if __name__ == "__main__":
    unittest.main()
