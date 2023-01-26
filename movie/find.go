package movie

func FindName(imdbId string) string {
	switch imdbId {
	case "tt4154796":
		return "Avengers Endgame"
	case "tt1825683":
		return "Black Panther"
	}
	return "not found"
}
