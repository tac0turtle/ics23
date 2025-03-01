package ics23

import (
	"bytes"
	"testing"
)

func TestExistenceProof(t *testing.T) {
	cases := ExistenceProofTestData(t)

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			res, err := tc.Proof.Calculate()
			// short-circuit with error case
			if tc.IsErr && err == nil {
				t.Fatal("Expected error, but got none")
			}
			if tc.IsErr == false && err != nil {
				t.Fatal(err)
			}
			if !bytes.Equal(res, tc.Expected) {
				t.Errorf("Bad result: %s vs %s", toHex(res), toHex(tc.Expected))
			}
		})
	}
}

func TestCheckLeaf(t *testing.T) {
	cases := CheckLeafTestData(t)

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			err := tc.Leaf.CheckAgainstSpec(&ProofSpec{LeafSpec: tc.Spec})
			if tc.IsErr && err == nil {
				t.Fatal("Expected error, but got nil")
			} else if !tc.IsErr && err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}
		})
	}
}

func TestCheckAgainstSpec(t *testing.T) {
	cases := CheckAgainstSpecTestData(t)

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			err := tc.Proof.CheckAgainstSpec(tc.Spec)
			if tc.IsErr && err == nil {
				t.Fatal("Expected error, but got nil")
			} else if !tc.IsErr && err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}
		})
	}
}
