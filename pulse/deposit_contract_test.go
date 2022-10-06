package pulse

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/triedb"
)

func TestReplaceDepositContract(t *testing.T) {
	// Init
	db := rawdb.NewMemoryDatabase()
	triedb := triedb.NewDatabase(db, &triedb.Config{Preimages: true})
	tdb := state.NewDatabase(triedb, nil)
	state, _ := state.New(common.Hash{}, tdb)

	// Exec
	replaceDepositContract(state)

	// Verify
	balance := state.GetBalance(pulseDepositContractAddr)
	if balance.Cmp(common.U2560) != 0 {
		t.Errorf("Found unexpected deposit contract balance: %d", balance)
	}

	actualCode := state.GetCode(pulseDepositContractAddr)
	for i, b := range actualCode {
		if b != depositContractBytes[i] {
			t.Errorf("Invalid deposit contract code at index %d", i)
		}
	}

	// Verify Storage
	for i, store := range depositContractStorage {
		actualStorage := state.GetState(pulseDepositContractAddr, common.HexToHash(store[0]))
		expectedStorage := common.HexToHash(store[1])
		if actualStorage != expectedStorage {
			t.Errorf("Invalid storage entry %d, actual: %d, expected: %d", i, actualStorage, expectedStorage)
		} else {
			t.Log("Valid Storage entry")
		}
	}
}
