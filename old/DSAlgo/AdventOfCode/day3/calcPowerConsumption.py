import copy

class Solution:
    def findOnesAndZeroes(self, diagnostics, length):
        ones = [0]*length
        zeroes = [0]*length

        for d in diagnostics:
            for i in range(length):
                bit = d[i]
                if bit == '1':
                    ones[i] += 1
                else:
                    zeroes[i] += 1
        return ones, zeroes

    def calcPowerConsumption(self, diagnostics):
        gammaRate = 0
        epsilonRate = 0
        length = len(diagnostics[0])

        ones, zeroes = self.findOnesAndZeroes(diagnostics, length)

        for i in range(length):
            if ones[i] > zeroes[i]:
                gammaRate += (1*(2**(length-i-1)))
                epsilonRate += 0
            else:
                gammaRate += 0
                epsilonRate += (1*(2**(length-i-1)))

        return gammaRate*epsilonRate

    def calcLifeSupportingRating(self, diagnostics):
        oxyRate = 0
        co2Rate = 0
        length = len(diagnostics[0])

        oxygenDiagnostics = copy.deepcopy(diagnostics)
        for i in range(length):
            max = '0'
            ones, zeroes = self.findOnesAndZeroes(oxygenDiagnostics, length)
            if ones[i] >= zeroes[i]:
                max = '1'

            if len(oxygenDiagnostics) > 1:
                oxygenDiagnostics = list(filter(lambda x: x[i] == max, oxygenDiagnostics))

            if len(oxygenDiagnostics) == 1:
                break


        co2Diagnostics = copy.deepcopy(diagnostics)
        for i in range(length):
            max = '1'
            ones, zeroes = self.findOnesAndZeroes(co2Diagnostics, length)
            if zeroes[i] <= ones[i]:
                max = '0'

            if len(co2Diagnostics) > 1:
                co2Diagnostics = list(filter(lambda x: x[i] == max, co2Diagnostics))

            if len(co2Diagnostics) == 1:
                break

        for i in range(length):
            oxyRate += int(oxygenDiagnostics[0][i]) * (2**(length-i-1))
            co2Rate += int(co2Diagnostics[0][i]) * (2**(length-i-1))

        return oxyRate * co2Rate


if __name__ == '__main__':
    f = open("./dist/input.txt", "r")
    input = f.read()
    f.close()
    diagnostics = input.split('\n')

#     Test case 1
#     diagnostics = ["00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"]

    obj = Solution()
    print("power consumption: ", obj.calcPowerConsumption(diagnostics))
    print("Life support rating: ", obj.calcLifeSupportingRating(diagnostics))