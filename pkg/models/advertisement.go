package models

type Advertisement struct {
	Id 			uint `json:"-"`
	Title 		string `json:"title"`
	Description string `json:"description"`
	Price		uint `json:"price"`
	Photos 		[]string `json:"photos"`
}

type AdvertisementGetRequest struct {
	Fields 	[]string `json:"fields"`
}

func (r *Advertisement) ValidateFields() bool {
	if r.Photos == nil {
		return false
	}

	if len(r.Photos) > 3 || len(r.Title) > 200 || len(r.Description) > 1000 {
		return false
	}

	return true
}

func (r *AdvertisementGetRequest) ValidateFields() bool {
	validFields := []string{ "description", "photos" }

	if r.Fields == nil {
		return true
	}

	for _, field := range r.Fields {
		temp := false

		for _, validField := range validFields {
			if field == validField {
				temp = true
			}
		}
		
		if temp == false {
			return false
		}
	}

	return true
}