package model

type VaccineDetail struct {
	Administered              int    `json:"administered"`
	PeopleVaccinated          int    `json:"people_vaccinated"`
	PeoplePartiallyVaccinated int    `json:"people_partially_vaccinated"`
	Country                   string `json:"country"`
	Population                int    `json:"population"`
	SqKmArea                  string `json:"sqKmArea"`
	LifeExpectancy            string `json:"life_expectancy"`
	ElevationInMeters         int    `json:"elevation_in_meters"`
	Continent                 string `json:"continent"`
	Abbreviation              string `json:"abbreviation"`
	Location                  string `json:"location"`
	Iso                       int    `json:"iso"`
	CapitalCity               string `json:"capital_city"`
	Update                    string `json:"updated"`
}
