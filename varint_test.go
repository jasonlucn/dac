package dac

import (
	"encoding/binary"
	"fmt"
	"testing"
)

func TestEncodeVarint(t *testing.T) {
	x := uint64(123456)
	xbytes := EncodeVarint(x)
	fmt.Println(len(xbytes))
	fmt.Println(DecodeVarint(xbytes))
	
	varBuf := make([]byte, binary.MaxVarintLen64)
	wb := binary.PutUvarint(varBuf, 123456)
	fmt.Println(varBuf)
	fmt.Println(len(varBuf))
	fmt.Println(wb)
}
