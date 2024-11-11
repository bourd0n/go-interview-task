package models

import "os"

var LocalCountriesStore = make([]*CountryModel, 0)
var LocalTownsStore = make([]*Town, 0)
var LocalContinentStore = make([]*Continent, 0)

func CreateCountry(c *CountryModel) {
	if os.Getenv("IN_MEMORY") == "true" {

		_ = append(LocalCountriesStore, c)

		for i := range c.TownsNames {
			_ = append(LocalTownsStore, &Town{
				Name:  c.TownsNames[i],
				Popul: 0,
			})
		}

		_ = append(LocalContinentStore, &Continent{Name: c.ContinentName})

	} else if os.Getenv("DB") == "true" {
		//call sql
	}
}

func GetCountry(x string) *CountryModel {
	if os.Getenv("IN_MEMORY") == "true" {

		for i := range LocalCountriesStore {
			if LocalCountriesStore[i].Name == x {
				return LocalCountriesStore[i]
			}
		}

	} else if os.Getenv("DB") == "true" {
		// call sql
		return nil
	}
	return nil
}

func SetTownPopulation(town *Town) {
	if os.Getenv("IN_MEMORY") == "true" {

		for i := range LocalTownsStore {
			if LocalTownsStore[i].Name == town.Name {
				LocalTownsStore[i].Popul = town.Popul
			}
		}

	} else if os.Getenv("DB") == "true" {
		// call sql
	}
}
