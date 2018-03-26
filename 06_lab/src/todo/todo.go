package main

func addToDo(descritpion string) {
	put(descritpion, descritpion)
}

func markDone(description string) bool {
	return true
}

func todo() []string {
	all, err := list()
	result := []string{}
	if err == nil {
		for _, v := range all {
			result = append(result, v[1])
		}
	}
	return result
}

func genKey() string {
	return randStringRunes(6)
}
