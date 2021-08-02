import unittest
from array_jump import jump


class TestArrayJump(unittest.TestCase):
    def test_scenario1(self):
        self.assertEqual(jump([1, 3, 5, 8, 9, 2, 6, 7, 6, 8, 9]), 3)

    def test_scenario2(self):
        self.assertEqual(jump([1, 4, 3, 2, 6, 7]), 2)

    def test_scenario3(self):
        self.assertEqual(jump([2, 3, 1, 1, 2, 4, 2, 0, 1, 1]), 4)

    def test_scenario4(self):
        self.assertEqual(jump([0, 1, 1, 1, 1]), -1)
