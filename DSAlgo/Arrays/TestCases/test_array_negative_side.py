import unittest
from array_negative_side import array_negative_side


class TestArrayNegativeSide(unittest.TestCase):
    def test_array_negative_side(self):
        self.assertEqual(array_negative_side(
            [1, -1, 2, 3, -9]), [-1, -9, 2, 3, 1])
