package parasut

import (
	"fmt"
	"net/http"
)

// SalesInvoicesService ...
type SalesInvoicesService service

// List ...
func (s *SalesInvoicesService) List(filter *SalesInvoiceListFilter, opts *SalesInvoiceListOptions) (*SalesInvoicesResponse, *http.Response, error) {

	params := salesInvoiceListParams{
		Filter: filter,
		Sort:   opts.Sort,
		Page: &PageOptions{
			Number: opts.PageNumber,
			Size:   opts.PageSize,
		},
		Include: opts.Include,
	}

	url, err := addOptions("sales_invoices", params)

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
func (s *SalesInvoicesService) Get(id int, opts *SalesInvoiceListOptions) (*SalesInvoiceResponse, *http.Response, error) {

	params := salesInvoiceListParams{
		Include: opts.Include,
	}

	url := fmt.Sprintf("sales_invoices/%d", id)

	url, err := addOptions(url, params)

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
func (s *SalesInvoicesService) Create(si *SalesInvoice) (*SalesInvoiceResponse, *http.Response, error) {

	obj := s.prepareObj(si)

	url := "sales_invoices"

	req, err := s.client.NewRequest("POST", url, obj)

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
func (s *SalesInvoicesService) Update(id int, si *SalesInvoice) (*SalesInvoiceResponse, *http.Response, error) {

	obj := s.prepareObj(si)

	url := fmt.Sprintf("sales_invoices/%d", id)

	req, err := s.client.NewRequest("PUT", url, obj)

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
func (s *SalesInvoicesService) Delete(id int) (*http.Response, error) {

	url := fmt.Sprintf("sales_invoices/%d", id)

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

func (s *SalesInvoicesService) prepareObj(si *SalesInvoice) *salesInvoiceRequest {

	var (
		salesInvoicesDetailsData []*salesInvoiceDetailsData
		tagsData                 []*tagData
	)

	for _, sid := range si.SalesInvoiceDetails {
		tmp := &salesInvoiceDetailsData{
			Type: "sales_invoice_details",
			Attributes: &salesInvoiceDetails{
				Quantity:  sid.Quantity,
				UnitPrice: sid.UnitPrice,
				VatRate:   sid.VatRate,
			},
			Relationships: &relationships{
				"product": &relationship{
					Data: &productData{
						ID:   sid.ProductID,
						Type: "products",
					},
				},
			},
		}
		salesInvoicesDetailsData = append(salesInvoicesDetailsData, tmp)
	}

	for _, tagID := range si.TagsID {
		tmp := &tagData{
			ID:   tagID,
			Type: "tags",
		}
		tagsData = append(tagsData, tmp)
	}

	obj := &salesInvoiceRequest{
		Data: &salesInvoiceData{
			Type: "sales_invoices",
			Attributes: &salesInvoice{
				ItemType:    si.ItemType,
				IssueDate:   si.IssueDate,
				DueDate:     si.DueDate,
				Description: si.Description,
			},
			Relationships: &relationships{
				"details": &relationship{
					Data: salesInvoicesDetailsData,
				},
				"contact": &relationship{
					Data: contactData{
						ID:   si.ContactID,
						Type: "contacts",
					},
				},
				"category": &relationship{
					Data: categoryData{
						ID:   si.CategoryID,
						Type: "item_categories",
					},
				},
				"tags": &relationship{
					Data: tagsData,
				},
			},
		},
	}

	return obj
}
