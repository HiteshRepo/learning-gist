class Solution:
    
    #Function to reverse words in a given string.
    def reverseWords(self,S):
        words = S.split('.')
        i = 0
        j = len(words)-1
        while i<j:
            words[i],words[j] = words[j],words[i]
            i=i+1
            j=j-1
        return '.'.join(words)