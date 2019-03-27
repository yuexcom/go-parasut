package parasut

import (
	"fmt"
	"net/http"
)

// ContactsService ...
type ContactsService service

// List ...
func (s *ContactsService) List(filter *ContactListFilter, opts *ContactListOptions) (*ContactsResponse, *http.Response, error) {

	params := contactListParams{
		Filter: filter,
		Sort:   opts.Sort,
		Page: &PageOptions{
			Number: opts.PageNumber,
			Size:   opts.PageSize,
		},
		Include: opts.Include,
	}

	url, err := addOptions("contacts", params)

	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	clr := &ContactsResponse{}

	resp, err := s.client.Do(req, &clr)

	if err != nil {
		return nil, nil, err
	}

	return clr, resp, nil
}

// Get ...
func (s *ContactsService) Get(id int, opts *ContactListOptions) (*ContactResponse, *http.Response, error) {

	params := contactListParams{
		Include: opts.Include,
	}

	url := fmt.Sprintf("contacts/%d", id)

	url, err := addOptions(url, params)

	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", url, nil)

	if err != nil {
		return nil, nil, err
	}

	cr := &ContactResponse{}

	resp, err := s.client.Do(req, cr)

	if err != nil {
		return nil, nil, err
	}

	return cr, resp, nil
}

// Create ...
func (s *ContactsService) Create(c *Contact) (*ContactResponse, *http.Response, error) {

	obj := s.prepareObj(c)

	url := "contacts"

	req, err := s.client.NewRequest("POST", url, obj)

	if err != nil {
		return nil, nil, err
	}

	cr := &ContactResponse{}

	resp, err := s.client.Do(req, cr)

	if err != nil {
		return nil, nil, err
	}

	return cr, resp, nil
}

// Update ...
func (s *ContactsService) Update(id int, c *Contact) (*ContactResponse, *http.Response, error) {

	obj := s.prepareObj(c)

	url := fmt.Sprintf("contacts/%d", id)

	req, err := s.client.NewRequest("PUT", url, obj)

	if err != nil {
		return nil, nil, err
	}

	cr := &ContactResponse{}

	resp, err := s.client.Do(req, cr)

	if err != nil {
		return nil, nil, err
	}

	return cr, resp, nil
}

// Delete ...
func (s *ContactsService) Delete(id int) (*http.Response, error) {

	url := fmt.Sprintf("contacts/%d", id)

	req, err := s.client.NewRequest("DELETE", url, nil)

	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *ContactsService) prepareObj(c *Contact) *contactRequest {

	var contactPeopleData []*contactPersonData

	for _, cp := range c.ContactPeople {
		tmp := &contactPersonData{
			ID:   cp.ID,
			Type: "contact_people",
			Attributes: &contactPerson{
				Name:  cp.Name,
				Email: cp.Email,
				Phone: cp.Phone,
				Notes: cp.Notes,
			},
		}
		contactPeopleData = append(contactPeopleData, tmp)
	}

	obj := &contactRequest{
		Data: &contactData{
			Type: "contacts",
			Attributes: &contact{
				Name:        c.Name,
				AccountType: c.AccountType,
				ContactType: c.ContactType,
				Email:       c.Email,
				ShortName:   c.ShortName,
				TaxNumber:   c.TaxNumber,
				TaxOffice:   c.TaxOffice,
				City:        c.City,
				District:    c.District,
				Address:     c.Address,
				Phone:       c.Phone,
			},
			Relationships: &relationships{
				"contact_people": &relationship{
					Data: contactPeopleData,
				},
				"category": &relationship{
					Data: categoryData{
						ID:   c.CategoryID,
						Type: "item_categories",
					},
				},
			},
		},
	}

	return obj
}
