import unittest
from array_reverse import reverse_array, reverse_array_recursive, reverse_array_slicing


class TestArrayReverse(unittest.TestCase):
    def test_array_reverse(self):
        self.assertEqual(reverse_array([1, 2, 3, 4, 5]), [5, 4, 3, 2, 1])
        self.assertEqual(reverse_array_recursive(
            [4, 5, 1, 2], 0, 3), [2, 1, 5, 4])
        self.assertEqual(reverse_array_slicing(
            [1, 2, 3, 4, 5]), [5, 4, 3, 2, 1])
