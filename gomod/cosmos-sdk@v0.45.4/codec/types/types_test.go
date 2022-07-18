package types_test

import (
	"strings"
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/gogo/protobuf/proto"
	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
)

func TestAnyPackUnpack(t *testing.T) {
	registry := testdata.NewTestInterfaceRegistry()

	spot := &testdata.Dog{Name: "Spot"}
	var animal testdata.Animal


	any, err := types.NewAnyWithValue(spot)
	require.NoError(t, err)
	require.Equal(t, spot, any.GetCachedValue())
	err = registry.UnpackAny(any, &animal)
	require.NoError(t, err)
	require.Equal(t, spot, animal)
}

type TestI interface {
	DoSomething()
}



type FakeDog struct{}

var (
	_ proto.Message   = &FakeDog{}
	_ testdata.Animal = &FakeDog{}
)


func (dog FakeDog) Reset()                  {}
func (dog FakeDog) String() string          { return "fakedog" }
func (dog FakeDog) ProtoMessage()           {}
func (dog FakeDog) XXX_MessageName() string { return proto.MessageName(&testdata.Dog{}) }
func (dog FakeDog) Greet() string           { return "fakedog" }

func TestRegister(t *testing.T) {
	registry := types.NewInterfaceRegistry()
	registry.RegisterInterface("Animal", (*testdata.Animal)(nil))
	registry.RegisterInterface("TestI", (*TestI)(nil))


	require.NotPanics(t, func() {
		registry.RegisterImplementations((*testdata.Animal)(nil), &testdata.Dog{})
	})


	require.Panics(t, func() {
		registry.RegisterImplementations((*TestI)(nil), &testdata.Dog{})
	})


	require.Panics(t, func() {
		registry.RegisterImplementations((*TestI)(nil), nil)
	})


	require.Panics(t, func() {
		registry.RegisterInterface("not_an_interface", (*testdata.Dog)(nil))
	})


	require.NotPanics(t, func() {
		registry.RegisterImplementations((*testdata.Animal)(nil), &testdata.Dog{})
	})


	require.PanicsWithError(
		t,
		"concrete type *testdata.Dog has already been registered under typeURL /testdata.Dog, cannot register *types_test.FakeDog under same typeURL. "+
			"This usually means that there are conflicting modules registering different concrete types for a same interface implementation",
		func() {
			registry.RegisterImplementations((*testdata.Animal)(nil), &FakeDog{})
		},
	)
}

func TestUnpackInterfaces(t *testing.T) {
	registry := testdata.NewTestInterfaceRegistry()

	spot := &testdata.Dog{Name: "Spot"}
	any, err := types.NewAnyWithValue(spot)
	require.NoError(t, err)

	hasAny := testdata.HasAnimal{
		Animal: any,
		X:      1,
	}
	bz, err := hasAny.Marshal()
	require.NoError(t, err)

	var hasAny2 testdata.HasAnimal
	err = hasAny2.Unmarshal(bz)
	require.NoError(t, err)

	err = types.UnpackInterfaces(hasAny2, registry)
	require.NoError(t, err)

	require.Equal(t, spot, hasAny2.Animal.GetCachedValue())
}

func TestNested(t *testing.T) {
	registry := testdata.NewTestInterfaceRegistry()

	spot := &testdata.Dog{Name: "Spot"}
	any, err := types.NewAnyWithValue(spot)
	require.NoError(t, err)

	ha := &testdata.HasAnimal{Animal: any}
	any2, err := types.NewAnyWithValue(ha)
	require.NoError(t, err)

	hha := &testdata.HasHasAnimal{HasAnimal: any2}
	any3, err := types.NewAnyWithValue(hha)
	require.NoError(t, err)

	hhha := testdata.HasHasHasAnimal{HasHasAnimal: any3}


	bz, err := hhha.Marshal()
	require.NoError(t, err)


	var hhha2 testdata.HasHasHasAnimal
	err = hhha2.Unmarshal(bz)
	require.NoError(t, err)
	err = types.UnpackInterfaces(hhha2, registry)
	require.NoError(t, err)

	require.Equal(t, spot, hhha2.TheHasHasAnimal().TheHasAnimal().TheAnimal())
}

func TestAny_ProtoJSON(t *testing.T) {
	spot := &testdata.Dog{Name: "Spot"}
	any, err := types.NewAnyWithValue(spot)
	require.NoError(t, err)

	jm := &jsonpb.Marshaler{}
	json, err := jm.MarshalToString(any)
	require.NoError(t, err)
	require.Equal(t, "{\"@type\":\"/testdata.Dog\",\"name\":\"Spot\"}", json)

	registry := testdata.NewTestInterfaceRegistry()
	jum := &jsonpb.Unmarshaler{}
	var any2 types.Any
	err = jum.Unmarshal(strings.NewReader(json), &any2)
	require.NoError(t, err)
	var animal testdata.Animal
	err = registry.UnpackAny(&any2, &animal)
	require.NoError(t, err)
	require.Equal(t, spot, animal)

	ha := &testdata.HasAnimal{
		Animal: any,
	}
	err = ha.UnpackInterfaces(types.ProtoJSONPacker{JSONPBMarshaler: jm})
	require.NoError(t, err)
	json, err = jm.MarshalToString(ha)
	require.NoError(t, err)
	require.Equal(t, "{\"animal\":{\"@type\":\"/testdata.Dog\",\"name\":\"Spot\"}}", json)

	require.NoError(t, err)
	var ha2 testdata.HasAnimal
	err = jum.Unmarshal(strings.NewReader(json), &ha2)
	require.NoError(t, err)
	err = ha2.UnpackInterfaces(registry)
	require.NoError(t, err)
	require.Equal(t, spot, ha2.Animal.GetCachedValue())
}
