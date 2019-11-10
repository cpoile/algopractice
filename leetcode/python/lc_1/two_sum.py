from typing import List


def binSearch(nums: List[int], target: int) -> int:
    lo, hi = 0, len(nums) - 1
    while lo <= hi:
        mid = lo + (hi - lo) // 2
        if nums[mid] < target:
            lo = mid + 1
        elif nums[mid] > target:
            hi = mid - 1
        else:
            return mid
    return -1


class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        vmap = {}
        for a, ai in enumerate(nums):
            comp = target - ai
            if comp in vmap:
                return [vmap[comp], a]
            vmap[ai] = a
        return []
