package main

import(
	"fmt"
)

type Data struct {
	key 	string
	values 	[]string
}

func InitData(key, value string) (*Data) {
	return &Data{key, []string{value}}
}

func (d Data) IsIt(key string) bool {
	if d.key == key {
		return true
	} else {
		return false
	}
}

func (d Data) Print() {
	fmt.Printf("key[%s] = ", d.key)
	for _, val := range d.values {
		fmt.Printf("%s ", val)
	}
	fmt.Printf("\n")
}

func (d *Data) Append(value string) (int, bool) {
	for pos, val := range d.values {
		if val == value {
			return pos,true
		}
	}
	d.values = append(d.values, value)

	return len(d.values), true
}

type DataMap struct {
	data_map	map[string] []Data
}

func DataMapInit() (*DataMap) {
	return &(DataMap{map[string][]Data{}})
}

func (dm DataMap) Print() {
	for key, datas := range dm.data_map {
		fmt.Printf("key[%s] = ", key)
		for _, data := range datas {
			data.Print()
		}
	}
}

func (dm DataMap) Search(section string) ([]Data) {
	data_array, result := dm.data_map[section]
	if result {
		return data_array
	} else {
		return nil
	}
	
}

func (dm *DataMap) Add(section, key, value string) bool {
	data_array := dm.Search(section)
	if data_array == nil {
		dm.data_map[section] = []Data {*InitData(key, value)}
	} else {
		for _, data := range data_array {
			if data.IsIt(key) {
				data.Append(value)
				return true
			}
		}
		// the key cannot be found, so add a new one.
		dm.data_map[section] = append(data_array, *InitData(key, value))
	}

	return true
}

func main() {
	d0 := InitData("d0", "1")
	if d0.IsIt("d0") == true {
		fmt.Printf("d0 is the key!\n")
	} else {
		fmt.Printf("d0 is not the key!\n")
	}
	if d0.IsIt("d1") == true {
		fmt.Printf("d1 is the key!\n")
	} else {
		fmt.Printf("d1 is not the key!\n")
	}
	d0.Print()

	d0.Append("10")	
	d0.Print()

	d0.Append("-20")	
	d0.Print()

	datamap0 := DataMapInit()
	datamap0.Print()

	datamap0.Add("sec0", "key0", "0")
	datamap0.Print()
	datamap0.Add("sec1", "key1", "1")
	datamap0.Print()
	
}
