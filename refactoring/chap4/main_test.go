package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

var sampleProvince = Province{
	Name: "Asia",
	Producers: []Producer{
		{"Byzantium", 10, 9},
		{"Attalia", 12, 10},
		{"Sinope", 10, 6},
	},
	Demand: 30,
	Price:  20,
}

var sampleProvinceWithoutProducers = Province{
	Name:      "Asia",
	Producers: []Producer{},
	Demand:    30,
	Price:     20,
}

func TestProvinceShortFall(t *testing.T) {
	testCases := []struct {
		name          string
		input         Province
		expectedValue int
		expectErr     bool
	}{
		{
			name:          "normal success",
			input:         sampleProvince,
			expectedValue: 5,
			expectErr:     false,
		},
		{
			name:          "success with no producers",
			input:         sampleProvinceWithoutProducers,
			expectedValue: 30,
			expectErr:     false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if diff := cmp.Diff(tc.expectedValue, tc.input.ShortFall()); diff != "" {
				t.Errorf("test result was wrong: %+v", diff)
			}
		})
	}
}

func TestProvinceProfit(t *testing.T) {
	testCases := []struct {
		name          string
		input         Province
		expectedValue int
		expectErr     bool
	}{
		{
			name:          "success",
			input:         sampleProvince,
			expectedValue: 500,
			expectErr:     false,
		},
		{
			name:          "success with no producers",
			input:         sampleProvinceWithoutProducers,
			expectedValue: 0,
			expectErr:     false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if diff := cmp.Diff(tc.expectedValue, tc.input.Profit()); diff != "" {
				t.Errorf("test result was wrong: %+v", diff)
			}
		})
	}
}
