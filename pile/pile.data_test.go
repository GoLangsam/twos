// Copyright 2018 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pile

import (
	"time"
)

func testIndex() (ID, []Index) {
	name := ID(TypeOf(Index(417756)).Name())
	data := []Index{ 4, 7, 1, 1, 5, 6} // Note: One duplicate!
	return name, data
}

func testCardinality() (ID, []Cardinality) {
	name := ID(TypeOf(Cardinality(417756)).Name())
	data := []Cardinality{ 4, 7, 1, 1, 5, 6} // Note: One duplicate!
	return name, data
}

func testID() (ID, []ID) {
	name := ID(TypeOf(ID("test")).Name())
	data := []ID{ "4", "7", "1", "1", "5", "6"} // Note: One duplicate!
	return name, data
}

func testInterface() (ID, []interface{}) {
	name := ID("interface{}")
	data := []interface{}{ int(4), Cardinality(7), time.Monday, time.Monday, "Thursday", ID("Friday") }
	return name, data
}

func testType() (ID, []Type) {
	name := ID("Type")
	data := []Type{TypeOf(int(4)), TypeOf(Cardinality(7)), TypeOf(time.Monday) , TypeOf(time.Monday), TypeOf("Thursday"), TypeOf(ID("Friday")) }
	// Note: One duplicate!
	return name, data
}

func testPile() (ID, []Pile) {
	name := ID("Pile")
	pile := NewPileOfID("time.Monday", "time.Monday")
	data := []Pile{NewPileOfIndex("the fourth one", Index(4)), NewPileOfCardinality("Seven", Cardinality(7)), pile, pile, NewPileOfID("Thursday", "Thursday"), NewPileOfID("Friday", "Friday") } //
	// Note: One duplicate!
	return name, data
}

func testPair() (ID, []Pair) {
	name := ID("Pair")
	head1, _ := NewPileOfIndex("the fourth one", Index(4)).Tail()()
	head2, _ := NewPileOfCardinality("Seven", Cardinality(7)).Tail()()
	head3, _ := NewPileOfID("time.Monday", "time.Monday").Tail()()
	head4, _ := NewPileOfID("time.Monday", "time.Monday").Tail()()
	head5, _ := NewPileOfID("Thursday", "Thursday").Tail()()
	head6, _ := NewPileOfID("Friday", "Friday").Tail()()
	_ = head4 // Cheap Cheat: use head3 twice to get ==
	data := []Pair{head1(),head2() ,head3() ,head3() ,head5() ,head6() } //
	// Note: One duplicate!
	return name, data
}

