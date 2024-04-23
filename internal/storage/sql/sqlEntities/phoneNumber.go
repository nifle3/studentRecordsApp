package sqlEntities

import "github.com/google/uuid"

type PhoneNumber struct {
	Id          uuid.UUID `db:'id'`
	StudentId   uuid.UUID `db:'student_id'`
	CountryCode string    `db:'country_code'`
	CityCode    string    `db:'city_code'`
	Code        string    `db:'code'`
	Description string    `db:'description'`
}
