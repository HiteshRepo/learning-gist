import unittest
from array_sort_012 import sort012, sortCount012


class TestArraySort012(unittest.TestCase):
    def test_array_sort012(self):
        self.assertEqual(sort012([1, 0, 2, 0, 1, 2, 0, 0, 0, 2, 1]), [
                         0, 0, 0, 0, 0, 1, 1, 1, 2, 2, 2])

    def test_array_sortCount012(self):
        self.assertEqual(sortCount012([1, 0, 2, 0, 1, 2, 0, 0, 0, 2, 1]), [
                         0, 0, 0, 0, 0, 1, 1, 1, 2, 2, 2])
