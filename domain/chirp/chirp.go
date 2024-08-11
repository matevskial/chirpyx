package chirp

type Chirp struct {
	Id   int
	Body string
}

type ChirpRepository interface {
	Create(body string) (Chirp, error)
	FindAll() ([]Chirp, error)
}
