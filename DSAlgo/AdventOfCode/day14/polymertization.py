class Solution:
    def findPairs(self, template):
        pairs = []
        for i in range(len(template)-1):
            pairs.append(template[i:i+2])
        return pairs

    def insertPolymers(self, pairs, rules):
        ans = ''
        for pair in pairs:
            if pair in rules:
                polymer = rules[pair]
                if ans == '':
                    ans = pair[0] + polymer + pair[1]
                else:
                    ans += polymer + pair[1]
        return ans

    def getMinMax(self, template):
        charMap = {}
        for c in template:
            if c in charMap:
                charMap[c] += 1
            else:
                charMap[c] = 1

        return charMap[max(charMap, key=charMap.get)], charMap[min(charMap, key=charMap.get)]

    def polymertize(self, template, rules, step):

        while step > 0:
            pairs = self.findPairs(template)
            template = self.insertPolymers(pairs, rules)
            step -= 1

        max, min = self.getMinMax(template)
        return max - min



if __name__ == '__main__':
    f = open("./dist/input.txt", "r")
#     f = open("./dist/test.txt", "r")
    input = f.read()
    f.close()
    lines = input.strip().split('\n')

    rules = {}
    template = lines[0]
    lines = lines[2:]

    for l in lines:
        pair = l.strip().split('->')
        rules[pair[0].strip()] = pair[1].strip()

    step = 40
    obj = Solution()
    print("most common element - least common elemnet = ", obj.polymertize(template, rules, step))

