package types

import (

	_ "embed"
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi"
)


type HexString []byte


func (s HexString) MarshalJSON() ([]byte, error) {
	return json.Marshal(fmt.Sprintf("%x", string(s)))
}


func (s *HexString) UnmarshalJSON(data []byte) error {
	var x string
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	str, err := hex.DecodeString(x)
	if err != nil {
		return err
	}
	*s = str
	return nil
}


type CompiledContract struct {
	ABI abi.ABI
	Bin HexString
}

type jsonCompiledContract struct {
	ABI string
	Bin HexString
}


func (s CompiledContract) MarshalJSON() ([]byte, error) {
	abi1, err := json.Marshal(s.ABI)
	if err != nil {
		return nil, err
	}
	return json.Marshal(jsonCompiledContract{ABI: string(abi1), Bin: s.Bin})
}


func (s *CompiledContract) UnmarshalJSON(data []byte) error {
	var x jsonCompiledContract
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}

	s.Bin = x.Bin
	if err := json.Unmarshal([]byte(x.ABI), &s.ABI); err != nil {
		return fmt.Errorf("failed to unmarshal ABI: %w", err)
	}

	return nil
}

var (

	erc20JSON []byte


	ERC20Contract CompiledContract


	simpleStorageJSON []byte


	SimpleStorageContract CompiledContract


	testMessageCallJSON []byte


	TestMessageCall CompiledContract
)

func init() {
	err := json.Unmarshal(erc20JSON, &ERC20Contract)
	if err != nil {
		panic(err)
	}

	if len(ERC20Contract.Bin) == 0 {
		panic("load contract failed")
	}

	err = json.Unmarshal(testMessageCallJSON, &TestMessageCall)
	if err != nil {
		panic(err)
	}

	if len(TestMessageCall.Bin) == 0 {
		panic("load contract failed")
	}

	err = json.Unmarshal(simpleStorageJSON, &SimpleStorageContract)
	if err != nil {
		panic(err)
	}

	if len(TestMessageCall.Bin) == 0 {
		panic("load contract failed")
	}
}
