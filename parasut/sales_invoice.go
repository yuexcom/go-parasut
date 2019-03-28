package parasut

import (
	"fmt"
	"net/http"
)

// SalesInvoicesService ...
type SalesInvoicesService service

// List ...
func (s *SalesInvoicesService) List(opts *SalesInvoiceListOptions) (*SalesInvoicesResponse, *http.Response, error) {

	url, err := addOptions("sales_invoices", opts)

	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", url, nil)

	if err != nil {
		return nil, nil, err
	}

	silr := &SalesInvoicesResponse{}

	resp, err := s.client.Do(req, silr)

	if err != nil {
		return nil, nil, err
	}

	return silr, resp, nil
}

// Get ...
func (s *SalesInvoicesService) Get(id string, opts *SalesInvoiceGetOptions) (*SalesInvoiceResponse, *http.Response, error) {

	url := fmt.Sprintf("sales_invoices/%s", id)

	url, err := addOptions(url, opts)

	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", url, nil)

	if err != nil {
		return nil, nil, err
	}

	sir := &SalesInvoiceResponse{}

	resp, err := s.client.Do(req, sir)

	if err != nil {
		return nil, nil, err
	}

	return sir, resp, nil
}

// Create ...
func (s *SalesInvoicesService) Create(si *SalesInvoiceRequest) (*SalesInvoiceResponse, *http.Response, error) {

	url := "sales_invoices"

	req, err := s.client.NewRequest("POST", url, si)

	if err != nil {
		return nil, nil, err
	}

	sir := &SalesInvoiceResponse{}

	resp, err := s.client.Do(req, sir)

	if err != nil {
		return nil, nil, err
	}

	return sir, resp, nil
}

// Update ...
func (s *SalesInvoicesService) Update(id string, si *SalesInvoiceRequest) (*SalesInvoiceResponse, *http.Response, error) {

	url := fmt.Sprintf("sales_invoices/%s", id)

	req, err := s.client.NewRequest("PUT", url, si)

	if err != nil {
		return nil, nil, err
	}

	sir := &SalesInvoiceResponse{}

	resp, err := s.client.Do(req, sir)

	if err != nil {
		return nil, nil, err
	}

	return sir, resp, nil
}

// Delete ...
func (s *SalesInvoicesService) Delete(id string) (*http.Response, error) {

	url := fmt.Sprintf("sales_invoices/%s", id)

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
