package main

func main() {
	days(8)
}

func days(day int) {
	switch day {
	case 1:
		println("sunday")
		break
	case 2:
		println("monday")
		break
	case 3:
		println("tuesday")
		break
	case 4:
		println("wednesday")
		break
	case 5:
		println("thursday")
		break
	case 6:
		println("friday")
		break
	case 7:
		println("saturday")
		break
	default:
		println("enter only 1 - 7")
	}
}
