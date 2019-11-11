package address

type LocationAddress struct {
	Country  string `json:"country"`
	Province string `json:"province"`
	City     string `json:"city"`
}

var cityDB *City
var cityDbError error

func init() {
	cityDB, cityDbError = NewCity("address/address.datx")
}

func GetAddressByIP(ip string) (*LocationAddress, error) {
	if cityDbError != nil {
		return nil, cityDbError
	}

	locations, err := cityDB.Find(ip)
	if err != nil {
		return nil, err
	}

	address := &LocationAddress{}
	if len(locations) >= 3 {
		address.Country = locations[0]
		address.Province = locations[1]
		address.City = locations[2]
	}

	return address, nil
}
