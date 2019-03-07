package parasut

import "fmt"

// GetSalesInvoices ...
func (p *Parasut) GetSalesInvoices(opts *ListOptions) (*SalesInvoiceListResponse, error) {

	url, err := prepareURL("sales_invoices", opts)

	if err != nil {
		return nil, err
	}

	req, err := p.newRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	clr := &SalesInvoiceListResponse{}

	_, err = p.doRequest(req, clr)

	if err != nil {
		return nil, err
	}

	return clr, nil
}

// GetSalesInvoice ...
func (p *Parasut) GetSalesInvoice(id int) (*SalesInvoiceResponse, error) {

	url := fmt.Sprintf("sales_invoices/%d", id)

	req, err := p.newRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	cr := &SalesInvoiceResponse{}

	_, err = p.doRequest(req, cr)

	if err != nil {
		return nil, err
	}

	return cr, nil
}

// CreateSalesInvoice ...
func (p *Parasut) CreateSalesInvoice(ccr *SalesInvoiceCreateRequest) (*SalesInvoiceResponse, error) {

	url, err := prepareURL("sales_invoices", nil)

	if err != nil {
		return nil, err
	}

	req, err := p.newRequest("POST", url, ccr)

	if err != nil {
		return nil, err
	}

	cr := &SalesInvoiceResponse{}

	_, err = p.doRequest(req, cr)

	if err != nil {
		return nil, err
	}

	return cr, nil
}

// UpdateSalesInvoice ...
func (p *Parasut) UpdateSalesInvoice(id int, cur *SalesInvoiceUpdateRequest) (*SalesInvoiceResponse, error) {

	endpoint := fmt.Sprintf("sales_invoices/%d", id)

	url, err := prepareURL(endpoint, nil)

	if err != nil {
		return nil, err
	}

	req, err := p.newRequest("PUT", url, cur)

	if err != nil {
		return nil, err
	}

	cr := &SalesInvoiceResponse{}

	_, err = p.doRequest(req, cr)

	if err != nil {
		return nil, err
	}

	return cr, nil
}

// DeleteSalesInvoice ...
func (p *Parasut) DeleteSalesInvoice(id int) (*SalesInvoiceDeleteResponse, error) {

	endpoint := fmt.Sprintf("sales_invoices/%d", id)

	url, err := prepareURL(endpoint, nil)

	if err != nil {
		return nil, err
	}

	req, err := p.newRequest("DELETE", url, nil)

	if err != nil {
		return nil, err
	}

	cdr := &SalesInvoiceDeleteResponse{}

	_, err = p.doRequest(req, cdr)

	if err != nil {
		return nil, err
	}

	return cdr, nil
}
