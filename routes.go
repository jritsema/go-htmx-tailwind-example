package main

import (
	"net/http"

	"github.com/jritsema/gotoolbox/web"
)

///////////////////////////////////////////////////////////
//  UI Interactions
///////////////////////////////////////////////////////////
//
// Add -> GET /company/add/ -> company-add.html
// (target body with row-add.html and row.html)
//
//    Save -> POST /company -> add, companies.html
//		(target body without row-add.html)
//
//		Cancel -> GET /company -> nothing, companies.html
//
// Edit -> GET /company/edit/{id} -> row-edit.html
// 		Save -> PUT /company/{id} -> update, row.html
//		Cancel -> GET /company/{id} -> nothing, row.html
//
// Delete -> DELETE /company/{id} -> delete, companies.html
///////////////////////////////////////////////////////////

func index(r *http.Request) *web.Response {
	return web.HTML(http.StatusOK, html, "index.html", data.companies, nil)
}

// GET /company/add
func addCompany(r *http.Request) *web.Response {
	return web.HTML(http.StatusOK, html, "company-add.html", data.companies, nil)
}

// POST /company
func saveNewCompany(r *http.Request) *web.Response {
	row := Company{}
	r.ParseForm()
	row.Company = r.Form.Get("company")
	row.Contact = r.Form.Get("contact")
	row.Country = r.Form.Get("country")
	data.add(row)
	return web.HTML(http.StatusOK, html, "companies.html", data.companies, nil)
}

// GET /company
func cancelSaveNewCompany(r *http.Request) *web.Response {
	return web.HTML(http.StatusOK, html, "companies.html", data.companies, nil)
}

// /GET company/edit/{id}
func editCompany(r *http.Request) *web.Response {
	id := r.PathValue("id")
	row := data.getByID(id)
	return web.HTML(http.StatusOK, html, "row-edit.html", row, nil)
}

// PUT /company/{id}
func saveExistingCompany(r *http.Request) *web.Response {
	id := r.PathValue("id")
	row := data.getByID(id)
	r.ParseForm()
	row.Company = r.Form.Get("company")
	row.Contact = r.Form.Get("contact")
	row.Country = r.Form.Get("country")
	data.update(row)
	return web.HTML(http.StatusOK, html, "row.html", row, nil)
}

// GET /company/{id}
func cancelSaveExistingCompany(r *http.Request) *web.Response {
	id := r.PathValue("id")
	row := data.getByID(id)
	return web.HTML(http.StatusOK, html, "row.html", row, nil)
}

// DELETE /company/{id}
func deleteCompany(r *http.Request) *web.Response {
	id := r.PathValue("id")
	data.delete(id)
	return web.HTML(http.StatusOK, html, "companies.html", data.companies, nil)
}
