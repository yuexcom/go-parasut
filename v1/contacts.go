package parasut

import "fmt"

// GetContacts ...
func (p *Parasut) GetContacts(opts *ListOptions) (*ContactListResponse, error) {

	url, err := prepareURL("contacts", opts)

	if err != nil {
		return nil, err
	}

	req, err := p.newRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	clr := &ContactListResponse{}

	_, err = p.doRequest(req, clr)

	if err != nil {
		return nil, err
	}

	return clr, nil
}

// GetContact ...
func (p *Parasut) GetContact(id int) (*ContactResponse, error) {

	url := fmt.Sprintf("contacts/%d", id)

	req, err := p.newRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	cr := &ContactResponse{}

	_, err = p.doRequest(req, cr)

	if err != nil {
		return nil, err
	}

	return cr, nil
}

// CreateContact ...
func (p *Parasut) CreateContact(ccr *ContactCreateRequest) (*ContactResponse, error) {

	url, err := prepareURL("contacts", nil)

	if err != nil {
		return nil, err
	}

	req, err := p.newRequest("POST", url, ccr)

	if err != nil {
		return nil, err
	}

	cr := &ContactResponse{}

	_, err = p.doRequest(req, cr)

	if err != nil {
		return nil, err
	}

	return cr, nil
}

// UpdateContact ...
func (p *Parasut) UpdateContact(id int, cur *ContactUpdateRequest) (*ContactResponse, error) {

	endpoint := fmt.Sprintf("contacts/%d", id)

	url, err := prepareURL(endpoint, nil)

	if err != nil {
		return nil, err
	}

	req, err := p.newRequest("PUT", url, cur)

	if err != nil {
		return nil, err
	}

	cr := &ContactResponse{}

	_, err = p.doRequest(req, cr)

	if err != nil {
		return nil, err
	}

	return cr, nil
}

// DeleteContact ...
func (p *Parasut) DeleteContact(id int) (*ContactDeleteResponse, error) {

	endpoint := fmt.Sprintf("contacts/%d", id)

	url, err := prepareURL(endpoint, nil)

	if err != nil {
		return nil, err
	}

	req, err := p.newRequest("DELETE", url, nil)

	if err != nil {
		return nil, err
	}

	cdr := &ContactDeleteResponse{}

	_, err = p.doRequest(req, cdr)

	if err != nil {
		return nil, err
	}

	return cdr, nil
}
