package nMeetings

func MaxMeetings(start, end []int, n int) int {
	ans := make([]int, 0)
	sort(end, start)
	limit := end[0]
	ans = append(ans, start[0])

	for i:=1; i<len(end); i++ {
		if start[i] > limit {
			ans = append(ans, start[i])
			limit = end[i]
		}
	}

	return len(ans)
}


func sort(m ,n []int) {
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

		tmp = n[count]
		n[count] = n[small]
		n[small] = tmp

		count += 1
	}
}
