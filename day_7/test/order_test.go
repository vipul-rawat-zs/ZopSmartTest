package main

import (
	"reflect"
	"testing"
)

type input struct {
	id     string
	items  []string
	amount float64
}

func compOrders(o1, o2 *order) bool {
	if o1.id != o2.id {
		return false
	}
	if compStringSlice(o1.items, o2.items) {
		return false
	}
	if o1.amount != o2.amount {
		return false
	}
	if o1.includesAlcohol != o2.includesAlcohol {
		return false
	}
	if o1.shippingFee != o2.shippingFee {
		return false
	}
	return true
}

func compStringSlice(s1, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

func TestNew(t *testing.T) {
	cases := []struct {
		input  input
		output *order
	}{
		{input{"1022", []string{"item1", "item2"}, 100}, &order{1022, []string{"item1", "item2"}, 0, false, 100.0}},
		{input{"1021", []string{"itemx", "itemy"}, 250}, &order{1021, []string{"itemx", "itemy"}, 0, false, 250.0}},
		{input{"12", []string{"item1", "item2", "item3"}, 10}, &order{10, []string{"item1", "item2", "item3"}, 0, false, 10.0}},
	}

	for _, v := range cases {
		t.Run("Running test on New", func(t *testing.T) {
			out := New(v.input.id, v.input.items, v.input.amount)
			// if compOrders(out, v.output) {
			if !reflect.DeepEqual(out, v.output) {
				t.Errorf("expected %v but got %v", v.output, out)
			}
		})
	}
}

func TestRecalculate(t *testing.T) {
	cases := []struct {
		input  order
		output float64
	}{
		{order{10, []string{"item"}, 0, false, 240}, 10},
		{order{11, []string{"item"}, 0, false, 250}, 0},
	}

	for _, v := range cases {
		t.Run("Running test on Recalculate", func(t *testing.T) {
			v.input.Recalculate()
			if v.input.shippingFee != v.output {
				t.Errorf("expected %v, but got %v", v.output, v.input.shippingFee)
			}
		})
	}
}
