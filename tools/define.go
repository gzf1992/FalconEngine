package tools

import (
	"encoding/binary"
	"fmt"
)

type FalconSearchEncoder interface {
	FalconEncoding() ([]byte, error)
	ToString() string
}

type FalconSearchDecoder interface {
	FalconDecoding(bytes []byte) error
	ToString() string
	// FalconPrepare(length int64) (int64,error)
}

type FalconCoder interface {
	FalconEncoding() ([]byte, error)
	FalconDecoding(bytes []byte) error
	ToString() string
}


type DictValue struct {
	Val uint64
	ExtVal uint64
}

func NewDicValue() *DictValue{
	return &DictValue{}
}

func (dv *DictValue) FalconEncoding() ([]byte,error) {
	b:=make([]byte,24)
	binary.LittleEndian.PutUint64(b[:8],uint64(16))
	binary.LittleEndian.PutUint64(b[8:16],dv.Val)
	binary.LittleEndian.PutUint64(b[16:],dv.ExtVal)
	return b,nil

}

func (dv *DictValue) FalconDecoding(src []byte) error {
	if len(src)!=24{
		return fmt.Errorf("Length is not 24 byte")
	}
	dv.Val=binary.LittleEndian.Uint64(src[8:16])
	dv.ExtVal=binary.LittleEndian.Uint64(src[16:])
	return nil
}

func (dv *DictValue) ToString() string {
	return fmt.Sprintf(`{ "Val": %d , "ExtVal"：%d }`,dv.Val,dv.ExtVal)
}



type DocId struct{
	DocID uint32
	Weight uint32
}

func (di *DocId) ToString() string {
	return fmt.Sprintf(`{"id":%d,"weight":%d}`,di.DocID,di.Weight)
}


type Document struct {

}