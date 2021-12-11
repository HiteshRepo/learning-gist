
uniqueDigitMap = {2: 1, 4: 4, 3: 7, 7:8}
uniqueDigitSignalMap = {}

class Solution:
    def isPresent(self, digit):
        if len(digit) in uniqueDigitMap:
            uniqueDigitSignalMap[uniqueDigitMap[len(digit)]] = digit.split()
            return True
        return False

    def count1478(self, finalSignals, i):
        count = 0
        while i < len(finalSignals):
            uniqueDigits = finalSignals[i]
            for digit in uniqueDigits:
                if self.isPresent(digit):
                    count += 1
            i += 2
        return count

    def deduce(self, signals):
        uniqueDigitSignalMap = {}
        digits = signals[0]
        for digit in digits:
            self.isPresent(digit)

        print(uniqueDigitSignalMap)




if __name__ == '__main__':
#     f = open("./dist/input.txt", "r")
    f = open("./dist/test.txt", "r")
    input = f.read()
    f.close()
    signals = input.strip().split('\n')

    finalSignals = []
    for signal in signals:
        parts = signal.strip().split('|')
        finalSignals.append(parts[0].strip().split())
        finalSignals.append(parts[1].strip().split())

    obj = Solution()
    obj.deduce(finalSignals)
#     print("number of 1,4,7,8 in signals: ", obj.count1478(finalSignals,1))