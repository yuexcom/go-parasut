package parasut

import (
	"fmt"
	"net/http"
)

// ContactsService ...
type ContactsService service

// List ...
func (s *ContactsService) List(opts *ContactListOptions) (*ContactsResponse, *http.Response, error) {

	url, err := addOptions("contacts", opts)

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

	url := fmt.Sprintf("contacts/%d", id)

	url, err := addOptions(url, opts)

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
func (s *ContactsService) Create(ccr *ContactCreateRequest) (*ContactResponse, *http.Response, error) {

	url := "contacts"

	req, err := s.client.NewRequest("POST", url, ccr)

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
func (s *ContactsService) Update(id int, cur *ContactUpdateRequest) (*ContactResponse, *http.Response, error) {

	url := fmt.Sprintf("contacts/%d", id)

	req, err := s.client.NewRequest("PUT", url, cur)

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
func (s *ContactsService) Delete(id int) (*ContactResponse, *http.Response, error) {

	url := fmt.Sprintf("contacts/%d", id)

	req, err := s.client.NewRequest("DELETE", url, nil)

	if err != nil {
		return nil, nil, err
	}

	cdr := &ContactResponse{}

	resp, err := s.client.Do(req, cdr)

	if err != nil {
		return nil, nil, err
	}

	return cdr, resp, nil
}
