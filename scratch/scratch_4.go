package main

import (
	"fmt"
	"github.com/pkg/errors"
	"reflect"
)

func main() {
	var x float64 = 3.2
	v := reflect.ValueOf(x)
	t := reflect.TypeOf(x)

	fmt.Println(t)
	fmt.Println(v)
	fmt.Println(v.String())
	fmt.Println(v.Kind() == reflect.Float64)
	fmt.Println("Type:", v.Type())
	fmt.Println("Kind:", v.Kind())
	fmt.Println("Type == TypeOf", v.Type() == reflect.TypeOf(x))
	fmt.Println(v.Float())

	type myInt int
	var y myInt = 42
	yv := reflect.ValueOf(y)
	fmt.Println(yv.Kind(), yv.Type())
	fmt.Printf("after Interface: %T\n", yv.Interface())
	fmt.Printf("just the interface: %v\n", yv.Interface())
	fmt.Printf("just the value: %v\n", yv)
	fmt.Println(yv.Interface())
	fmt.Println(reflect.TypeOf(v))
	fmt.Println(reflect.TypeOf(v.Interface()))
	fmt.Printf("value of x is %v\n", v)
	fmt.Printf("value of x is %7.1e\n", v.Interface())

	//v.SetFloat(7.1) // panic
	fmt.Printf("can set? %t\n\n", v.CanSet())
	// think of it as f(x). Can't set x within f, right?

	// Now, try f(&x) -- won't work.
	// Why? We don't want to set &x (the pointer)
	w := reflect.ValueOf(&x)
	fmt.Printf("value of w: %v\n", w)
	fmt.Printf("type of w: %v\n", w.Type())
	fmt.Printf("kind of w: %v\n", w.Kind())
	fmt.Printf("can set w? %t\n\n", w.CanSet())

	// We want to set *x (the value in the addressable location x)
	we := w.Elem()
	fmt.Printf("value of we: %v\n", we)
	fmt.Printf("type of we: %v\n", we.Type())
	fmt.Printf("kind of we: %v\n", we.Kind())
	fmt.Printf("can set we? %t\n", we.CanSet())

	fmt.Printf("\nSetting to 7.1.\n")
	we.SetFloat(7.1)
	fmt.Printf("value of we: %v\n", we)
	fmt.Printf("interface of we: %v\n\n", we.Interface())

	type myStruct struct {
		Name   string
		Age    int
		Height float64
	}
	m := myStruct{"Chris", 29, 6 + 3/12.0}

	ve := reflect.ValueOf(&m).Elem()
	vt := ve.Type()
	fmt.Printf("struct type: %s\n", vt)
	for i := 0; i < ve.NumField(); i++ {
		f := ve.Field(i)
		fmt.Printf("name: %s, type: %s, interface: %v\n",
			vt.Field(i).Name, f.Type(), f.Interface())
	}
	ve.Field(0).SetString("Christopher")
	ve.FieldByName("Age").SetInt(31)
	ve.Field(2).SetFloat(6 + 4.0/12)
	fmt.Printf("%s is now: %v\n\n", vt, m)

	mp := myStructOfPtr{NewString("Chris2"), 25, NewFloat(5 + 3/12.0),
		nestedStruct{NewInt(180), []int{175, 170, 160}, NewFloat(47.5),
			&wifeStruct{NewInt(105), map[string]int{"before": 100, "way Before": 105, "much before": 110},
				NewFloat(10.3)}}, nil}
	vtp := reflect.TypeOf(mp)
	fmt.Printf("struct type: %s\n", vtp)
	for i := 0; i < vtp.NumField(); i++ {
		f := vtp.Field(i)
		fmt.Printf("name: %s, type: %s, elem: %v\n",
			f.Name, f.Type, reflect.ValueOf(mp).Field(i))
	}

	fmt.Println("\nTrying to print via ValueOf...")
	names := []string{"Name", "Age", "Height", "MyWeight"}
	vep2 := reflect.ValueOf(&mp).Elem()
	for _, i := range names {
		sf, _ := reflect.TypeOf(mp).FieldByName(i)
		vByName := vep2.FieldByName(i)
		if vByName.Kind() == reflect.Ptr {
			vByName = vByName.Elem()
		}
		vByIdx := vep2.FieldByIndex(sf.Index)
		if vByIdx.Kind() == reflect.Ptr {
			vByIdx = vByIdx.Elem()
		}
		fmt.Printf("name: %s, type: %s, index: %v, elem: %v (by name: %v)\n",
			sf.Name, sf.Type, sf.Index, vByIdx, vByName)
		if vByIdx.Kind() == reflect.Struct {
			// works:
			// sf2, _ := reflect.TypeOf(vByIdx.Elem().Interface()).FieldByName("Weight")
			sf2, _ := reflect.TypeOf(vByName.Interface()).FieldByName("Weight")
			vByName2 := vByName.FieldByName("Weight").Elem()
			vByIdx2 := vep2.FieldByIndex([]int{3, 0}).Elem()
			fmt.Printf("name: %s, type: %s, index: %v, elem: %v (by name: %v)\n",
				sf2.Name, sf2.Type, sf2.Index, vByIdx2, vByName2)
		}
	}

	pathMap := map[string]interface{}{
		"settings": map[string]interface{}{
			"devKey":   true,
			"cloudKey": "safekey",
			"moreSettings": map[string]interface{}{
				"devPwd":   "you'llneverguessit",
				"cloudPwd": "totallySafePwd",
			},
		},
		"plugins": map[string]interface{}{
			"myPlugin1_num_of_users": 123,
		},
		"freeValue":    false,
		"anotherValue": 123.91,
	}
	fmt.Println("\n***map:\n", pathMap)

	fmt.Printf("***in path form: %q\n", getPaths(pathMap))

	vep := reflect.ValueOf(&mp).Elem()
	vep.Field(0).Elem().SetString("Christoph")
	vep.FieldByName("Age").SetInt(32)
	vep.Field(2).Elem().SetFloat(6 + 5.0/12)
	vep.Field(3).Field(0).Elem().SetInt(185)
	vep.Field(3).FieldByName("Weight").Elem().SetInt(190)
	fmt.Printf("%s is now: %v", vtp, mp)

	fmt.Printf("\n\n")
	s := &myStructOfPtr{NewString("Bob"), 28, NewFloat(4 + 3/12.0),
		nestedStruct{NewInt(180), []int{175, 170, 160, 58}, NewFloat(47.5),
			&wifeStruct{NewInt(105), map[string]int{"before": 999, "way Before": 105, "much before": 110, "now!": 100},
				NewFloat(10.3)}}, []string{"nice", "dad"}}
	fmt.Printf("Name: %s, Age: %v, Height: %v, Weight: %v, Stats: %q\n",
		getVal(s, []string{"Name"}),
		getVal(s, []string{"Age"}),
		getVal(s, []string{"Height"}),
		getVal(s, []string{"MyWeight", "Weight"}),
		getVal(s, []string{"Stats"}))

	setVal(s, []string{"Name"}, "Bob2")
	setVal(s, []string{"Age"}, 99)
	setVal(s, []string{"Height"}, 6.2)
	setVal(s, []string{"MyWeight", "Weight"}, 160)
	setVal(s, []string{"Stats"}, []string{"really", "nice", "dad"})

	fmt.Printf("Name: %s, Age: %v, Height: %v, Weight: %v, Stats: %q\n\n",
		getVal(s, []string{"Name"}),
		getVal(s, []string{"Age"}),
		getVal(s, []string{"Height"}),
		getVal(s, []string{"MyWeight", "Weight"}),
		getVal(s, []string{"Stats"}))

	s2 := &myStructOfPtr{nil, 32, NewFloat(4 + 3/12.0),
		nestedStruct{NewInt(180), []int{77, 170, 999}, NewFloat(47.5),
			&wifeStruct{nil, map[string]int{"before": 100, "way Before": 105, "much before": 66},
				NewFloat(10.4)}}, []string{"nice", "dad"}}

	fmt.Println(s)
	fmt.Println(s2)

	s3, err := merge(s, s2)
	fmt.Printf("\nok? %t\nBase: %v\nPatch: %v\nMerge: %v\n", err == nil, s, s2, s3)
}

func mergeTestStructs(base, patch testStruct) (*testStruct, error) {
	ret, err := merge(base, patch)
	if err != nil {
		return nil, err
	}
	retTS := ret.(testStruct)
	return &retTS, nil
}

func mergeTestStructsPtrs(base, patch *testStruct) (*testStruct, error) {
	ret, err := merge(base, patch)
	if err != nil {
		return nil, err
	}
	retTS := ret.(testStruct)
	return &retTS, nil
}

// Will not add new fields to base.
// Will replace entire maps and slices (at the moment).
func merge(base interface{}, patch interface{}) (interface{}, error) {
	if reflect.TypeOf(base) != reflect.TypeOf(patch) {
		return nil, fmt.Errorf("cannot merge different types. base type: %s, patch type: %s",
			reflect.TypeOf(base), reflect.TypeOf(patch))
	}

	baseType := reflect.TypeOf(base)
	if baseType.Kind() == reflect.Ptr {
		baseType = baseType.Elem()
	}

	ret := reflect.New(baseType)
	if err := initStruct(baseType, ret.Elem()); err != nil {
		return nil, err
	}

	if err := mergeRec(base, patch, ret.Elem()); err != nil {
		return nil, err
	}
	return ret.Elem().Interface(), nil
}

func mergeRec(base interface{}, patch interface{}, ret reflect.Value) (err error) {
	//defer func() {
	//	if x := recover(); x != nil {
	//		err = fmt.Errorf("panic in mergerec. btype: %v, ptype: %v",
	//			reflect.TypeOf(base), reflect.TypeOf(patch))
	//	}
	//}()

	bt := reflect.TypeOf(base)
	bv := reflect.ValueOf(base)
	pv := reflect.ValueOf(patch)
	if bt.Kind() == reflect.Ptr { // or slice, map
		bt = bt.Elem()
		bv = bv.Elem()
		// if the patch value is nil, just assign the base value and we're done.
		if pv.IsNil() {
			ret.Set(bv)
			return nil
		}
		pv = pv.Elem()
	}
	for i := 0; i < bt.NumField(); i++ {
		patchWasNil := false
		bti := bt.Field(i).Type
		bvi := bv.Field(i)
		pvi := pv.Field(i)
		rvi := ret.Field(i)
		switch bti.Kind() {
		case reflect.Ptr:
			patchWasNil = pvi.IsNil()
			bti = bti.Elem()
			bvi = bvi.Elem()
			pvi = pvi.Elem()
			rvi = rvi.Elem()
		case reflect.Slice:
			patchWasNil = pvi.IsNil()
		case reflect.Map:
			patchWasNil = pvi.IsNil()
		}

		if bti.Kind() == reflect.Struct {
			err := mergeRec(bv.Field(i).Interface(), pv.Field(i).Interface(), rvi)
			if err != nil {
				return fmt.Errorf("failed in mergeRec on field # %d of type %v, err: %v", i, bti, err)
			}
		} else {
			if !patchWasNil && bvi != pvi {
				rvi.Set(pvi)
			} else {
				rvi.Set(bvi)
			}
		}
	}
	return nil
}

func initStruct(t reflect.Type, v reflect.Value) (err error) {
	defer func() {
		if x := recover(); x != nil {
			err = fmt.Errorf("panic in initStruct. type: %v, kind: %v, is value settable? %t. "+
				"if not settable, check if are you trying to merge a struct with an unexported field",
				t, v.Kind(), v.CanSet())
		}
	}()

	// WIP: if we want to support structs with unexported fields, we need to do that explicitly.
	for i := 0; i < v.NumField(); i++ {
		ft := t.Field(i).Type
		fv := v.Field(i)

		switch ft.Kind() {
		case reflect.Slice:
			fv.Set(reflect.MakeSlice(ft, 0, 0))
		case reflect.Map:
			fv.Set(reflect.MakeMap(ft))
		case reflect.Chan:
			fv.Set(reflect.MakeChan(ft, 0))
		case reflect.Struct:
			if e := initStruct(ft, fv); e != nil {
				return e
			}
		case reflect.Ptr:
			p := reflect.New(ft.Elem())
			if ft.Elem().Kind() == reflect.Struct {
				if e := initStruct(ft.Elem(), p.Elem()); e != nil {
					return e
				}
			}
			fv.Set(p)
		default:
		}
	}
	return nil
}

func MergeOrig(base interface{}, patch interface{}) (interface{}, error) {
	if reflect.TypeOf(base) != reflect.TypeOf(patch) {
		return nil, fmt.Errorf("cannot merge different types. base type: %s, patch type: %s",
			reflect.TypeOf(base), reflect.TypeOf(patch))
	}

	baseType := reflect.TypeOf(base)
	if baseType.Kind() == reflect.Ptr {
		baseType = baseType.Elem()
	}

	ret := reflect.New(baseType)
	if err := initStruct(baseType, ret.Elem()); err != nil {
		return nil, err
	}

	if err := mergeRecOrig(base, patch, ret.Elem()); err != nil {
		return nil, err
	}
	return ret.Elem().Interface(), nil
}

// mergeRec recursively merges into ret
func mergeRecOrig(base interface{}, patch interface{}, merged reflect.Value) (err error) {
	bt := reflect.TypeOf(base)
	bv := reflect.ValueOf(base)
	pv := reflect.ValueOf(patch)
	if bt.Kind() == reflect.Ptr {
		// TODO: find out how to assign nil (not just a zero value)
		if bv.IsNil() && pv.IsNil() {
			// if base and patch are nil, leave it
			return nil
		} else if pv.IsNil() {
			// if the patch value is nil, assign base value and done
			merged.Set(bv.Elem())
			return nil
		} else if bv.IsNil() {
			// if the base value is nil, assign patch value and done
			merged.Set(pv.Elem())
			return nil
		}
		bt = bt.Elem()
		bv = bv.Elem()
		pv = pv.Elem()
	}
	for i := 0; i < bt.NumField(); i++ {
		patchWasNil := false
		bti := bt.Field(i).Type
		bvi := bv.Field(i)
		pvi := pv.Field(i)
		mvi := merged.Field(i)
		switch bti.Kind() {
		case reflect.Ptr:
			patchWasNil = pvi.IsNil()
			bti = bti.Elem()
			bvi = bvi.Elem()
			pvi = pvi.Elem()
			mvi = mvi.Elem()
		case reflect.Slice:
			patchWasNil = pvi.IsNil()
		case reflect.Map:
			patchWasNil = pvi.IsNil()
		}

		if !mvi.CanSet() {
			continue
		}

		if bti.Kind() == reflect.Struct {
			err := mergeRecOrig(bv.Field(i).Interface(), pv.Field(i).Interface(), mvi)
			if err != nil {
				return fmt.Errorf("failed in mergeRec on field # %d of type %v, err: %v", i, bti, err)
			}
		} else {
			if !patchWasNil && bvi != pvi {
				mvi.Set(pvi)
			} else {
				mvi.Set(bvi)
			}
		}
	}
	return nil
}

// mergeRec recursively merges into ret
func mergeRec2(base interface{}, patch interface{}, merged interface{}) (err error) {
	bt := reflect.TypeOf(base)
	mt := reflect.TypeOf(merged)
	bv := reflect.ValueOf(base)
	mv := reflect.ValueOf(merged)
	pv := reflect.ValueOf(patch)

	if bt != mt {
		fmt.Printf("****bt %v != mt %v\n", bt, mt)
	}
	// handle pointers
	if bt.Kind() == reflect.Ptr {
		bvNil := bv.IsNil()
		pvNil := pv.IsNil()
		bt = bt.Elem()
		mt = mt.Elem()
		bv = bv.Elem()
		mv = mv.Elem()
		pv = pv.Elem()
		// if the base or patch value is nil, overwrite and we're done.
		if bvNil && pvNil {
			mv.Set(reflect.Zero(reflect.ValueOf(merged).Type().Elem()))
			return nil
		} else if bvNil && mv.CanSet() {
			mv.Set(pv)
			return nil
		} else if pvNil && mv.CanSet() {
			mv.Set(bv)
			return nil
		}
	}

	// ignore unexported fields
	//if !mv.CanSet() {
	//	return
	//}

	patchWasNil := false
	fmt.Printf("**** bvkind %v, mvkind %v\n", bv.Kind(), mv.Kind())
	switch bt.Kind() {
	case reflect.Struct:
		fmt.Printf("*** bt numfield %v, mt numbfield %v\n", bt.NumField(), mt.NumField())
		for i := 0; i < bt.NumField(); i++ {
			if !mv.Field(i).CanSet() {
				continue
			}
			err := mergeRec2(bv.Field(i).Interface(), pv.Field(i).Interface(), mv.Field(i).Interface())
			if err != nil {
				return errors.Wrap(err, fmt.Sprintf("failed in mergeRec on field # %d of type %v, err: %v", i, bt.Field(i).Type, err))
			}
		}
	default:
		if mv.Kind() == reflect.Ptr {
			mv = mv.Elem()
		}
		if !patchWasNil && bv != pv {
			mv.Set(pv)
		} else {
			mv.Set(bv)
		}
	}

	return nil
}

// initStruct creates a new struct of type t
func initStructOrig(tnotused reflect.Type, v reflect.Value) error {
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		ft := t.Field(i).Type
		fv := v.Field(i)

		// ignore unexported fields
		if !fv.CanSet() {
			continue
		}

		switch ft.Kind() {
		case reflect.Slice:
			fv.Set(reflect.MakeSlice(ft, 0, 0))
		case reflect.Map:
			fv.Set(reflect.MakeMap(ft))
		case reflect.Chan:
			fv.Set(reflect.MakeChan(ft, 0))
		case reflect.Struct:
			if err := initStruct(ft, fv); err != nil {
				return err
			}
		case reflect.Ptr:
			p := reflect.New(ft.Elem())
			if ft.Elem().Kind() == reflect.Struct {
				if err := initStruct(ft.Elem(), p.Elem()); err != nil {
					return err
				}
			}
			fv.Set(p)
		default:
		}
	}
	return nil
}

func getPaths(m map[string]interface{}) [][]string {
	return getPathsRec(m, nil)
}

//func getPathsRec(src interface{}) [][]string {
//	var allPaths [][]string
//	if reflect.ValueOf(src).Kind() == reflect.Map {
//		for k, v := range src.(map[string]interface{}) {
//			for _, path := range getPathsRec(v) {
//				allPaths = append(allPaths, append([]string{k}, path...))
//			}
//		}
//	} else {
//		allPaths = append(allPaths, []string{})
//	}
//	return allPaths
//}

func getPathsRec(src interface{}, curPath []string) [][]string {
	if reflect.ValueOf(src).Kind() == reflect.Map {
		paths := [][]string{}
		for k, v := range src.(map[string]interface{}) {
			paths = append(paths, getPathsRec(v, append(curPath, k))...)
		}
		return paths
	}

	return [][]string{curPath}
}

//func getPaths(m map[string]interface{}) [][]string {
//	return getPathsRec(m, nil)
//}

//func getPathsRec(src interface{}, curPath []string, allPaths [][]string) [][]string {
//	if reflect.ValueOf(src).Kind() == reflect.Map {
//		for k, v := range src.(map[string]interface{}) {
//			allPaths = getPathsRec(v, append(curPath, k), allPaths)
//		}
//	} else {
//		allPaths = append(allPaths, curPath)
//	}
//	return allPaths
//}

func getVal(src interface{}, path []string) reflect.Value {
	var val reflect.Value
	if reflect.ValueOf(src).Kind() == reflect.Ptr {
		val = reflect.ValueOf(src).Elem().FieldByName(path[0])
	} else {
		val = reflect.ValueOf(src).FieldByName(path[0])
	}
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	if val.Kind() == reflect.Struct {
		return getVal(val.Interface(), path[1:])
	}
	return val
}

func setVal(src interface{}, path []string, newVal interface{}) {
	val := getVal(src, path)
	switch val.Kind() {
	case reflect.Bool:
		val.SetBool(newVal.(bool))
	case reflect.String:
		val.SetString(newVal.(string))
	case reflect.Int:
		val.SetInt(int64(newVal.(int)))
	case reflect.Int8:
		val.SetInt(int64(newVal.(int8)))
	case reflect.Int16:
		val.SetInt(int64(newVal.(int16)))
	case reflect.Int32:
		val.SetInt(int64(newVal.(int32)))
	case reflect.Int64:
		val.SetInt(newVal.(int64))
	case reflect.Float32:
		val.SetFloat(float64(newVal.(float32)))
	case reflect.Float64:
		val.SetFloat(newVal.(float64))
	case reflect.Uint8:
		val.SetUint(uint64(newVal.(uint8)))
	case reflect.Uint16:
		val.SetUint(uint64(newVal.(uint16)))
	case reflect.Uint32:
		val.SetUint(uint64(newVal.(uint32)))
	case reflect.Uint64:
		val.SetUint(newVal.(uint64))
	case reflect.Slice:
		val.Set(reflect.ValueOf(newVal))
		// TODO: maps?
	}
}

type wifeStruct struct {
	Weight            *int
	OverTime          map[string]int
	ChangeFromHighest *float64
}

func (ns wifeStruct) String() string {
	var w int
	var cfh float64
	if ns.Weight == nil {
		w = 0
	} else {
		w = *ns.Weight
	}
	if ns.ChangeFromHighest == nil {
		cfh = 0
	} else {
		cfh = *ns.ChangeFromHighest
	}

	return fmt.Sprintf("{weight: %d, OverTime: %v, ChangeFromHighest: %3.2f}",
		w, ns.OverTime, cfh)
}

type nestedStruct struct {
	Weight            *int
	OverTime          []int
	ChangeFromHighest *float64
	WifeWeight        *wifeStruct
}

func (ns nestedStruct) String() string {
	var w int
	var cfh float64
	if ns.Weight == nil {
		w = 0
	} else {
		w = *ns.Weight
	}
	if ns.ChangeFromHighest == nil {
		cfh = 0
	} else {
		cfh = *ns.ChangeFromHighest
	}
	return fmt.Sprintf("{weight: %d, OverTime: %v, ChangeFromHighest: %3.2f, WifeWeight: %v}",
		w, ns.OverTime, cfh, ns.WifeWeight)
}

type myStructOfPtr struct {
	Name     *string
	Age      int
	Height   *float64
	MyWeight nestedStruct
	Stats    []string
}

func (m myStructOfPtr) String() string {
	var n string
	var h float64
	var a int
	if m.Name == nil {
		n = "nil"
	} else {
		n = *m.Name
	}
	if m.Height == nil {
		h = 0
	} else {
		h = *m.Height
	}
	return fmt.Sprintf("{%s, %d, %d, %f, %v}", n, m.Age, a, h, m.MyWeight)
}

func NewBool(b bool) *bool        { return &b }
func NewInt(n int) *int           { return &n }
func NewInt64(n int64) *int64     { return &n }
func NewString(s string) *string  { return &s }
func NewFloat(f float64) *float64 { return &f }

func mergeSimple(base, patch simple) (*simple, error) {
	ret, err := merge(base, patch)
	if err != nil {
		return nil, err
	}
	retS := ret.(simple)
	return &retS, nil
}

type simple struct {
	I  int
	S2 simple2
}

type simple2 struct {
	I int
}

type testStruct struct {
	I        int
	I8       int8
	I16      int16
	I32      int32
	I64      int64
	F        float64
	F32      float32
	S        string
	Ui       uint
	Ui8      uint8
	Ui16     uint32
	Ui32     uint32
	Ui64     uint64
	Pi       *int
	Pi8      *int8
	Pi16     *int16
	Pi32     *int32
	Pi64     *int64
	Pf       *float64
	Pf32     *float32
	Ps       *string
	Pui      *uint
	Pui8     *uint8
	Pui16    *uint16
	Pui32    *uint32
	Pui64    *uint64
	Sls      []string
	Sli      []int
	Slf      []float64
	Msi      map[string]int
	Mis      map[int]string
	Mspi     map[string]*int
	Mips     map[int]*string
	Struct1  testStructEmbed
	Struct1p *testStructEmbed
}

type testStructEmbed struct {
	I        int
	I8       int8
	I16      int16
	I32      int32
	I64      int64
	F        float64
	F32      float32
	S        string
	Ui       uint
	Ui8      uint8
	Ui16     uint32
	Ui32     uint32
	Ui64     uint64
	Pi       *int
	Pi8      *int8
	Pi16     *int16
	Pi32     *int32
	Pi64     *int64
	Pf       *float64
	Pf32     *float32
	Ps       *string
	Pui      *uint
	Pui8     *uint8
	Pui16    *uint16
	Pui32    *uint32
	Pui64    *uint64
	Sls      []string
	Sli      []int
	Slf      []float64
	Msi      map[string]int
	Mis      map[int]string
	Mspi     map[string]*int
	Mips     map[int]*string
	Struct2  testStructEmbed2
	Struct2p *testStructEmbed2
}

type testStructEmbed2 struct {
	I     int
	I8    int8
	I16   int16
	I32   int32
	I64   int64
	F     float64
	F32   float32
	S     string
	Ui    uint
	Ui8   uint8
	Ui16  uint32
	Ui32  uint32
	Ui64  uint64
	Pi    *int
	Pi8   *int8
	Pi16  *int16
	Pi32  *int32
	Pi64  *int64
	Pf    *float64
	Pf32  *float32
	Ps    *string
	Pui   *uint
	Pui8  *uint8
	Pui16 *uint16
	Pui32 *uint32
	Pui64 *uint64
	Sls   []string
	Sli   []int
	Slf   []float64
	Msi   map[string]int
	Mis   map[int]string
	Mspi  map[string]*int
	Mips  map[int]*string
}
