package main

type Pair struct {
	timestamp int
	value     string
}

type TimeMap struct {
	store map[string][]Pair
}

func Constructor() TimeMap {
	return TimeMap{store: map[string][]Pair{}}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	if _, alreadyExist := this.store[key]; !alreadyExist {
		this.store[key] = []Pair{}
	}
	this.store[key] = append(this.store[key], Pair{timestamp: timestamp, value: value})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	localStore, exists := this.store[key]

	if !exists || len(localStore) == 0 {
		return ""
	}

	left, right, mid := 1, len(localStore), 0
	if timestamp < localStore[0].timestamp {
		return ""
	} else if timestamp > localStore[right-1].timestamp {
		return localStore[right-1].value
	}

	for left < right {
		mid = (left + right) / 2
		if timestamp == localStore[mid].timestamp {
			return localStore[mid].value
		} else if localStore[mid].timestamp < timestamp {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return localStore[left-1].value
}
