package main

type covidGlobalResult		struct {
	Global					struct {
		Confirmed			int 		`json:"TotalConfirmed"`
		Deaths				int			`json:"TotalDeaths"`
		Recovered			int			`json:"TotalRecovered"`
	}									`json:"Global"`
}

type covidCountryResults	[]struct {
	Country     			string		`json:"Country"`
	Confirmed   			int   		`json:"Confirmed"`
	Deaths      			int   		`json:"Deaths"`
	Recovered   			int   		`json:"Recovered"`
}