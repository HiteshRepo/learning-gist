openList = ['[', '{', '(', '<']
closeList = [']', '}', ')', '>']
invalidMap = {')':3, ']':57, '}':1197, '>':25137}
scoreMap = {')':1, ']': 2, '}': 3, '>': 4}

class Solution:
    incompleteMap = dict()

    def checkBalanced(self, line):
        st = []

        for ch in line:
            if ch in openList:
                st.append(ch)
            elif ch in closeList:
                pos = closeList.index(ch)
                if len(st) > 0 and openList[pos] == st[-1]:
                    st.pop()
                else:
                    return False, ch, []
        if len(st) > 0:
            return False, '', st
        return True, '', []

    def syntaxChecker(self, lines):
        syntaxErrScore = 0
        for line in lines:
            balanced, ch, lst = self.checkBalanced(line)
            if not balanced and ch != '':
                syntaxErrScore += invalidMap[ch]
            if not balanced and ch == '':
                self.incompleteMap[line] = lst
        return syntaxErrScore


    def calculateScore(self, closeStack):
        score = 0
        while len(closeStack) > 0 :
            ch = closeStack.pop()
            score *= 5
            score += scoreMap[closeList[openList.index(ch)]]
        return score

    def findMiddleScore(self, scores):
        pass

    def autocompleteScore(self, lines):
        scoreTracker = []
        self.syntaxChecker(lines)
        for key, value in self.incompleteMap.items():
            score = self.calculateScore(value)
            scoreTracker.append(score)
        scoreTracker.sort()
        middle = float(len(scoreTracker))/2
        if middle % 2 != 0:
            return scoreTracker[int(middle - .5)]
        else:
            return (scoreTracker[int(middle)], scoreTracker[int(middle-1)])



if __name__ == '__main__':
    f = open("./dist/input.txt", "r")
#     f = open("./dist/test.txt", "r")
    input = f.read()
    f.close()
    lines = input.strip().split('\n')

    obj = Solution()
#     print('syntax error score: ', obj.syntaxChecker(lines))
    print('autocomplete score: ', obj.autocompleteScore(lines))