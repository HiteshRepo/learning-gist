import sys
class Solution:
    #Function to find the days of buying and selling stock for max profit.
	def stockBuySell(self, A, n):
		buyingDay = 0
		sellingDay = 0
		profit = 0
		days = []
		
		i = 1
		while i<n:
		    if A[i] >= A[i-1]:
		        sellingDay += 1
		    else:
		        if buyingDay != sellingDay:
		            profit = profit + A[sellingDay] - A[buyingDay]
		            days.append([buyingDay, sellingDay])
		        buyingDay = sellingDay = i
		    i = i+1
		if buyingDay != sellingDay:
		    profit = profit + A[sellingDay] - A[buyingDay]
		    days.append([buyingDay, sellingDay])
		    
		return days