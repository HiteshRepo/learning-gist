
class Solution:
    def permute(self, S, i, arr):
        if i == len(S):
            arr.append(''.join(S))
        else:
            for j in range(i, len(S)):
                S[i], S[j] = S[j],S[i]
                self.permute(S, i+1, arr)
                S[j], S[i] = S[i],S[j]
        
    def find_permutation(self, S):
        arr = []
        data = list(S)
        self.permute(data, 0, arr)
        return arr
        