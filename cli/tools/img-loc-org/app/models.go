package app

type RevGeoType struct {
	Type     string `json:"type"`
	Features []struct {
		Type       string `json:"type"`
		Properties struct {
			Datasource struct {
				Sourcename  string `json:"sourcename"`
				Attribution string `json:"attribution"`
				License     string `json:"license"`
				URL         string `json:"url"`
			} `json:"datasource"`
			Name         string  `json:"name"`
			Country      string  `json:"country"`
			CountryCode  string  `json:"country_code"`
			State        string  `json:"state"`
			City         string  `json:"city"`
			Postcode     string  `json:"postcode"`
			District     string  `json:"district"`
			Suburb       string  `json:"suburb"`
			Street       string  `json:"street"`
			Housenumber  any     `json:"housenumber"`
			Lon          float64 `json:"lon"`
			Lat          float64 `json:"lat"`
			Distance     float64 `json:"distance"`
			ResultType   string  `json:"result_type"`
			Formatted    string  `json:"formatted"`
			AddressLine1 string  `json:"address_line1"`
			AddressLine2 string  `json:"address_line2"`
			Category     string  `json:"category"`
			Timezone     struct {
				Name             string `json:"name"`
				OffsetSTD        string `json:"offset_STD"`
				OffsetSTDSeconds int    `json:"offset_STD_seconds"`
				OffsetDST        string `json:"offset_DST"`
				OffsetDSTSeconds int    `json:"offset_DST_seconds"`
				AbbreviationSTD  string `json:"abbreviation_STD"`
				AbbreviationDST  string `json:"abbreviation_DST"`
			} `json:"timezone"`
			PlusCode string `json:"plus_code"`
			Rank     struct {
				Importance float64 `json:"importance"`
				Popularity float64 `json:"popularity"`
			} `json:"rank"`
			PlaceID string `json:"place_id"`
		} `json:"properties"`
		Geometry struct {
			Type        string    `json:"type"`
			Coordinates []float64 `json:"coordinates"`
		} `json:"geometry"`
		Bbox []float64 `json:"bbox"`
	} `json:"features"`
	Query struct {
		Lat      float64 `json:"lat"`
		Lon      float64 `json:"lon"`
		PlusCode string  `json:"plus_code"`
	} `json:"query"`
}
