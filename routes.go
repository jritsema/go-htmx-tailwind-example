package main

import (
	"net/http"

	"github.com/jritsema/gotoolbox/web"
)

// Delete -> DELETE /company/{id} -> delete, companys.html

// Edit   -> GET /company/edit/{id} -> row-edit.html
// Save   ->   PUT /company/{id} -> update, row.html
// Cancel ->	 GET /company/{id} -> nothing, row.html

// Add    -> GET /company/add/ -> companys-add.html (target body with row-add.html and row.html)
// Save   ->   POST /company -> add, companys.html (target body without row-add.html)
// Cancel ->	 GET /company -> nothing, companys.html

func index(r *http.Request) *web.Response {
	return web.HTML(http.StatusOK, html, "index.html", data, nil)
}

// GET /company/add
func companyAdd(r *http.Request) *web.Response {
	return web.HTML(http.StatusOK, html, "company-add.html", data, nil)
}

// /GET company/edit/{id}
func companyEdit(r *http.Request) *web.Response {
	id, _ := web.PathLast(r)
	row := getCompanyByID(id)
	return web.HTML(http.StatusOK, html, "row-edit.html", row, nil)
}

// GET /company
// GET /company/{id}
// DELETE /company/{id}
// PUT /company/{id}
// POST /company
func companies(r *http.Request) *web.Response {
	id, segments := web.PathLast(r)
	switch r.Method {

	case http.MethodDelete:
		deleteCompany(id)
		return web.HTML(http.StatusOK, html, "companies.html", data, nil)

	//cancel
	case http.MethodGet:
		if segments > 1 {
			//cancel edit
			row := getCompanyByID(id)
			return web.HTML(http.StatusOK, html, "row.html", row, nil)
		} else {
			//cancel add
			return web.HTML(http.StatusOK, html, "companies.html", data, nil)
		}

	//save edit
	case http.MethodPut:
		row := getCompanyByID(id)
		r.ParseForm()
		row.Company = r.Form.Get("company")
		row.Contact = r.Form.Get("contact")
		row.Country = r.Form.Get("country")
		updateCompany(row)
		return web.HTML(http.StatusOK, html, "row.html", row, nil)

	//save add
	case http.MethodPost:
		row := Company{}
		r.ParseForm()
		row.Company = r.Form.Get("company")
		row.Contact = r.Form.Get("contact")
		row.Country = r.Form.Get("country")
		addCompany(row)
		return web.HTML(http.StatusOK, html, "companies.html", data, nil)
	}

	return web.Empty(http.StatusNotImplemented)
}
