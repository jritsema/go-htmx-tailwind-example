package main

import "strconv"

var data []Company

type Company struct {
	ID      string
	Company string
	Contact string
	Country string
}

func init() {
	data = []Company{
		{
			ID:      "1",
			Company: "Amazon",
			Contact: "Jeff Bezos",
			Country: "United States",
		},
		{
			ID:      "2",
			Company: "Apple",
			Contact: "Tim Cook",
			Country: "United States",
		},
		{
			ID:      "3",
			Company: "Microsoft",
			Contact: "Satya Nadella",
			Country: "United States",
		},
	}
}

func getCompanyByID(id string) Company {
	var result Company
	for _, i := range data {
		if i.ID == id {
			result = i
			break
		}
	}
	return result
}

func updateCompany(company Company) {
	result := []Company{}
	for _, i := range data {
		if i.ID == company.ID {
			i.Company = company.Company
			i.Contact = company.Contact
			i.Country = company.Country
		}
		result = append(result, i)
	}
	data = result
}

func addCompany(company Company) {
	max := 0
	for _, i := range data {
		n, _ := strconv.Atoi(i.ID)
		if n > max {
			max = n
		}
	}
	max++
	id := strconv.Itoa(max)

	data = append(data, Company{
		ID:      id,
		Company: company.Company,
		Contact: company.Contact,
		Country: company.Country,
	})
}

func deleteCompany(id string) {
	result := []Company{}
	for _, i := range data {
		if i.ID != id {
			result = append(result, i)
		}
	}
	data = result
}
