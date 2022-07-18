package keeper_test

import "testing"

func BenchmarkGetValidator(b *testing.B) {


	var powersNumber = 900

	var totalPower int64 = 0
	var powers = make([]int64, powersNumber)
	for i := range powers {
		powers[i] = int64(i)
		totalPower += int64(i)
	}

	app, ctx, _, valAddrs, vals := initValidators(b, totalPower, len(powers), powers)

	for _, validator := range vals {
		app.StakingKeeper.SetValidator(ctx, validator)
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		for _, addr := range valAddrs {
			_, _ = app.StakingKeeper.GetValidator(ctx, addr)
		}
	}
}
