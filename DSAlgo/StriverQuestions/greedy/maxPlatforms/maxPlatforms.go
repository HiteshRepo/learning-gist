package maxPlatforms

func FindPlatformsOptimal(arr, dep []int, n int) int {
	ans := 1
	count := 1
	sort(arr)
	sort(dep)

	i:=1
	j:=0

	for i<n && j<n {
		if arr[i]<=dep[j] {
			count += 1
			i++
		} else {
			count -= 1
			j++
		}

		ans = max(ans,count)
	}


	return ans
}

func FindPlatforms(arr, dep []int, n int) int {
	ans := 1
	for i:=0; i<n; i++ {
		count := 1
		for j:=i+1; j<n; j++ {
			if (arr[i]>=arr[j] && arr[i]<=dep[j]) || (arr[j]>=arr[i] && arr[j]<=dep[i]) {
				count++
			}
		}
		ans = max(ans, count)
	}

	return ans
}

func max(m, n int) int {
	if m>n {
		return m
	}
	return n
}


func sort(m []int) {
	count := 0

	for count < len(m) {
		small := count
		for i:=count; i<len(m); i++ {
			if m[i] < m[small] {
				small = i
			}
		}
		tmp := m[count]
		m[count] = m[small]
		m[small] = tmp

		count += 1
	}
}