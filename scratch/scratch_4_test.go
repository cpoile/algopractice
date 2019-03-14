package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// the base structs
var baseStructEmbed2A, baseStructEmbed2B, baseStructEmbed2C, baseStructEmbed2D testStructEmbed2
var baseStructEmbedBaseA, baseStructEmbedBaseB testStructEmbed
var base testStruct

// the patch structs
var patchStructEmbed2A, patchStructEmbed2B, patchStructEmbed2C, patchStructEmbed2D testStructEmbed2
var patchStructEmbedBaseA, patchStructEmbedBaseB testStructEmbed
var patch testStruct

// The merged structs
var mergeStructEmbed2A, mergeStructEmbed2B, mergeStructEmbed2C, mergeStructEmbed2D testStructEmbed2
var mergeStructEmbedBaseA, mergeStructEmbedBaseB testStructEmbed
var expectedMerged testStruct

func TestMerge(t *testing.T) {
	t.Run("merge identical structs", func(t *testing.T) {
		setupStructs(t)

		merged, err := mergeTestStructs(base, patch)
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, expectedMerged, *merged)
	})

	t.Run("merge identical structs as pointers", func(t *testing.T) {
		setupStructs(t)

		merged, err := mergeTestStructsPtrs(&base, &patch)
		if err != nil {
			t.Error(err)
		}
		assert.Equal(t, expectedMerged, *merged)
	})

	t.Run("different base vals are overwritten by patch - simple", func(t *testing.T) {
		s1 := simple{42, simple2{30}}
		s2 := simple{13, simple2{30}}
		expected := simple{13, simple2{30}}
		merged, err := mergeSimple(s1, s2)
		if err != nil {
			t.Error(err)
		}
		assert.NotEqual(t, s1, *merged)
		assert.Equal(t, expected, *merged)

	})

	t.Run("different base vals are overwritten by patch", func(t *testing.T) {
		setupStructs(t)

		base.F = 1342.12
		base.Struct1.Pi = NewInt(937)
		base.Struct1p.Ui = 734
		base.Struct1.Struct2.Sli = []int{123123, 1243123}

		merged, err := mergeTestStructs(base, patch)
		if err != nil {
			t.Error(err)
		}

		assert.NotEqual(t, base, *merged)
		assert.Equal(t, patch, *merged)
	})

	t.Run("nil values in patch are ignored", func(t *testing.T) {
		setupStructs(t)

		patch.Pi = nil
		patch.Struct1.Pi16 = nil

		merged, err := mergeTestStructs(base, patch)
		if err != nil {
			t.Error(err)
		}

		assert.NotEqual(t, patch, *merged)
		assert.Equal(t, expectedMerged, *merged)
	})

	t.Run("nil structs in patch are ignored", func(t *testing.T) {
		setupStructs(t)

		patch.Struct1p = nil
		patch.Struct1.Struct2p = nil

		merged, err := mergeTestStructs(base, patch)
		if err != nil {
			t.Error(err)
		}

		assert.NotEqual(t, patch, *merged)
		assert.Equal(t, expectedMerged, *merged)
	})

	t.Run("nil slices in patch are ignored", func(t *testing.T) {
		setupStructs(t)

		patch.Sls = nil
		patch.Struct1.Sli = nil
		patch.Struct1.Struct2p.Slf = nil

		merged, err := mergeTestStructs(base, patch)
		if err != nil {
			t.Error(err)
		}

		assert.NotEqual(t, patch, *merged)
		assert.Equal(t, expectedMerged, *merged)
	})

	t.Run("nil maps in patch are ignored", func(t *testing.T) {
		setupStructs(t)

		patch.Msi = nil
		patch.Mspi = nil
		patch.Struct1.Mis = nil
		patch.Struct1.Struct2p.Mspi = nil

		merged, err := mergeTestStructs(base, patch)
		if err != nil {
			t.Error(err)
		}

		assert.NotEqual(t, patch, *merged)
		assert.Equal(t, expectedMerged, *merged)
	})
}

func setupStructs(t *testing.T) {
	baseStructEmbed2A = testStructEmbed2{1, 2, 3, 4, 5, 1.1, 2.2, "test", 10, 11, 12, 12, 13,
		NewInt(14), NewInt8(15), NewInt16(16), NewInt32(17), NewInt64(18),
		NewFloat64(19.9), NewFloat32(20.1), NewString("test pointer"),
		NewUint(21), NewUint8(22), NewUint16(23), NewUint32(24), NewUint64(25),
		[]string{"test", "slice", "strings"}, []int{1, 2, 3, 4}, []float64{1.1, 2.2, 3.3},
		map[string]int{"this": 1, "is": 2, "a": 3, "map": 4}, map[int]string{1: "this", 2: "is", 3: "another"},
		map[string]*int{"wow": NewInt(1), "a map": NewInt(2), "of pointers!": NewInt(3)},
		map[int]*string{1: NewString("Another"), 2: NewString("map of"), 3: NewString("pointers, wow!")},
	}

	baseStructEmbed2B = testStructEmbed2{1, 2, 3, 4, 5, 1.1, 2.2, "test", 10, 11, 12, 12, 13,
		NewInt(14), NewInt8(15), NewInt16(16), NewInt32(17), NewInt64(18),
		NewFloat64(19.9), NewFloat32(20.1), NewString("test pointer"),
		NewUint(21), NewUint8(22), NewUint16(23), NewUint32(24), NewUint64(25),
		[]string{"test", "slice", "strings"}, []int{1, 2, 3, 4}, []float64{1.1, 2.2, 3.3},
		map[string]int{"this": 1, "is": 2, "a": 3, "map": 4}, map[int]string{1: "this", 2: "is", 3: "another"},
		map[string]*int{"wow": NewInt(1), "a map": NewInt(2), "of pointers!": NewInt(3)},
		map[int]*string{1: NewString("Another"), 2: NewString("map of"), 3: NewString("pointers, wow!")},
	}

	baseStructEmbed2C = testStructEmbed2{1, 2, 3, 4, 5, 1.1, 2.2, "test", 10, 11, 12, 12, 13,
		NewInt(14), NewInt8(15), NewInt16(16), NewInt32(17), NewInt64(18),
		NewFloat64(19.9), NewFloat32(20.1), NewString("test pointer"),
		NewUint(21), NewUint8(22), NewUint16(23), NewUint32(24), NewUint64(25),
		[]string{"test", "slice", "strings"}, []int{1, 2, 3, 4}, []float64{1.1, 2.2, 3.3},
		map[string]int{"this": 1, "is": 2, "a": 3, "map": 4}, map[int]string{1: "this", 2: "is", 3: "another"},
		map[string]*int{"wow": NewInt(1), "a map": NewInt(2), "of pointers!": NewInt(3)},
		map[int]*string{1: NewString("Another"), 2: NewString("map of"), 3: NewString("pointers, wow!")},
	}

	baseStructEmbed2D = testStructEmbed2{1, 2, 3, 4, 5, 1.1, 2.2, "test", 10, 11, 12, 12, 13,
		NewInt(14), NewInt8(15), NewInt16(16), NewInt32(17), NewInt64(18),
		NewFloat64(19.9), NewFloat32(20.1), NewString("test pointer"),
		NewUint(21), NewUint8(22), NewUint16(23), NewUint32(24), NewUint64(25),
		[]string{"test", "slice", "strings"}, []int{1, 2, 3, 4}, []float64{1.1, 2.2, 3.3},
		map[string]int{"this": 1, "is": 2, "a": 3, "map": 4}, map[int]string{1: "this", 2: "is", 3: "another"},
		map[string]*int{"wow": NewInt(1), "a map": NewInt(2), "of pointers!": NewInt(3)},
		map[int]*string{1: NewString("Another"), 2: NewString("map of"), 3: NewString("pointers, wow!")},
	}

	baseStructEmbedBaseA = testStructEmbed{1, 2, 3, 4, 5, 1.1, 2.2, "test", 10, 11, 12, 12, 13,
		NewInt(14), NewInt8(15), NewInt16(16), NewInt32(17), NewInt64(18),
		NewFloat64(19.9), NewFloat32(20.1), NewString("test pointer"),
		NewUint(21), NewUint8(22), NewUint16(23), NewUint32(24), NewUint64(25),
		[]string{"test", "slice", "strings"}, []int{1, 2, 3, 4}, []float64{1.1, 2.2, 3.3},
		map[string]int{"this": 1, "is": 2, "a": 3, "map": 4}, map[int]string{1: "this", 2: "is", 3: "another"},
		map[string]*int{"wow": NewInt(1), "a map": NewInt(2), "of pointers!": NewInt(3)},
		map[int]*string{1: NewString("Another"), 2: NewString("map of"), 3: NewString("pointers, wow!")},
		baseStructEmbed2A, &baseStructEmbed2B,
	}

	baseStructEmbedBaseB = testStructEmbed{1, 2, 3, 4, 5, 1.1, 2.2, "test", 10, 11, 12, 12, 13,
		NewInt(14), NewInt8(15), NewInt16(16), NewInt32(17), NewInt64(18),
		NewFloat64(19.9), NewFloat32(20.1), NewString("test pointer"),
		NewUint(21), NewUint8(22), NewUint16(23), NewUint32(24), NewUint64(25),
		[]string{"test", "slice", "strings"}, []int{1, 2, 3, 4}, []float64{1.1, 2.2, 3.3},
		map[string]int{"this": 1, "is": 2, "a": 3, "map": 4}, map[int]string{1: "this", 2: "is", 3: "another"},
		map[string]*int{"wow": NewInt(1), "a map": NewInt(2), "of pointers!": NewInt(3)},
		map[int]*string{1: NewString("Another"), 2: NewString("map of"), 3: NewString("pointers, wow!")},
		baseStructEmbed2C, &baseStructEmbed2D,
	}

	base = testStruct{1, 2, 3, 4, 5, 1.1, 2.2, "test", 10, 11, 12, 12, 13,
		NewInt(14), NewInt8(15), NewInt16(16), NewInt32(17), NewInt64(18),
		NewFloat64(19.9), NewFloat32(20.1), NewString("test pointer"),
		NewUint(21), NewUint8(22), NewUint16(23), NewUint32(24), NewUint64(25),
		[]string{"test", "slice", "strings"}, []int{1, 2, 3, 4}, []float64{1.1, 2.2, 3.3},
		map[string]int{"this": 1, "is": 2, "a": 3, "map": 4}, map[int]string{1: "this", 2: "is", 3: "another"},
		map[string]*int{"wow": NewInt(1), "a map": NewInt(2), "of pointers!": NewInt(3)},
		map[int]*string{1: NewString("Another"), 2: NewString("map of"), 3: NewString("pointers, wow!")},
		baseStructEmbedBaseA, &baseStructEmbedBaseB,
	}

	patchStructEmbed2A = testStructEmbed2{1, 2, 3, 4, 5, 1.1, 2.2, "test", 10, 11, 12, 12, 13,
		NewInt(14), NewInt8(15), NewInt16(16), NewInt32(17), NewInt64(18),
		NewFloat64(19.9), NewFloat32(20.1), NewString("test pointer"),
		NewUint(21), NewUint8(22), NewUint16(23), NewUint32(24), NewUint64(25),
		[]string{"test", "slice", "strings"}, []int{1, 2, 3, 4}, []float64{1.1, 2.2, 3.3},
		map[string]int{"this": 1, "is": 2, "a": 3, "map": 4}, map[int]string{1: "this", 2: "is", 3: "another"},
		map[string]*int{"wow": NewInt(1), "a map": NewInt(2), "of pointers!": NewInt(3)},
		map[int]*string{1: NewString("Another"), 2: NewString("map of"), 3: NewString("pointers, wow!")},
	}

	patchStructEmbed2B = testStructEmbed2{1, 2, 3, 4, 5, 1.1, 2.2, "test", 10, 11, 12, 12, 13,
		NewInt(14), NewInt8(15), NewInt16(16), NewInt32(17), NewInt64(18),
		NewFloat64(19.9), NewFloat32(20.1), NewString("test pointer"),
		NewUint(21), NewUint8(22), NewUint16(23), NewUint32(24), NewUint64(25),
		[]string{"test", "slice", "strings"}, []int{1, 2, 3, 4}, []float64{1.1, 2.2, 3.3},
		map[string]int{"this": 1, "is": 2, "a": 3, "map": 4}, map[int]string{1: "this", 2: "is", 3: "another"},
		map[string]*int{"wow": NewInt(1), "a map": NewInt(2), "of pointers!": NewInt(3)},
		map[int]*string{1: NewString("Another"), 2: NewString("map of"), 3: NewString("pointers, wow!")},
	}

	patchStructEmbed2C = testStructEmbed2{1, 2, 3, 4, 5, 1.1, 2.2, "test", 10, 11, 12, 12, 13,
		NewInt(14), NewInt8(15), NewInt16(16), NewInt32(17), NewInt64(18),
		NewFloat64(19.9), NewFloat32(20.1), NewString("test pointer"),
		NewUint(21), NewUint8(22), NewUint16(23), NewUint32(24), NewUint64(25),
		[]string{"test", "slice", "strings"}, []int{1, 2, 3, 4}, []float64{1.1, 2.2, 3.3},
		map[string]int{"this": 1, "is": 2, "a": 3, "map": 4}, map[int]string{1: "this", 2: "is", 3: "another"},
		map[string]*int{"wow": NewInt(1), "a map": NewInt(2), "of pointers!": NewInt(3)},
		map[int]*string{1: NewString("Another"), 2: NewString("map of"), 3: NewString("pointers, wow!")},
	}

	patchStructEmbed2D = testStructEmbed2{1, 2, 3, 4, 5, 1.1, 2.2, "test", 10, 11, 12, 12, 13,
		NewInt(14), NewInt8(15), NewInt16(16), NewInt32(17), NewInt64(18),
		NewFloat64(19.9), NewFloat32(20.1), NewString("test pointer"),
		NewUint(21), NewUint8(22), NewUint16(23), NewUint32(24), NewUint64(25),
		[]string{"test", "slice", "strings"}, []int{1, 2, 3, 4}, []float64{1.1, 2.2, 3.3},
		map[string]int{"this": 1, "is": 2, "a": 3, "map": 4}, map[int]string{1: "this", 2: "is", 3: "another"},
		map[string]*int{"wow": NewInt(1), "a map": NewInt(2), "of pointers!": NewInt(3)},
		map[int]*string{1: NewString("Another"), 2: NewString("map of"), 3: NewString("pointers, wow!")},
	}

	patchStructEmbedBaseA = testStructEmbed{1, 2, 3, 4, 5, 1.1, 2.2, "test", 10, 11, 12, 12, 13,
		NewInt(14), NewInt8(15), NewInt16(16), NewInt32(17), NewInt64(18),
		NewFloat64(19.9), NewFloat32(20.1), NewString("test pointer"),
		NewUint(21), NewUint8(22), NewUint16(23), NewUint32(24), NewUint64(25),
		[]string{"test", "slice", "strings"}, []int{1, 2, 3, 4}, []float64{1.1, 2.2, 3.3},
		map[string]int{"this": 1, "is": 2, "a": 3, "map": 4}, map[int]string{1: "this", 2: "is", 3: "another"},
		map[string]*int{"wow": NewInt(1), "a map": NewInt(2), "of pointers!": NewInt(3)},
		map[int]*string{1: NewString("Another"), 2: NewString("map of"), 3: NewString("pointers, wow!")},
		patchStructEmbed2A, &patchStructEmbed2B,
	}

	patchStructEmbedBaseB = testStructEmbed{1, 2, 3, 4, 5, 1.1, 2.2, "test", 10, 11, 12, 12, 13,
		NewInt(14), NewInt8(15), NewInt16(16), NewInt32(17), NewInt64(18),
		NewFloat64(19.9), NewFloat32(20.1), NewString("test pointer"),
		NewUint(21), NewUint8(22), NewUint16(23), NewUint32(24), NewUint64(25),
		[]string{"test", "slice", "strings"}, []int{1, 2, 3, 4}, []float64{1.1, 2.2, 3.3},
		map[string]int{"this": 1, "is": 2, "a": 3, "map": 4}, map[int]string{1: "this", 2: "is", 3: "another"},
		map[string]*int{"wow": NewInt(1), "a map": NewInt(2), "of pointers!": NewInt(3)},
		map[int]*string{1: NewString("Another"), 2: NewString("map of"), 3: NewString("pointers, wow!")},
		patchStructEmbed2C, &patchStructEmbed2D,
	}

	patch = testStruct{1, 2, 3, 4, 5, 1.1, 2.2, "test", 10, 11, 12, 12, 13,
		NewInt(14), NewInt8(15), NewInt16(16), NewInt32(17), NewInt64(18),
		NewFloat64(19.9), NewFloat32(20.1), NewString("test pointer"),
		NewUint(21), NewUint8(22), NewUint16(23), NewUint32(24), NewUint64(25),
		[]string{"test", "slice", "strings"}, []int{1, 2, 3, 4}, []float64{1.1, 2.2, 3.3},
		map[string]int{"this": 1, "is": 2, "a": 3, "map": 4}, map[int]string{1: "this", 2: "is", 3: "another"},
		map[string]*int{"wow": NewInt(1), "a map": NewInt(2), "of pointers!": NewInt(3)},
		map[int]*string{1: NewString("Another"), 2: NewString("map of"), 3: NewString("pointers, wow!")},
		patchStructEmbedBaseA, &patchStructEmbedBaseB,
	}

	mergeStructEmbed2A = testStructEmbed2{1, 2, 3, 4, 5, 1.1, 2.2, "test", 10, 11, 12, 12, 13,
		NewInt(14), NewInt8(15), NewInt16(16), NewInt32(17), NewInt64(18),
		NewFloat64(19.9), NewFloat32(20.1), NewString("test pointer"),
		NewUint(21), NewUint8(22), NewUint16(23), NewUint32(24), NewUint64(25),
		[]string{"test", "slice", "strings"}, []int{1, 2, 3, 4}, []float64{1.1, 2.2, 3.3},
		map[string]int{"this": 1, "is": 2, "a": 3, "map": 4}, map[int]string{1: "this", 2: "is", 3: "another"},
		map[string]*int{"wow": NewInt(1), "a map": NewInt(2), "of pointers!": NewInt(3)},
		map[int]*string{1: NewString("Another"), 2: NewString("map of"), 3: NewString("pointers, wow!")},
	}

	mergeStructEmbed2B = testStructEmbed2{1, 2, 3, 4, 5, 1.1, 2.2, "test", 10, 11, 12, 12, 13,
		NewInt(14), NewInt8(15), NewInt16(16), NewInt32(17), NewInt64(18),
		NewFloat64(19.9), NewFloat32(20.1), NewString("test pointer"),
		NewUint(21), NewUint8(22), NewUint16(23), NewUint32(24), NewUint64(25),
		[]string{"test", "slice", "strings"}, []int{1, 2, 3, 4}, []float64{1.1, 2.2, 3.3},
		map[string]int{"this": 1, "is": 2, "a": 3, "map": 4}, map[int]string{1: "this", 2: "is", 3: "another"},
		map[string]*int{"wow": NewInt(1), "a map": NewInt(2), "of pointers!": NewInt(3)},
		map[int]*string{1: NewString("Another"), 2: NewString("map of"), 3: NewString("pointers, wow!")},
	}

	mergeStructEmbed2C = testStructEmbed2{1, 2, 3, 4, 5, 1.1, 2.2, "test", 10, 11, 12, 12, 13,
		NewInt(14), NewInt8(15), NewInt16(16), NewInt32(17), NewInt64(18),
		NewFloat64(19.9), NewFloat32(20.1), NewString("test pointer"),
		NewUint(21), NewUint8(22), NewUint16(23), NewUint32(24), NewUint64(25),
		[]string{"test", "slice", "strings"}, []int{1, 2, 3, 4}, []float64{1.1, 2.2, 3.3},
		map[string]int{"this": 1, "is": 2, "a": 3, "map": 4}, map[int]string{1: "this", 2: "is", 3: "another"},
		map[string]*int{"wow": NewInt(1), "a map": NewInt(2), "of pointers!": NewInt(3)},
		map[int]*string{1: NewString("Another"), 2: NewString("map of"), 3: NewString("pointers, wow!")},
	}

	mergeStructEmbed2D = testStructEmbed2{1, 2, 3, 4, 5, 1.1, 2.2, "test", 10, 11, 12, 12, 13,
		NewInt(14), NewInt8(15), NewInt16(16), NewInt32(17), NewInt64(18),
		NewFloat64(19.9), NewFloat32(20.1), NewString("test pointer"),
		NewUint(21), NewUint8(22), NewUint16(23), NewUint32(24), NewUint64(25),
		[]string{"test", "slice", "strings"}, []int{1, 2, 3, 4}, []float64{1.1, 2.2, 3.3},
		map[string]int{"this": 1, "is": 2, "a": 3, "map": 4}, map[int]string{1: "this", 2: "is", 3: "another"},
		map[string]*int{"wow": NewInt(1), "a map": NewInt(2), "of pointers!": NewInt(3)},
		map[int]*string{1: NewString("Another"), 2: NewString("map of"), 3: NewString("pointers, wow!")},
	}

	mergeStructEmbedBaseA = testStructEmbed{1, 2, 3, 4, 5, 1.1, 2.2, "test", 10, 11, 12, 12, 13,
		NewInt(14), NewInt8(15), NewInt16(16), NewInt32(17), NewInt64(18),
		NewFloat64(19.9), NewFloat32(20.1), NewString("test pointer"),
		NewUint(21), NewUint8(22), NewUint16(23), NewUint32(24), NewUint64(25),
		[]string{"test", "slice", "strings"}, []int{1, 2, 3, 4}, []float64{1.1, 2.2, 3.3},
		map[string]int{"this": 1, "is": 2, "a": 3, "map": 4}, map[int]string{1: "this", 2: "is", 3: "another"},
		map[string]*int{"wow": NewInt(1), "a map": NewInt(2), "of pointers!": NewInt(3)},
		map[int]*string{1: NewString("Another"), 2: NewString("map of"), 3: NewString("pointers, wow!")},
		mergeStructEmbed2A, &mergeStructEmbed2B,
	}

	mergeStructEmbedBaseB = testStructEmbed{1, 2, 3, 4, 5, 1.1, 2.2, "test", 10, 11, 12, 12, 13,
		NewInt(14), NewInt8(15), NewInt16(16), NewInt32(17), NewInt64(18),
		NewFloat64(19.9), NewFloat32(20.1), NewString("test pointer"),
		NewUint(21), NewUint8(22), NewUint16(23), NewUint32(24), NewUint64(25),
		[]string{"test", "slice", "strings"}, []int{1, 2, 3, 4}, []float64{1.1, 2.2, 3.3},
		map[string]int{"this": 1, "is": 2, "a": 3, "map": 4}, map[int]string{1: "this", 2: "is", 3: "another"},
		map[string]*int{"wow": NewInt(1), "a map": NewInt(2), "of pointers!": NewInt(3)},
		map[int]*string{1: NewString("Another"), 2: NewString("map of"), 3: NewString("pointers, wow!")},
		mergeStructEmbed2C, &mergeStructEmbed2D,
	}

	expectedMerged = testStruct{1, 2, 3, 4, 5, 1.1, 2.2, "test", 10, 11, 12, 12, 13,
		NewInt(14), NewInt8(15), NewInt16(16), NewInt32(17), NewInt64(18),
		NewFloat64(19.9), NewFloat32(20.1), NewString("test pointer"),
		NewUint(21), NewUint8(22), NewUint16(23), NewUint32(24), NewUint64(25),
		[]string{"test", "slice", "strings"}, []int{1, 2, 3, 4}, []float64{1.1, 2.2, 3.3},
		map[string]int{"this": 1, "is": 2, "a": 3, "map": 4}, map[int]string{1: "this", 2: "is", 3: "another"},
		map[string]*int{"wow": NewInt(1), "a map": NewInt(2), "of pointers!": NewInt(3)},
		map[int]*string{1: NewString("Another"), 2: NewString("map of"), 3: NewString("pointers, wow!")},
		mergeStructEmbedBaseA, &mergeStructEmbedBaseB,
	}

}

func NewInt8(n int8) *int8          { return &n }
func NewInt16(n int16) *int16       { return &n }
func NewInt32(n int32) *int32       { return &n }
func NewFloat64(f float64) *float64 { return &f }
func NewFloat32(f float32) *float32 { return &f }
func NewUint(n uint) *uint          { return &n }
func NewUint8(n uint8) *uint8       { return &n }
func NewUint16(n uint16) *uint16    { return &n }
func NewUint32(n uint32) *uint32    { return &n }
func NewUint64(n uint64) *uint64    { return &n }
