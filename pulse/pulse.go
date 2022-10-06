// Package pulse implements the PulseChain fork
package pulse

import (
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/params"
)

// Apply PrimordialPulse fork changes
func PrimordialPulseFork(state vm.StateDB, treasury *params.Treasury) {
	applySacrificeCredits(state, treasury)
	replaceDepositContract(state)
}
