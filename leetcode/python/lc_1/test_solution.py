from unittest import TestCase

from two_sum import Solution

class TestSolution(TestCase):
    def test_twoSum(self):
        sol = Solution()
        self.assertEqual([0, 2], sol.twoSum([2, 7, 11, 15], 13))
        self.assertEqual([0, 3], sol.twoSum([2, 7, 11, 15], 17))
        self.assertEqual([1, 2], sol.twoSum([2, 7, 11, 15], 18))
        self.assertEqual([3, 4], sol.twoSum([2, 7, 11, 15, 16], 31))
        self.assertEqual([], sol.twoSum([2, 7, 11, 15, 16], 30))
        self.assertEqual([0, 1], sol.twoSum([2, 2, 7, 11, 15, 16], 4))
        self.assertEqual([0, 3], sol.twoSum([0, 4, 3, 0], 0))
