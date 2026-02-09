import unittest

from lc0232 import MyQueue

class Test_MyQueue(unittest.TestCase):
    def test(self):
        q = MyQueue()
        q.push(1) # queue is: [1]
        self.assertEqual(q.peek(), 1)
        q.push(2) # queue is: [1, 2]
        self.assertEqual(q.peek(), 1)
        popped = q.pop() 
        self.assertEqual(popped, 1)

        self.assertFalse(q.empty()) # after popping value of 1 above, queue is: [2]
        self.assertEqual(q.pop(), 2)
        self.assertTrue(q.empty())

if __name__ == "__main__":
    unittest.main()
