package simple_crypto

import (
	"testing"
)

func TestValidateProof(t *testing.T) {
	if Valid_proof(1, 2) {
		t.Fatal(`Valid_proof(1, 2) = true, wanted false, "", error`)
	}
	if !Valid_proof(7334, 89597) {
		t.Fatal(`Valid_proof(7334, 89597) = false, wanted true, "", error`)
	}
}

func TestProofOfWork(t *testing.T) {
	test1 := Proof_of_work(1)
	if test1 != 72608 {
		t.Fatalf(`Proof_of_work("1") = [%d], wanted 72608, "", error`, test1)
	}
	test2 := Proof_of_work(88483)
	if test2 != 33174 {
		t.Fatalf(`Proof_of_work("88483") = [%d], wanted 33174, "", error`, test2)
	}

}

func TestTransactions(t *testing.T) {

	New_Transaction("from_test", "to_test", 5)
	New_Transaction("from_test1", "to_test1", 10)

	next_proof := Proof_of_work(Last_Block().Proof)

	dta := Mine()

	if len(dta.Transactions) != 3 {
		t.Fatalf(`Mine("") transaction len = [%d], wanted 3, "", error`, len(dta.Transactions))
	}

	if next_proof != dta.Proof {
		t.Fatalf(`Mine("") proof = [%d], wanted [%d], "", error`, dta.Proof, next_proof)
	}

}

func TestNextSet(t *testing.T) {

	New_Transaction("from_test", "to_test", 900)

	next_proof := Proof_of_work(Last_Block().Proof)

	dta := Mine()

	if len(dta.Transactions) != 2 {
		t.Fatalf(`Mine("") transaction len = [%d], wanted 2, "", error`, len(dta.Transactions))
	}

	if next_proof != dta.Proof {
		t.Fatalf(`Mine("") proof = [%d], wanted [%d], "", error`, dta.Proof, next_proof)
	}
}
