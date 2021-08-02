import unittest
from array_min_max import getMinMax, getMinMaxRecursive


class TestArrayMinMax(unittest.TestCase):
    def test_array_min_max(self):
        self.assertEqual(getMinMax(
            [1000, 11, 445, 1, 330, 3000]), (1, 3000))
        self.assertEqual(getMinMaxRecursive(
            0, 5, [1000, 11, 445, 1, 330, 3000]), (1, 3000))
