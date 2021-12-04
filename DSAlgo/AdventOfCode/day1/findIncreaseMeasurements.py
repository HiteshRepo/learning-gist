
class Solution:
    def findIncreaseCounts(self, measurements):
        prev = measurements[0]
        increaseCtr = 0

        for m in measurements:
            if m > prev:
                increaseCtr += 1
            prev = m

        return increaseCtr

    def findSlidingIncreaseCounts(self, measurements):
            prev = sum(measurements[0:3])
            increaseCtr = 0

            i = 0
            while i < len(measurements) - 2:
                curr = sum(measurements[i:i+3])
                if curr > prev:
                    increaseCtr += 1
                prev = curr
                i = i+1

            return increaseCtr


if __name__ == '__main__':
    f = open("./dist/input1.txt", "r")
    input = f.read()
    f.close()
    T = input.split('\n')
    measurements = list(map(int, T))

#     Test case 1
#     measurements = [199, 200, 208, 210, 200, 207, 240, 269, 260, 263]

    obj = Solution()
    print("number of times a depth measurement increases from the previous measurement: ", obj.findIncreaseCounts(measurements))

    f = open("./dist/input2.txt", "r")
    input = f.read()
    f.close()
    T = input.split('\n')
    measurements = list(map(int, T))

#     Test case 1
#     measurements = [199, 200, 208, 210, 200, 207, 240, 269, 260, 263]

    obj = Solution()
    print("number of times a depth measurement in sliding window increases from the previous sliding window measurement: ", obj.findSlidingIncreaseCounts(measurements))
