import unittest

from lc3508 import Router

class Test_Router(unittest.TestCase):
    def test(self):
        r = Router(3)
        self.assertTrue(r.addPacket(1, 4, 90))
        self.assertTrue(r.addPacket(2, 5, 90))
        self.assertFalse(r.addPacket(1, 4, 90)) # this is a duplicate of a first packet.
        self.assertTrue(r.addPacket(3, 5, 95))

        self.assertTrue(r.addPacket(4, 5, 105))
        # This operation exceeds the Router's memory limit of 3,
        # so the [1, 4, 90] packet has to get removed.

        forwarded = r.forwardPacket() # Has to return [2, 5, 90].
        self.assertEqual(forwarded[0], 2)
        self.assertEqual(forwarded[1], 5)
        self.assertEqual(forwarded[2], 90)

        self.assertTrue(r.addPacket(5, 2, 110))
        self.assertEqual(r.getCount(5, 100, 110), 1)

if __name__ == "__main__":
    unittest.main()
