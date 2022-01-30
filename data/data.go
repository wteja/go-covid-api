package data

type CovidData struct {
	ConfirmDate    string
	No             *int
	Age            *int
	Gender         string
	GenderEn       string
	Nation         string
	NationEn       string
	Province       string
	ProvinceId     *int
	District       string
	ProvinceEn     string
	StatQuarantine *int
}

type CovidDataList struct {
	Data []CovidData
}

// Get a group of data.
func (c *CovidDataList) Group() (map[string]int, map[string]int) {
	provinces := make(map[string]int)
	ages := map[string]int{
		"0-30":  0,
		"31-60": 0,
		"61+":   0,
		"N/A":   0,
	}

	var provinceKey string

	for _, value := range c.Data {
		// Count by ages.
		ages[GetAgeKey(value.Age)] += 1

		// Create dynamic key for provinces map.
		provinceKey = GetProvinceKey(value.Province)
		if _, ok := provinces[provinceKey]; !ok {
			provinces[provinceKey] = 1
		} else {
			provinces[provinceKey] += 1
		}
	}

	return provinces, ages
}

// Returns corresponding key for AgeGroup map.
func GetAgeKey(age *int) string {
	if age == nil || *age < 0 {
		return "N/A"
	} else if *age >= 0 && *age <= 30 {
		return "0-30"
	} else if *age >= 31 && *age <= 60 {
		return "31-60"
	} else {
		return "61+"
	}
}

// Returns corresponding key for Profince map.
func GetProvinceKey(province string) string {
	if len(province) == 0 {
		return "N/A"
	} else {
		return province
	}
}
