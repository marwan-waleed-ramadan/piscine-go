package piscine

func PodiumPosition(podium [][]string) [][]string {
	size := len(podium)
	for i := 0; i < size/2; i++ {
		podium[i], podium[size-1-i] = podium[size-1-i], podium[i]
	}
	return podium
}
