package types

type WalkingReq struct {
	Origin        string `schema:"origin"`
	Destination   string `schema:"destination"`
	OriginId      string `schema:"origin_id,omitempty"`
	DestinationId string `schema:"destination_id,omitempty"`
}

type WalkResp struct {
	Count string `json:"count"`
	Route struct {
		Origin      string `json:"origin"`
		Destination string `json:"destination"`
		Paths       []struct {
			Distance string `json:"distance"`
			Duration string `json:"duration"`
			Steps    []struct {
				Instruction     string      `json:"instruction"`
				Orientation     interface{} `json:"orientation"`
				Road            interface{} `json:"road"`
				Distance        string      `json:"distance"`
				Duration        string      `json:"duration"`
				Polyline        string      `json:"polyline"`
				Action          interface{} `json:"action"`
				AssistantAction interface{} `json:"assistant_action"`
				WalkType        string      `json:"walk_type"`
			} `json:"steps"`
		} `json:"paths"`
	} `json:"route"`
}

type IntegratedReq struct {
	Origin      string `schema:"origin"`
	Destination string `schema:"destination"`
	City        string `schema:"city"`
	Cityd       string `schema:"cityd"`
	Extensions  string `schema:"extensions,omitempty"`
	Strategy    int    `schema:"strategy,omitempty"`
	Nightflag   int    `schema:"nightflag,omitempty"`
	Date        string `schema:"date,omitempty"`
	Time        string `schema:"time,omitempty"`
}

type IntegratedResp struct {
	Count string `json:"count"`
	Route struct {
		Origin      string `json:"origin"`
		Destination string `json:"destination"`
		Distance    string `json:"distance"`
		TaxiCost    string `json:"taxi_cost"`
		Transits    []struct {
			Cost            string `json:"cost"`
			Duration        string `json:"duration"`
			Nightflag       string `json:"nightflag"`
			WalkingDistance string `json:"walking_distance"`
			Distance        string `json:"distance"`
			Missed          string `json:"missed"`
			Segments        []struct {
				Taxi    []interface{} `json:"taxi"`
				Walking interface{}   `json:"walking"`
				Bus     struct {
					Buslines []struct {
						DepartureStop struct {
							Name     string `json:"name"`
							Id       string `json:"id"`
							Location string `json:"location"`
						} `json:"departure_stop"`
						ArrivalStop struct {
							Name     string `json:"name"`
							Id       string `json:"id"`
							Location string `json:"location"`
						} `json:"arrival_stop"`
						Name       string      `json:"name"`
						Id         string      `json:"id"`
						Type       string      `json:"type"`
						Distance   string      `json:"distance"`
						Duration   string      `json:"duration"`
						Polyline   string      `json:"polyline"`
						Bustimetag string      `json:"bustimetag"`
						StartTime  interface{} `json:"start_time"`
						EndTime    interface{} `json:"end_time"`
						ViaNum     string      `json:"via_num"`
						ViaStops   []struct {
							Name     string `json:"name"`
							Id       string `json:"id"`
							Location string `json:"location"`
						} `json:"via_stops"`
					} `json:"buslines"`
				} `json:"bus"`
				Entrance interface{} `json:"entrance"`
				Exit     interface{} `json:"exit"`
				Railway  struct {
					ViaStops []interface{} `json:"via_stops"`
					Alters   []interface{} `json:"alters"`
					Spaces   []interface{} `json:"spaces"`
				} `json:"railway"`
			} `json:"segments"`
		} `json:"transits"`
	} `json:"route"`
}

type DrivingReq struct {
	Origin          string `schema:"origin"`
	Destination     string `schema:"destination"`
	DestinationId   string `schema:"destinationid,omitempty"`
	DestinationType string `schema:"destinationtype,omitempty"`
	Strategy        int    `schema:"strategy,omitempty"`
	Waypoint        string `schema:"waypoint,omitempty"`
	AvoidPolygons   string `schema:"avoidpolygons,omitempty"`
	Province        string `schema:"province,omitempty"`
	Number          string `schema:"number,omitempty"`
	Cartype         string `schema:"cartype,omitempty"`
	Ferry           int    `schema:"ferry,omitempty"`
	Roadaggregation bool   `schema:"roadaggregation,omitempty"`
	Nosteps         bool   `schema:"nosteps,omitempty"`
}

type DrivingResp struct {
	Count string `json:"count"`
	Route struct {
		Origin      string `json:"origin"`
		Destination string `json:"destination"`
		Paths       []struct {
			Distance     string `json:"distance"`
			Duration     string `json:"duration"`
			Strategy     string `json:"strategy"`
			Tolls        string `json:"tolls"`
			TollDistance string `json:"toll_distance"`
			Steps        []struct {
				Instruction     string        `json:"instruction"`
				Orientation     string        `json:"orientation"`
				Distance        string        `json:"distance"`
				Tolls           string        `json:"tolls"`
				TollDistance    string        `json:"toll_distance"`
				TollRoad        []interface{} `json:"toll_road"`
				Duration        string        `json:"duration"`
				Polyline        string        `json:"polyline"`
				Action          interface{}   `json:"action"`
				AssistantAction interface{}   `json:"assistant_action"`
				Road            string        `json:"road,omitempty"`
			} `json:"steps"`
			Restriction   string `json:"restriction"`
			TrafficLights string `json:"traffic_lights"`
		} `json:"paths"`
	} `json:"route"`
}
