package models

type ConditionResponse struct {
	Alerts    []*Alert `json:"alerts,omitempty"`
	Summary   string   `json:"summary"`
	Temp      *float64 `json:"temp"`
	FeelsTemp *float64 `json:"feels_temp"`
	FeelsLike string   `json:"feels_like"`
	Error     string   `json:"error,omitempty"`
}

type OpenWeatherOneCallResponse struct {
	Code     *int              `json:"cod"` // present on error
	Lat      float64           `json:"lat"`
	Lon      float64           `json:"lon"`
	Timezone string            `json:"timezone"`
	Current  CurrentConditions `json:"current"`
	Alerts   []*Alert          `json:"alerts"`
}

type CurrentConditions struct {
	Temp      *float64 `json:"temp"`
	FeelsLike *float64 `json:"feels_like"`
	Weather   []*WeatherItem
}

type WeatherItem struct {
	ID          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type Alert struct {
	SenderName  string   `json:"sender_name"`
	Event       string   `json:"event"`
	Start       int64    `json:"start"`
	End         int64    `json:"end"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
}
