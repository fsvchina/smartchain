package config_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/config"
	"github.com/cosmos/cosmos-sdk/client/flags"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/cosmos/cosmos-sdk/x/staking/client/cli"
)

const (
	nodeEnv   = "NODE"
	testNode1 = "http:
	testNode2 = "http:
)


func initClientContext(t *testing.T, envVar string) (client.Context, func()) {
	home := t.TempDir()
	clientCtx := client.Context{}.
		WithHomeDir(home).
		WithViper("")

	clientCtx.Viper.BindEnv(nodeEnv)
	if envVar != "" {
		os.Setenv(nodeEnv, envVar)
	}

	clientCtx, err := config.ReadFromClientConfig(clientCtx)
	require.NoError(t, err)

	return clientCtx, func() { _ = os.RemoveAll(home) }
}

func TestConfigCmd(t *testing.T) {
	clientCtx, cleanup := initClientContext(t, testNode1)
	defer func() {
		os.Unsetenv(nodeEnv)
		cleanup()
	}()


	cmd := config.Cmd()
	args := []string{"node", testNode2}
	_, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, args)
	require.NoError(t, err)


	b := bytes.NewBufferString("")
	cmd.SetOut(b)
	cmd.SetArgs([]string{"node"})
	cmd.Execute()
	out, err := ioutil.ReadAll(b)
	require.NoError(t, err)
	require.Equal(t, string(out), testNode1+"\n")
}

func TestConfigCmdEnvFlag(t *testing.T) {
	const (
		defaultNode = "http:
	)

	tt := []struct {
		name    string
		envVar  string
		args    []string
		expNode string
	}{
		{"env var is set with no flag", testNode1, []string{"validators"}, testNode1},
		{"env var is set with a flag", testNode1, []string{"validators", fmt.Sprintf("--%s=%s", flags.FlagNode, testNode2)}, testNode2},
		{"env var is not set with no flag", "", []string{"validators"}, defaultNode},
		{"env var is not set with a flag", "", []string{"validators", fmt.Sprintf("--%s=%s", flags.FlagNode, testNode2)}, testNode2},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			clientCtx, cleanup := initClientContext(t, tc.envVar)
			defer func() {
				if tc.envVar != "" {
					os.Unsetenv(nodeEnv)
				}
				cleanup()
			}()
			
			cmd := cli.GetQueryCmd()
			_, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)
			require.Error(t, err)
			require.Contains(t, err.Error(), tc.expNode, "Output does not contain expected Node")
		})
	}
}
