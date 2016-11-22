package tedfeed

<<<<<<< HEAD
=======
import "encoding/xml"

// Link maps a subset of the atom link element
>>>>>>> b0c3149d7bd5a0e4e0534915b42a809b3da9ddfc
type Link struct {
	Rel  string `xml:"rel,attr"`
	HRef string `xml:"href,attr"`
}

<<<<<<< HEAD
=======
// Entry maps a subset of the atom entry element
>>>>>>> b0c3149d7bd5a0e4e0534915b42a809b3da9ddfc
type Entry struct {
	Id          string `xml:"id"`
	TalkId      string `xml:"ted:talkid"`
	ImageURL    string `xml:"ted:image"`
	Duration    string `xml:"ted:duration"`
	SpeakerName string `xml:"ted:speakername"`
	Title       string `xml:"title"`
	Link        []Link `xml:"link"`
	Update      string `xml:"update"`
	Summary     string `xml:"summary"`
}

<<<<<<< HEAD
type Feed struct {
	XMLName string  `xml:"feed"`
	Updated string  `xml:"updated"`
	Title   string  `xml:"title"`
	Entry   []Entry `xml:"entry"`
}

//exercise 3: adding method who iterate over Feed Type and returns a map[Title]Link
func (fd Feed) GetLinksList() map[string]string {

	//creating map
	m := make(map[string]string)

	//iterate over tedfeed.Entry[].Link[]
	for _, entry := range fd.Entry {
		for _, link := range entry.Link {

			//must get only Rel == "enclosure" link
=======
// Feed maps a subset of the main atom feed element
type Feed struct {
	XMLName xml.Name `xml:"feed"`
	Updated string   `xml:"updated"`
	Title   string   `xml:"title"`
	Entry   []Entry  `xml:"entry"`
}

// GetLinksList returns a map with the title and the URL of the video to download
func (fd Feed) GetLinksList() map[string]string {
	// Looking for links containing enclosures
	m := make(map[string]string)
	for _, entry := range fd.Entry {
		for _, link := range entry.Link {
>>>>>>> b0c3149d7bd5a0e4e0534915b42a809b3da9ddfc
			if link.Rel == "enclosure" {
				m[entry.Title] = link.HRef
			}
		}
	}
<<<<<<< HEAD

=======
>>>>>>> b0c3149d7bd5a0e4e0534915b42a809b3da9ddfc
	return m
}
