import unittest
from array_minimize_heights import minimize_heights


class TestArrayMinimizeHeights(unittest.TestCase):
    def test_array_minimize_heights1(self):
        self.assertEqual(minimize_heights(
            [1, 5, 8, 10], 2), 5)

    def test_array_minimize_heights2(self):
        self.assertEqual(minimize_heights(
            [3, 9, 12, 16, 20], 3), 11)
