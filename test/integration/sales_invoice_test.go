package integration

import (
	"net/http"
	"testing"

	"github.com/yuexcom/go-parasut/parasut"
)

const (
	sampleProductID  = "10562965"
	sampleCategoryID = "2610632"
	sampleContactID  = "13208345"
	sampleTagID      = "194913"
)

func TestSalesInvoiceCRUD(t *testing.T) {

	req := &parasut.SalesInvoiceRequest{
		Data: &parasut.SalesInvoiceData{
			Type: "sales_invoices",
			Attributes: &parasut.SalesInvoice{
				ItemType:    "invoice",
				IssueDate:   "2019-03-28",
				Description: "Lorem ipsum dolor sit amet.",
			},
			Relationships: &parasut.Relationships{
				"details": &parasut.Relationship{
					Data: []*parasut.SalesInvoiceDetailsData{
						&parasut.SalesInvoiceDetailsData{
							Type: "sales_invoice_details",
							Attributes: &parasut.SalesInvoiceDetails{
								Quantity:  1.0,
								UnitPrice: 5.0,
								VatRate:   18.0,
							},
							Relationships: &parasut.Relationships{
								"product": &parasut.Relationship{
									Data: &parasut.ProductData{
										ID:   sampleProductID,
										Type: "products",
									},
								},
							},
						},
					},
				},
				"category": &parasut.Relationship{
					Data: &parasut.CategoryData{
						ID:   sampleCategoryID,
						Type: "item_categories",
					},
				},
				"contact": &parasut.Relationship{
					Data: &parasut.ContactData{
						ID:   sampleContactID,
						Type: "contacts",
					},
				},
				"tags": &parasut.Relationship{
					Data: []*parasut.TagData{
						&parasut.TagData{
							ID:   sampleTagID,
							Type: "tags",
						},
					},
				},
			},
		},
	}

	// CREATE

	salesInvoice, _, err := client.SalesInvoices.Create(req)

	if err != nil {
		t.Errorf("create error, %s", err)
	}

	// LIST

	listOpts := &parasut.SalesInvoiceListOptions{
		Filter: &parasut.SalesInvoiceListFilter{
			IssueDate: "2019-03-28",
		},
	}

	salesInvoices, _, err := client.SalesInvoices.List(listOpts)

	if err != nil {
		t.Errorf("list error, %s", err)
	}

	if len(salesInvoices.Data) != 1 {
		t.Errorf("wrong record count, want %d got %d", 1, len(salesInvoices.Data))
	}

	// GET

	salesInvoice, _, err = client.SalesInvoices.Get(salesInvoice.Data.ID, nil)

	if err != nil {
		t.Errorf("get error, %s", err)
	}

	// UPDATE

	req = &parasut.SalesInvoiceRequest{
		Data: &parasut.SalesInvoiceData{
			Type: "sales_invoices",
			Attributes: &parasut.SalesInvoice{
				ItemType:    "invoice",
				IssueDate:   "2019-03-28",
				Description: "Lorem ipsum dolor sit fatura.",
			},
		},
	}

	salesInvoice, _, err = client.SalesInvoices.Update(salesInvoice.Data.ID, req)

	if err != nil {
		t.Errorf("update error, %s", err)
	}

	if salesInvoice.Data.Attributes.Description != "Lorem ipsum dolor sit fatura." {
		t.Errorf("update modify error")
	}

	// DELETE

	resp, err := client.SalesInvoices.Delete(salesInvoice.Data.ID)

	if err != nil {
		t.Errorf("delete error, %s", err)
	}

	if resp.StatusCode != http.StatusNoContent {
		t.Errorf("wrong http status code, want %d got %d", http.StatusNoContent, resp.StatusCode)
	}

	_, _, err = client.SalesInvoices.Get(salesInvoice.Data.ID, nil)

	if err == nil {
		t.Errorf("still exist after delete, id: %s", salesInvoice.Data.ID)
	}
}
