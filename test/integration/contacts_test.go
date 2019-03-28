package integration

import (
	"net/http"
	"testing"

	"github.com/yuexcom/go-parasut/parasut"
)

func TestContactCRUD(t *testing.T) {

	req := &parasut.ContactRequest{
		Data: &parasut.ContactData{
			Type: "contacts",
			Attributes: &parasut.Contact{
				Name:        "John Doe Ltd. Şti.",
				ShortName:   "JD Ltd. Şti.",
				AccountType: "customer",
				TaxNumber:   "9999999999",
			},
			Relationships: &parasut.Relationships{
				"contact_people": &parasut.Relationship{
					Data: []*parasut.ContactPersonData{
						&parasut.ContactPersonData{
							Type: "contact_people",
							Attributes: &parasut.ContactPerson{
								Name:  "John Doe",
								Email: "john@doe.com",
								Phone: "111 222 33 44",
								Notes: "Lorem ipsum dolor sit amet.",
							},
						},
					},
				},
				// "category": &parasut.Relationship{ // TODO: bu eklenince "Müşteri/tedarikçi grubu tipi geçersiz." hatası dönüyor.
				// 	Data: &parasut.CategoryData{
				// 		ID:   "2610632",
				// 		Type: "item_categories",
				// 	},
				// },
			},
		},
	}

	// CREATE

	contact, _, err := client.Contacts.Create(req)

	if err != nil {
		t.Errorf("create error, %s", err)
	}

	// LIST

	listOpts := &parasut.ContactListOptions{
		Filter: &parasut.ContactListFilter{
			TaxNumber: "9999999999",
		},
	}

	contacts, _, err := client.Contacts.List(listOpts)

	if err != nil {
		t.Errorf("list error, %s", err)
	}

	if len(contacts.Data) != 1 {
		t.Errorf("wrong record count, want %d got %d", 1, len(contacts.Data))
	}

	// GET

	contact, _, err = client.Contacts.Get(contact.Data.ID, nil)

	if err != nil {
		t.Errorf("get error, %s", err)
	}

	// UPDATE

	req = &parasut.ContactRequest{
		Data: &parasut.ContactData{
			ID:   contact.Data.ID,
			Type: "contacts",
			Attributes: &parasut.Contact{
				Name:        "John Doe II Ltd. Şti.",
				ShortName:   "JD2 Ltd. Şti.",
				AccountType: "customer",
				TaxNumber:   "0987654321",
			},
		},
	}

	contact, _, err = client.Contacts.Update(contact.Data.ID, req)

	if err != nil {
		t.Errorf("update error, %s", err)
	}

	if contact.Data.Attributes.Name != "John Doe II Ltd. Şti." {
		t.Errorf("update modify error")
	}

	// DELETE

	resp, err := client.Contacts.Delete(contact.Data.ID)

	if err != nil {
		t.Errorf("delete error, %s", err)
	}

	if resp.StatusCode != http.StatusNoContent {
		t.Errorf("wrong http status code, want %d got %d", http.StatusNoContent, resp.StatusCode)
	}

	_, _, err = client.Contacts.Get(contact.Data.ID, nil)

	if err == nil {
		t.Errorf("still exist after delete, id: %s", contact.Data.ID)
	}
}
