package main

import "strconv"

// Company represents a company
type Company struct {
	ID      string
	Company string
	Contact string
	Country string
}

// List of companies
type Companies struct {
	companies []Company
}

// Company data access
var data Companies

func init() {

	data = Companies{
		companies: []Company{
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
		},
	}
}

func (c *Companies) getByID(id string) Company {
	var result Company
	for _, i := range c.companies {
		if i.ID == id {
			result = i
			break
		}
	}
	return result
}

func (c *Companies) update(company Company) {
	result := []Company{}
	for _, i := range c.companies {
		if i.ID == company.ID {
			i.Company = company.Company
			i.Contact = company.Contact
			i.Country = company.Country
		}
		result = append(result, i)
	}
	c.companies = result
}

func (c *Companies) add(company Company) {
	max := 0
	for _, i := range c.companies {
		n, _ := strconv.Atoi(i.ID)
		if n > max {
			max = n
		}
	}
	max++
	id := strconv.Itoa(max)

	c.companies = append(c.companies, Company{
		ID:      id,
		Company: company.Company,
		Contact: company.Contact,
		Country: company.Country,
	})
}

func (c *Companies) delete(id string) {
	result := []Company{}
	for _, i := range c.companies {
		if i.ID != id {
			result = append(result, i)
		}
	}
	c.companies = result
}
