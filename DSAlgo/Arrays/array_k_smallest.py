class Solution:
    def maxHeapify(self, A, n, i):
        largest = i
        l = (2*i)+1
        r = (2*i)+2
        while l<n and A[l] > A[largest]:
            largest = l
        while r<n and A[r] > A[largest]:
            largest = r
        if largest != i:
            A[i], A[largest] = A[largest],A[i]
            self.maxHeapify(A, n, largest)
    
    def minHeapify(self, A, n, i):
        smallest = i
        l = (2*i)+1
        r = (2*i)+2
        while l<n and A[l] < A[smallest]:
            smallest = l
        while r<n and A[r] < A[smallest]:
            smallest = r
        if smallest != i:
            A[i], A[smallest] = A[smallest],A[i]
            self.minHeapify(A, n, smallest)
        
    def kthSmallest(self,arr, l, r, k):
        '''
        arr : given array
        l : starting index of the array i.e 0
        r : ending index of the array i.e size-1
        k : find kth smallest element and return using this function
        '''
        # n = len(arr)
        # i = int(n/2) - 1
        # while i>=0:
        #     self.maxHeapify(arr, n, i)
        #     i = i-1
        # i = n-1
        # cnt = 0
        # while i>=0:
        #     if cnt == k:
        #         return arr[0]
        #     arr[i], arr[0] = arr[0], arr[i]
        #     self.maxHeapify(arr, i, 0)
        #     i = i-1
        #     cnt = cnt+1
        n = len(arr)
        i = int(n/2) - 1
        while i>=0:
            self.minHeapify(arr, n, i)
            i = i-1
        i = n-1
        cnt = 1
        while i>=0:
            if cnt == k:
                return arr[0]
            arr[i], arr[0] = arr[0], arr[i]
            self.minHeapify(arr, i, 0)
            i = i-1
            cnt = cnt+1


# print(getKthMinMax('max', [7, 10, 4, 3, 20, 15], 2))
