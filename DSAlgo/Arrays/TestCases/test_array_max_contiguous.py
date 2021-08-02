import unittest
from array_max_contiguous import array_max_contiguous


class TestArrayMaxContiguous(unittest.TestCase):
    def test_array_max_contiguous_1(self):
        self.assertEqual(array_max_contiguous(
            [1, 2, -3, 4, 5]), 9)

    def test_array_max_contiguous_2(self):
        self.assertEqual(array_max_contiguous(
            [-2, 5, -1]), 5)

    def test_array_max_contiguous_3(self):
        self.assertEqual(array_max_contiguous(
            [-1, -2, -3, -4]), -1)
