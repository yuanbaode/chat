package models

import (
	"testing"
	"fmt"
)

func TestUnmarshalInfo(t *testing.T) {

	s := `第2408548期：\n大双/5`
	i := InfoStored{}
	UnmarshalInfo([]byte(s), &i)

	fmt.Println(i)

}
