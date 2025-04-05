package main

import (
	"testing"

	flow "github.com/mzmbq/flow-launcher-go"
)

func TestFetchRandomMeal(t *testing.T) {
	ms, err := fetchRandomMeal()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if ms == nil {
		t.Fatal("expected a meal, got nil")
	}
	if len(ms.Meals) == 0 {
		t.Fatal("expected at least one meal")
	}
	if len(ms.Meals) != 1 {
		t.Fatal("expected exactly one meal")
	}
}

func TestSearchMealByName(t *testing.T) {
	ms, err := searchMealByName("pasta")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if ms == nil {
		t.Fatalf("expected meals, got nil")
	}
	if len(ms.Meals) == 0 {
		t.Fatal("expected at least one meal for 'pasta' search")
	}
}

func TestHandleQuery(t *testing.T) {
	tests := []struct {
		name       string
		parameters []string
		wantEmpty  bool
	}{
		{
			name:       "Random meal",
			parameters: []string{},
			wantEmpty:  false,
		},
		{
			name:       "Search meal",
			parameters: []string{"pasta"},
			wantEmpty:  false,
		},
		{
			name:       "Search bad",
			parameters: []string{"salkfhasopiufbaisf"},
			wantEmpty:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := &flow.Request{Parameters: tt.parameters}
			resp := handleQuery(req)
			if resp == nil {
				t.Fatal("expected response, got nil")
			}
			if !tt.wantEmpty && len(resp.Results) == 0 {
				t.Fatal("expected at least one result")
			}
			if tt.wantEmpty && len(resp.Results) != 0 {
				t.Fatal("expected empty result")
			}
		})
	}
}
