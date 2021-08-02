import unittest
from array_union_intersecion import array_union_intersecion


class TestArrayUnionIntersection(unittest.TestCase):
    def test_array_union_intersection(self):
        self.assertEqual(array_union_intersecion(
            [1, 2, 3, 4, 5], [1, 2, 3, 4]), [5, 4])
