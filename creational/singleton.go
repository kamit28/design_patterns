package creational

import "sync"

type Database interface {
	GetPopulation(name string) int
}

type database struct {
	capitals map[string]int
}

func (db *database) GetPopulation(name string) int {
	return db.capitals[name]
}

var once sync.Once
var instance *database

func GetDatabase() *database {
	once.Do(func() {
		db := database{}
		db.capitals = readData()
		instance = &db
	})
	return instance
}

// Ideally this would be read from a File / Database / external interface
func readData() map[string]int {
	m := map[string]int{
		"Tokyo":        37435191,
		"Delhi":        29399141,
		"Shanghai":     26317104,
		"Sao Paulo":    21846507,
		"Mexico City":  21671908,
		"Cairo":        20484965,
		"Dhaka":        20283552,
		"Mumbai":       20185064,
		"Beijing":      20035455,
		"Osaka":        19222665,
		"Karachi":      16459472,
		"Chongqing":    16382376,
		"Istanbul":     15415197,
		"Buenos Aires": 15257673,
		"Kolkata":      14974073,
		"Kinshasa":     14970460,
		"Lagos":        14862111,
		"Manila":       14158573,
		"Tianjin":      13794450,
		"Guangzhou":    13635397,
	}
	return m
}
