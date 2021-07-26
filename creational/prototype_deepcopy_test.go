package creational

import (
	"testing"
)

func TestOrderDeepCopy(t *testing.T) {
	order := Order{
		"M/S PN Associates - Order 001",
		1388.90,
		[]OrderItem{
			{"Note Books", 2.99, 100},
			{"Sticky Notes", 1.70, 250},
			{"HB Pencils - 10 pack", 4.00, 100},
		},
		&Address{"102 Station Rd", "Sydney", "2000"},
	}

	repeatOrder := order.DeepCopy()

	if repeatOrder.Name != order.Name || repeatOrder.OrderTotal != order.OrderTotal {
		t.Error("order elements could not be correctly copied")
	}

	if repeatOrder.ShippingAddress.Street != order.ShippingAddress.Street ||
		repeatOrder.ShippingAddress.City != order.ShippingAddress.City ||
		repeatOrder.ShippingAddress.PostCode != order.ShippingAddress.PostCode {
		t.Error("address could not be copied successfully")
	}

	for i, item := range repeatOrder.Items {
		if item.ItemName != order.Items[i].ItemName ||
			item.Price != order.Items[i].Price ||
			item.Quantity != order.Items[i].Quantity {
			t.Error("order Items could not be copied sucessfully")
		}
	}
}

func TestOrderDeepCopySerialization(t *testing.T) {
	order := Order{
		"M/S PN Associates - Order 001",
		1388.90,
		[]OrderItem{
			{"Note Books", 2.99, 100},
			{"Sticky Notes", 1.70, 250},
			{"HB Pencils - 10 pack", 4.00, 100},
		},
		&Address{"102 Station Rd", "Sydney", "2000"},
	}

	repeatOrder := order.DeepCopySerialization()

	if repeatOrder.Name != order.Name || repeatOrder.OrderTotal != order.OrderTotal {
		t.Error("order elements could not be correctly copied")
	}

	if repeatOrder.ShippingAddress.Street != order.ShippingAddress.Street ||
		repeatOrder.ShippingAddress.City != order.ShippingAddress.City ||
		repeatOrder.ShippingAddress.PostCode != order.ShippingAddress.PostCode {
		t.Error("address could not be copied successfully")
	}

	for i, item := range repeatOrder.Items {
		if item.ItemName != order.Items[i].ItemName ||
			item.Price != order.Items[i].Price ||
			item.Quantity != order.Items[i].Quantity {
			t.Error("order Items could not be copied sucessfully")
		}
	}
}
