

class Solution:

    #Function to find the minimum number of swaps required to sort the array.
    def minSwaps(self, nums):
        numsWithPos = {}
        visited = {}

        for i in range(len(nums)):
            numsWithPos[nums[i]] = i
            visited[nums[i]] = False

        sortedKeys = sorted(numsWithPos.keys())

        swaps = 0

        for i in range(len(sortedKeys)):
            if visited[sortedKeys[i]] or i == numsWithPos[sortedKeys[i]]:
                continue

            cycleLen = 0
            j = i

            while(True):
                visited[sortedKeys[j]] = True
                cycleLen = cycleLen + 1
                j = numsWithPos[sortedKeys[j]]
                if visited[sortedKeys[j]]:
                    break

            swaps = swaps + cycleLen-1


        return swaps


#{
#  Driver Code Starts


if __name__ == '__main__':
    T=int(input())
    for i in range(T):
        n = int(input())
        nums = list(map(int, input().split()))
        obj = Solution()
        ans = obj.minSwaps(nums)
        print(ans)

# } Driver Code Ends