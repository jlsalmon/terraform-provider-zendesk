package zendesk

import (
	"fmt"
	. "github.com/golang/mock/gomock"
	"github.com/nukosuke/go-zendesk/zendesk"
	"github.com/nukosuke/go-zendesk/zendesk/mock"
	"testing"
)

var testBrand = zendesk.Brand{
	ID:              47,
	URL:             "https://company.zendesk.com/api/v2/brands/47.json",
	Name:            "Brand 1",
	BrandURL:        "https://brand1.com",
	HasHelpCenter:   true,
	HelpCenterState: "enabled",
	Active:          true,
	Default:         true,
	Logo: zendesk.Attachment{
		ID:          928374,
		FileName:    "brand1_logo.png",
		ContentURL:  "https://company.zendesk.com/logos/brand1_logo.png",
		ContentType: "image/png",
		Size:        166144,
	},
	Subdomain:         "brand1",
	HostMapping:       "brand1.com",
	SignatureTemplate: "{{agent.signature}}",
}

func TestCreateBrand(t *testing.T) {
	ctrl := NewController(t)
	defer ctrl.Finish()

	m := mock.NewClient(ctrl)

	m.EXPECT().CreateBrand(Any()).Return(testBrand, nil)

	i := newIdentifiableGetterSetter()
	err := createBrand(i, m)
	if err != nil {
		t.Fatalf("Create brand returned an error %v", err)
	}

	if i.Id() != fmt.Sprintf("%d", testBrand.ID) {
		t.Fatalf("Created object does not have the correct brand id. Was: %s. Expected %d", i.Id(), testBrand.ID)
	}

	if i.Get("logo_attachment_id") != testBrand.Logo.ID {
		t.Fatalf("Created object does not have the correct logo id. Was: %d. Expected %d", i.Get("logo_attachment_id"), testBrand.Logo.ID)
	}
}

func TestReadBrand(t *testing.T) {
	ctrl := NewController(t)
	defer ctrl.Finish()

	m := mock.NewClient(ctrl)
	m.EXPECT().GetBrand(testBrand.ID).Return(testBrand, nil)
	i := newIdentifiableGetterSetter()
	i.SetId(fmt.Sprintf("%d", testBrand.ID))

	err := readBrand(i, m)
	if err != nil {
		t.Fatalf("readBrand returned an error: %v", err)
	}

	if v := i.Get("subdomain"); v != testBrand.Subdomain {
		t.Fatalf("Subdomain was not set to the expected value. Was: %s Expected %s", v, testBrand.Subdomain)
	}
}

func TestUpdateBrand(t *testing.T) {
	updatedBrand := testBrand
	updatedBrand.Name = "1234"

	i := newIdentifiableGetterSetter()
	i.SetId(fmt.Sprintf("%d", testBrand.ID))

	ctrl := NewController(t)
	defer ctrl.Finish()

	m := mock.NewClient(ctrl)
	m.EXPECT().UpdateBrand(testBrand.ID, Any()).Return(updatedBrand, nil)

	err := updateBrand(i, m)
	if err != nil {
		t.Fatalf("update brand returned an error: %v", err)
	}

	if v := i.Get("name"); v != updatedBrand.Name {
		t.Fatalf("Update did not set name to the expected value. Was %s expected %s", v, updatedBrand.Name)
	}
}

func TestDeleteBrand(t *testing.T) {
	id := int64(1234)
	i := newIdentifiableGetterSetter()
	i.SetId(fmt.Sprintf("%d", id))

	ctrl := NewController(t)
	defer ctrl.Finish()

	m := mock.NewClient(ctrl)
	m.EXPECT().DeleteBrand(id).Return(nil)

	err := deleteBrand(i, m)
	if err != nil {
		t.Fatalf("delete brand returned an error: %v", err)
	}
}
