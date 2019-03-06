package scratch

import (
	"fmt"
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
	fmt.Println(v.Float(), "\n")

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
	fmt.Printf("value of x is %7.1e\n", v)
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

	mp := myStructOfPtr{NewString("Chris2"), NewInt(25), NewFloat(5 + 3/12.0),
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
	fmt.Println("\nmap:", pathMap)

	fmt.Printf("in path form: %q", getPaths(pathMap))

	vep := reflect.ValueOf(&mp).Elem()
	vep.Field(0).Elem().SetString("Christoph")
	vep.FieldByName("Age").Elem().SetInt(32)
	vep.Field(2).Elem().SetFloat(6 + 5.0/12)
	vep.Field(3).Field(0).Elem().SetInt(185)
	vep.Field(3).FieldByName("Weight").Elem().SetInt(190)
	fmt.Printf("%s is now: %v", vtp, mp)

	fmt.Println("\n")
	s := &myStructOfPtr{NewString("Bob"), NewInt(28), NewFloat(4 + 3/12.0),
		nestedStruct{NewInt(180), []int{175, 170, 160}, NewFloat(47.5),
			&wifeStruct{NewInt(105), map[string]int{"before": 100, "way Before": 105, "much before": 110},
				NewFloat(10.3)}}, []string{"nice", "dad"}}
	fmt.Printf("Name: %s, Age: %v, Height: %v, Weight: %v, Stats: %q\n",
		getVal(s, []string{"Name"}),
		getVal(s, []string{"Age"}),
		getVal(s, []string{"Height"}),
		getVal(s, []string{"MyWeight", "Weight"}),
		getVal(s, []string{"Stats"}))

	setVal(s, []string{"Name"}, "Bob2")
	setVal(s, []string{"Age"}, 32)
	setVal(s, []string{"Height"}, 6.2)
	setVal(*s, []string{"MyWeight", "Weight"}, 160)
	setVal(s, []string{"Stats"}, []string{"really", "nice", "dad"})

	fmt.Printf("Name: %s, Age: %v, Height: %v, Weight: %v, Stats: %q\n\n",
		getVal(s, []string{"Name"}),
		getVal(s, []string{"Age"}),
		getVal(s, []string{"Height"}),
		getVal(s, []string{"MyWeight", "Weight"}),
		getVal(s, []string{"Stats"}))

	s2 := &myStructOfPtr{NewString("Bob3"), NewInt(32), NewFloat(4 + 3/12.0),
		nestedStruct{NewInt(180), []int{175, 170, 160}, NewFloat(47.5),
			&wifeStruct{NewInt(105), map[string]int{"before": 100, "way Before": 105, "much before": 110},
				NewFloat(10.3)}}, []string{"nice", "dad"}}

	s3, ok := merge(s2, s)
	fmt.Printf("\nMerged? %t, %v\n", ok, s3)
}

func merge(base interface{}, patch interface{}) (interface{}, bool) {
	if reflect.TypeOf(base) != reflect.TypeOf(patch) {
		return nil, false
	}
	baseType := reflect.TypeOf(base)
	if baseType.Kind() == reflect.Ptr {
		baseType = baseType.Elem()
	}
	ret := reflect.New(baseType)
	initStruct(baseType, ret.Elem())
	fmt.Printf("ret: %T, %v\n", ret.Elem(), ret.Elem())

	//fmt.Printf("ret.Elem(): %T, .Interface() %T\n", ret.Elem(), ret.Elem().Interface())
	//fmt.Printf("raw myStructOfPtr: %T\n", reflect.TypeOf(myStructOfPtr{}))
	//fmt.Printf("equal? %t, %T, %T", baseType == reflect.TypeOf(myStructOfPtr{}),
	//	baseType, reflect.TypeOf(myStructOfPtr{}))

	ok := mergeRec(base, patch, ret.Elem())
	if !ok {
		return nil, false
	}
	fmt.Printf("\nReturned: %T, %v", ret, ret)
	return ret, true
}

func mergeRec(base interface{}, patch interface{}, ret reflect.Value) bool {
	bt := reflect.TypeOf(base)
	pt := reflect.TypeOf(patch)
	bv := reflect.ValueOf(base)
	pv := reflect.ValueOf(patch)
	if bt.Kind() == reflect.Ptr { // or slice, map
		bt = bt.Elem()
		pt = pt.Elem()
		bv = bv.Elem()
		pv = pv.Elem()
	} else {
		// the interfaces were not passed as pointers; we have to make them settable
		bv = reflect.ValueOf(&base).Elem()
		pv = reflect.ValueOf(&patch).Elem()
	}
	// always want ret to be settable
	rv := reflect.ValueOf(&ret).Elem()
	fmt.Printf("is rv settable? %t\n", rv.CanSet())

	for i := 0; i < bt.NumField(); i++ {
		bti := bt.Field(i).Type
		pti := pt.Field(i).Type
		if bti.Kind() == reflect.Ptr { // or slice, map
			bti = bti.Elem()
			pti = pti.Elem()
		}
		rvi := rv.Field(i).Elem()
		fmt.Printf("Field %d: %v, type: %s, is struct? %t, is settable? %t, index: %v\n", i, bt.Field(i).Name, bt.Field(i).Type,
			bti.Kind() == reflect.Struct, bv.CanSet(), bt.Field(i).Index)
		fmt.Printf("Return field: type: %s, is struct? %t, is settable? %t\n",
			rvi.Type(), rvi.Kind() == reflect.Struct, rvi.CanSet())

		if bt.Field(i).Type.Kind() == reflect.Struct {
			newi, ok := mergeRec(bv.Field(i).Interface(), pv.Field(i).Interface())
			if !ok {
				return newi, false
			}
			rvi.Set(newi)
		} else {
			// for now, just use base:
			//rvi.Set(bv.Field(i))
			rv.Field(i).SetString("test") // rvi is "value obtained using unexported field"

		}
	}
	return rv, true
}

func initStruct(t reflect.Type, v reflect.Value) {
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
			initStruct(ft, fv)
		case reflect.Ptr:
			fmt.Printf("ptr is: %v\n", ft.Elem())
			p := reflect.New(ft.Elem())
			if ft.Elem().Kind() == reflect.Struct {
				initStruct(ft.Elem(), p.Elem())
			}
			fv.Set(p)
		default:
		}
	}
}

func getPaths(m map[string]interface{}) [][]string {
	return getPathsRec(m, nil, nil)
}

func getPathsRec(src interface{}, curPath []string, allPaths [][]string) [][]string {
	if reflect.ValueOf(src).Kind() == reflect.Map {
		for k, v := range src.(map[string]interface{}) {
			allPaths = getPathsRec(v, append(curPath, k), allPaths)
		}
	} else {
		allPaths = append(allPaths, curPath)
	}
	return allPaths
}

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
	return fmt.Sprintf("{weight: %d, OverTime: %v, ChangeFromHighest: %v}",
		*ns.Weight, ns.OverTime, *ns.ChangeFromHighest)
}

type nestedStruct struct {
	Weight            *int
	OverTime          []int
	ChangeFromHighest *float64
	WifeWeight        *wifeStruct
}

func (ns nestedStruct) String() string {
	return fmt.Sprintf("{weight: %d, OverTime: %v, ChangeFromHighest: %3.2f, WifeWeight: %v}",
		*ns.Weight, ns.OverTime, *ns.ChangeFromHighest, ns.WifeWeight)
}

type myStructOfPtr struct {
	Name     *string
	Age      *int
	Height   *float64
	MyWeight nestedStruct
	Stats    []string
}

func (m myStructOfPtr) String() string {
	return fmt.Sprintf("{%s, %d, %f, %v}", *m.Name, *m.Age, *m.Height, m.MyWeight)
}

func NewBool(b bool) *bool        { return &b }
func NewInt(n int) *int           { return &n }
func NewInt64(n int64) *int64     { return &n }
func NewString(s string) *string  { return &s }
func NewFloat(f float64) *float64 { return &f }
