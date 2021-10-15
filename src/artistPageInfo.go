package src

func ArtistPageInfo(artistID int) Artist {
	var artist Artist
	artist.ID = All.Artists[artistID].ID
	artist.Image = All.Artists[artistID].Image
	artist.Name = All.Artists[artistID].Name
	artist.Members = All.Artists[artistID].Members
	artist.CreationDate = All.Artists[artistID].CreationDate
	artist.FirstAlbum = All.Artists[artistID].FirstAlbum
	artist.DatesLocations = All.Relation.Index[artistID].DatesLocations
	return artist
}

// Mapping is
// func Mapping(Map map[string][]string) map[string][]string {
// 	Mapa := map[string][]string{}
// 	for i, v := range Map {
// 		str := ""
// 		runa := []rune(i)
// 		for _, val := range runa {
// 			if val == '-' {
// 				val = ','
// 				str += string(val)
// 				str += string(" ")
// 			} else if val == '_' {
// 				val = ' '
// 				str = string(val)
// 			} else {
// 				str += string(val)
// 			}
// 		}
// 		Mapa[str] = v
// 	}
// 	return Mapa
// }
