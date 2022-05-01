import unittest
from array_cyclic_rotateonce import array_cyclic_rotateonce


class TestArrayCyclicRotateOnce(unittest.TestCase):
    def test_array_cyclic_rotate_once(self):
        self.assertEqual(array_cyclic_rotateonce(
            [1, 2, 3, 4, 5]), [5, 1, 2, 3, 4])
