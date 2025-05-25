package server

import "github.com/etzba/gopu/wire"

// locations is a global variable saved in memory to allow post and put requests to change it
// if the server restarted it will be as the list. It is saved in memory as this server won't include database connection
var locations = []wire.Location{
	{
		Id:         1,
		Latitude:   29.97921394086218,
		Longtitude: 31.134401361961217,
		Name:       "Great Pyramid at Giza",
	}, {
		Id:         2,
		Latitude:   40.689412075433374,
		Longtitude: -74.04439311569635,
		Name:       "Statue of Liberty",
	}, {
		Id:         3,
		Latitude:   48.858648890929786,
		Longtitude: 2.294352550396703,
		Name:       "Eiffel Tower",
	}, {
		Id:         4,
		Latitude:   -25.343756441863693,
		Longtitude: 131.03478543537872,
		Name:       "Ayers Rock (Uluru)",
	}, {
		Id:         5,
		Latitude:   -3.0663960573117492,
		Longtitude: 37.35571305086078,
		Name:       "Mt. Kilimanjaro",
	}, {
		Id:         6,
		Latitude:   27.175402426543172,
		Longtitude: 78.04221727597596,
		Name:       "Taj Mahal",
	}, {
		Id:         7,
		Latitude:   37.80829524296391,
		Longtitude: -122.47502910436772,
		Name:       "Golden Gate Bridge",
	}, {
		Id:         8,
		Latitude:   -13.162948060479428,
		Longtitude: -72.545229918406,
		Name:       "Machu Picchu",
	}, {
		Id:         9,
		Latitude:   52.516418207265474,
		Longtitude: 13.377693367823799,
		Name:       "Brandenburg Gate",
	}, {
		Id:         10,
		Latitude:   43.64272145135072,
		Longtitude: -79.38711044806657,
		Name:       "CN Tower",
	}, {
		Id:         11,
		Latitude:   35.361324531464874,
		Longtitude: 138.72719167099413,
		Name:       "Mt. Fuji",
	}, {
		Id:         12,
		Latitude:   55.75266168393862,
		Longtitude: 37.62310824099749,
		Name:       "St. Basil's Cathedral",
	}, {
		Id:         13,
		Latitude:   38.625311272663076,
		Longtitude: -90.18673967532403,
		Name:       "Gateway Arch",
	},
}
