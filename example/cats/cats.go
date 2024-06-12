package cats

import "time"

type Fact struct {
	Id        string    `json:"_id"`
	V         int       `json:"__v"`
	Text      string    `json:"text"`
	UpdatedAt time.Time `json:"updatedAt"`
	Deleted   bool      `json:"deleted"`
	Source    string    `json:"source"`
}

type CatFactsAPI struct {
	Facts       func() ([]Fact, error)                   `method:"GET" path:"/facts"`
	AnimalFacts func(animal_type string) ([]Fact, error) `method:"GET" path:"/facts" query:"animal_type"`
	Fact        func(id string) (*Fact, error)           `method:"GET" path:"/facts/{id}"`
}
