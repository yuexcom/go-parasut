package parasut

import (
	"fmt"
	"net/http"
)

// SalesInvoicesService ...
type SalesInvoicesService service

// List ...
func (s *SalesInvoicesService) List(opts *ListOptions) (*SalesInvoiceListResponse, *http.Response, error) {

	url, err := addOptions("sales_invoices", opts)

	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", url, nil)

	if err != nil {
		return nil, nil, err
	}

	silr := &SalesInvoiceListResponse{}

	resp, err := s.client.Do(req, silr)

	if err != nil {
		return nil, nil, err
	}

	return silr, resp, nil
}

// Get ...
func (s *SalesInvoicesService) Get(id int) (*SalesInvoiceResponse, *http.Response, error) {

	url := fmt.Sprintf("sales_invoices/%d", id)

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
func (s *SalesInvoicesService) Create(ccr *SalesInvoiceCreateRequest) (*SalesInvoiceResponse, *http.Response, error) {

	url := "sales_invoices"

	req, err := s.client.NewRequest("POST", url, ccr)

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
func (s *SalesInvoicesService) Update(id int, cur *SalesInvoiceUpdateRequest) (*SalesInvoiceResponse, *http.Response, error) {

	url := fmt.Sprintf("sales_invoices/%d", id)

	req, err := s.client.NewRequest("PUT", url, cur)

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
func (s *SalesInvoicesService) Delete(id int) (*SalesInvoiceDeleteResponse, *http.Response, error) {

	url := fmt.Sprintf("sales_invoices/%d", id)

	req, err := s.client.NewRequest("DELETE", url, nil)

	if err != nil {
		return nil, nil, err
	}

	sidr := &SalesInvoiceDeleteResponse{}

	resp, err := s.client.Do(req, sidr)

	if err != nil {
		return nil, nil, err
	}

	return sidr, resp, nil
}
