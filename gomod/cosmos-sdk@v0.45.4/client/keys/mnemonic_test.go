package keys

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/testutil"
)

func Test_RunMnemonicCmdNormal(t *testing.T) {
	cmd := MnemonicKeyCommand()
	_ = testutil.ApplyMockIODiscardOutErr(cmd)
	cmd.SetArgs([]string{})
	require.NoError(t, cmd.Execute())
}

func Test_RunMnemonicCmdUser(t *testing.T) {
	cmd := MnemonicKeyCommand()
	_ = testutil.ApplyMockIODiscardOutErr(cmd)

	cmd.SetArgs([]string{fmt.Sprintf("--%s=1", flagUserEntropy)})
	err := cmd.Execute()
	require.Error(t, err)
	require.Equal(t, "EOF", err.Error())


	mockIn := testutil.ApplyMockIODiscardOutErr(cmd)
	mockIn.Reset("Hi!\n")
	err = cmd.Execute()
	require.Error(t, err)
	require.Equal(t,
		"256-bits is 43 characters in Base-64, and 100 in Base-6. You entered 3, and probably want more",
		err.Error())


	fakeEntropy := strings.Repeat(":)", 40) + "\ny\n"
	mockIn.Reset(fakeEntropy)
	require.NoError(t, cmd.Execute())


	fakeEntropy = strings.Repeat(":)", 40) + "\n"
	mockIn.Reset(fakeEntropy)
	require.Error(t, cmd.Execute())


	fakeEntropy = strings.Repeat(":)", 40) + "\nn\n"
	mockIn.Reset(fakeEntropy)
	require.NoError(t, cmd.Execute())
}
