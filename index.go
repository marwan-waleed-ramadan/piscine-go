package piscine

func Index(s string, toFind string) int {
	toFindLen := len(toFind)
	sLen := len(s)
	for index := 0; index <= sLen-toFindLen; index++ {
		if s[index:index+toFindLen] == toFind {
			return index
		}
	}
	return -1
}
