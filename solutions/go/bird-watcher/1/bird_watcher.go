package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
    var sum int =0
	for _,value := range birdsPerDay{
        sum+=value
    }
    return sum
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	birds:=birdsPerDay[week*7-7:week*7]
    return TotalBirdCount(birds)
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for index,value:=range birdsPerDay{
        if index%2==0{
            birdsPerDay[index]=value+1
        }
    }
    return birdsPerDay
}
