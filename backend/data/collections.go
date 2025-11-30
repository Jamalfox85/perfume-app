package data



type Collection struct {
	Id					string     			`json:"id"`
	UserId   			*string   			`json:"user_id"`
	Name	  			string   			`json:"name"`
	Description   		*string      		`json:"description"`
	Perfumes 			[]Perfume			`json:"perfumes"`
}