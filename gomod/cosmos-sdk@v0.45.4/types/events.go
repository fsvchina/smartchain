package types

import (
	"encoding/json"
	"fmt"
	"reflect"
	"sort"
	"strings"

	"github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/cosmos/cosmos-sdk/codec"
)







type EventManager struct {
	events Events
}

func NewEventManager() *EventManager {
	return &EventManager{EmptyEvents()}
}

func (em *EventManager) Events() Events { return em.events }



func (em *EventManager) EmitEvent(event Event) {
	em.events = em.events.AppendEvent(event)
}



func (em *EventManager) EmitEvents(events Events) {
	em.events = em.events.AppendEvents(events)
}


func (em EventManager) ABCIEvents() []abci.Event {
	return em.events.ToABCIEvents()
}


func (em *EventManager) EmitTypedEvent(tev proto.Message) error {
	event, err := TypedEventToEvent(tev)
	if err != nil {
		return err
	}

	em.EmitEvent(event)
	return nil
}


func (em *EventManager) EmitTypedEvents(tevs ...proto.Message) error {
	events := make(Events, len(tevs))
	for i, tev := range tevs {
		res, err := TypedEventToEvent(tev)
		if err != nil {
			return err
		}
		events[i] = res
	}

	em.EmitEvents(events)
	return nil
}


func TypedEventToEvent(tev proto.Message) (Event, error) {
	evtType := proto.MessageName(tev)
	evtJSON, err := codec.ProtoMarshalJSON(tev, nil)
	if err != nil {
		return Event{}, err
	}

	var attrMap map[string]json.RawMessage
	err = json.Unmarshal(evtJSON, &attrMap)
	if err != nil {
		return Event{}, err
	}

	attrs := make([]abci.EventAttribute, 0, len(attrMap))
	for k, v := range attrMap {
		attrs = append(attrs, abci.EventAttribute{
			Key:   []byte(k),
			Value: v,
		})
	}

	return Event{
		Type:       evtType,
		Attributes: attrs,
	}, nil
}


func ParseTypedEvent(event abci.Event) (proto.Message, error) {
	concreteGoType := proto.MessageType(event.Type)
	if concreteGoType == nil {
		return nil, fmt.Errorf("failed to retrieve the message of type %q", event.Type)
	}

	var value reflect.Value
	if concreteGoType.Kind() == reflect.Ptr {
		value = reflect.New(concreteGoType.Elem())
	} else {
		value = reflect.Zero(concreteGoType)
	}

	protoMsg, ok := value.Interface().(proto.Message)
	if !ok {
		return nil, fmt.Errorf("%q does not implement proto.Message", event.Type)
	}

	attrMap := make(map[string]json.RawMessage)
	for _, attr := range event.Attributes {
		attrMap[string(attr.Key)] = attr.Value
	}

	attrBytes, err := json.Marshal(attrMap)
	if err != nil {
		return nil, err
	}

	err = jsonpb.Unmarshal(strings.NewReader(string(attrBytes)), protoMsg)
	if err != nil {
		return nil, err
	}

	return protoMsg, nil
}





type (

	Event abci.Event


	Events []Event
)



func NewEvent(ty string, attrs ...Attribute) Event {
	e := Event{Type: ty}

	for _, attr := range attrs {
		e.Attributes = append(e.Attributes, attr.ToKVPair())
	}

	return e
}


func NewAttribute(k, v string) Attribute {
	return Attribute{k, v}
}


func EmptyEvents() Events {
	return make(Events, 0)
}

func (a Attribute) String() string {
	return fmt.Sprintf("%s: %s", a.Key, a.Value)
}


func (a Attribute) ToKVPair() abci.EventAttribute {
	return abci.EventAttribute{Key: toBytes(a.Key), Value: toBytes(a.Value)}
}


func (e Event) AppendAttributes(attrs ...Attribute) Event {
	for _, attr := range attrs {
		e.Attributes = append(e.Attributes, attr.ToKVPair())
	}
	return e
}


func (e Events) AppendEvent(event Event) Events {
	return append(e, event)
}


func (e Events) AppendEvents(events Events) Events {
	return append(e, events...)
}



func (e Events) ToABCIEvents() []abci.Event {
	res := make([]abci.Event, len(e))
	for i, ev := range e {
		res[i] = abci.Event{Type: ev.Type, Attributes: ev.Attributes}
	}

	return res
}

func toBytes(i interface{}) []byte {
	switch x := i.(type) {
	case []uint8:
		return x
	case string:
		return []byte(x)
	default:
		panic(i)
	}
}


var (
	EventTypeTx = "tx"

	AttributeKeyAccountSequence = "acc_seq"
	AttributeKeySignature       = "signature"
	AttributeKeyFee             = "fee"

	EventTypeMessage = "message"

	AttributeKeyAction = "action"
	AttributeKeyModule = "module"
	AttributeKeySender = "sender"
	AttributeKeyAmount = "amount"
)

type (

	StringEvents []StringEvent
)

func (se StringEvents) String() string {
	var sb strings.Builder

	for _, e := range se {
		sb.WriteString(fmt.Sprintf("\t\t- %s\n", e.Type))

		for _, attr := range e.Attributes {
			sb.WriteString(fmt.Sprintf("\t\t\t- %s\n", attr.String()))
		}
	}

	return strings.TrimRight(sb.String(), "\n")
}



func (se StringEvents) Flatten() StringEvents {
	flatEvents := make(map[string][]Attribute)

	for _, e := range se {
		flatEvents[e.Type] = append(flatEvents[e.Type], e.Attributes...)
	}
	keys := make([]string, 0, len(flatEvents))
	res := make(StringEvents, 0, len(flatEvents))

	for ty := range flatEvents {
		keys = append(keys, ty)
	}

	sort.Strings(keys)
	for _, ty := range keys {
		res = append(res, StringEvent{Type: ty, Attributes: flatEvents[ty]})
	}

	return res
}


func StringifyEvent(e abci.Event) StringEvent {
	res := StringEvent{Type: e.Type}

	for _, attr := range e.Attributes {
		res.Attributes = append(
			res.Attributes,
			Attribute{string(attr.Key), string(attr.Value)},
		)
	}

	return res
}



func StringifyEvents(events []abci.Event) StringEvents {
	res := make(StringEvents, 0, len(events))

	for _, e := range events {
		res = append(res, StringifyEvent(e))
	}

	return res.Flatten()
}



func MarkEventsToIndex(events []abci.Event, indexSet map[string]struct{}) []abci.Event {
	indexAll := len(indexSet) == 0
	updatedEvents := make([]abci.Event, len(events))

	for i, e := range events {
		updatedEvent := abci.Event{
			Type:       e.Type,
			Attributes: make([]abci.EventAttribute, len(e.Attributes)),
		}

		for j, attr := range e.Attributes {
			_, index := indexSet[fmt.Sprintf("%s.%s", e.Type, attr.Key)]
			updatedAttr := abci.EventAttribute{
				Key:   attr.Key,
				Value: attr.Value,
				Index: index || indexAll,
			}

			updatedEvent.Attributes[j] = updatedAttr
		}

		updatedEvents[i] = updatedEvent
	}

	return updatedEvents
}
