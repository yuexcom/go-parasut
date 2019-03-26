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
func (s *ContactsService) Get(id int, filter *ContactListFilter) (*ContactResponse, *http.Response, error) {

	obj := contactListParams{
		Filter:  filter,
		Sort:    "",
		Page:    &PageOptions{},
		Include: "",
	}

	url := fmt.Sprintf("contacts/%d", id)

	url, err := addOptions(url, obj)

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
func (s *ContactsService) Create(c *Contact, cp []*ContactPerson) (*ContactResponse, *http.Response, error) {

	var contactPeopleData []ContactPersonData

	for _, attributes := range cp {
		tmp := ContactPersonData{
			Type:       "contact_people",
			Attributes: attributes,
		}
		contactPeopleData = append(contactPeopleData, tmp)
	}

	obj := contactCreate{
		Data: &ContactData{
			Type:       "contacts",
			Attributes: c,
			Relationships: &Relationships{
				"contact_people": &Relationship{
					Data: contactPeopleData,
				},
			},
		},
	}

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
func (s *ContactsService) Update(id int, c *Contact, cp []*ContactPerson) (*ContactResponse, *http.Response, error) {

	var contactPeopleData []ContactPersonData

	for _, attributes := range cp {

		tmp := ContactPersonData{
			ID:   attributes.ID,
			Type: "contact_people",
			Attributes: &ContactPerson{
				Name:  attributes.Name,
				Email: attributes.Email,
				Phone: attributes.Phone,
				Notes: attributes.Notes,
			},
		}

		contactPeopleData = append(contactPeopleData, tmp)
	}

	obj := contactUpdate{
		Data: &ContactData{
			Type:       "contacts",
			Attributes: c,
			Relationships: &Relationships{
				"contact_people": &Relationship{
					Data: contactPeopleData,
				},
			},
		},
	}

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
