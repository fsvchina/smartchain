package types

import (
	protoio "github.com/gogo/protobuf/io"
)




type Snapshotter interface {

	Snapshot(height uint64, protoWriter protoio.Writer) error




	Restore(height uint64, format uint32, protoReader protoio.Reader) (SnapshotItem, error)
}



type ExtensionSnapshotter interface {
	Snapshotter


	SnapshotName() string




	SnapshotFormat() uint32


	SupportedFormats() []uint32
}
