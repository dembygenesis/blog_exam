package data_utils

func GetDistinctValuesOfArrayOfStrings(stringSlice []string) []string {
	hash := make(map[string]string)
	for _, v := range stringSlice {
		hash[v] = ""
	}
	var output []string
	for i, _ := range hash {
		output = append(output, i)
	}
	return output
}

func GetDistinctValuesOfArrayOfInts(intSlice []int) []int {
	hash := make(map[int]int)
	for _, v := range intSlice {
		hash[v] = 0
	}
	var output []int
	for i, _ := range hash {
		output = append(output, i)
	}
	return output
}