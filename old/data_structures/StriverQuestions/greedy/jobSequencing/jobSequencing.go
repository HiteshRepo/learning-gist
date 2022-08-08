package jobSequencing

func JobSequencing(jobs []Job) []int {
	sortJobsByDecreasingProfit(jobs)
	maxDeadline := getMaxDeadline(jobs)
	jobSequence := make([]*Job, maxDeadline)

	for _,job := range jobs {
		currDeadline := job.Deadline
		for currDeadline > 0 {
			if jobSequence[currDeadline-1] == nil {
				j := job
				jobSequence[currDeadline-1] = &j
				break
			}
			currDeadline -= 1
		}
	}

	count := 0
	totalProfit := 0
	for _,job := range jobSequence {
		if job != nil {
			count += 1
			totalProfit += job.Profit
		}
	}

	return []int{count, totalProfit}
}

func sortJobsByDecreasingProfit(jobs []Job) {
	for i:=0; i<len(jobs); i++ {
		maxProfitJob := i
		for j:=i+1; j<len(jobs); j++ {
			if jobs[maxProfitJob].Profit < jobs[j].Profit {
				maxProfitJob = j
			}
		}
		tmp := jobs[i]
		jobs[i] = jobs[maxProfitJob]
		jobs[maxProfitJob] = tmp
	}
}

func getMaxDeadline(jobs []Job) int {
	maxDeadline := jobs[0].Deadline
	for _,job := range jobs {
		if maxDeadline < job.Deadline {
			maxDeadline = job.Deadline
		}
	}
	return maxDeadline
}
