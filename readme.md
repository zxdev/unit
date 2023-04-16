# unit

reads and builds a ```map[string]string``` from a simple unit style resource file


./sample
```
[sample]
key1 = value1
key2 = value2

[sample2]
# comments are supported as are colon seperators and 
# multi-segemented values are not disturbed on parse
key1 : value1
key2 : value2,value3,value4
key3 : value3:value4:value5
key4 : value4=value5
```

```golang

	var u unit.Unit
	u.Parse("./sample", "sample")    

	// u = map{key1:value1, key2:value2}
	fmt.Println(u)
```