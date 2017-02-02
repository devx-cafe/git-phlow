package gitwrapper

type Fetch struct {
	gitFetchCommand string
}

//NewFetch
//Gives a new fetch command
func NewFetch() *Fetch {
	return &Fetch{gitFetchCommand:"fetch"}

}

//Fetch
func Fetch()  {

}




