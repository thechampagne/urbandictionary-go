# Urban Dictionary

[![](https://img.shields.io/github/v/tag/thechampagne/urbandictionary-go?label=version)](https://github.com/thechampagne/urbandictionary-go/releases/latest) [![](https://img.shields.io/github/license/thechampagne/urbandictionary-go)](https://github.com/thechampagne/urbandictionary-go/blob/main/LICENSE)

Urban Dictionary API client for **Go**.

### Download

```
go get github.com/thechampagne/urbandictionary-go
```

### Example

```go
func main() {
	urban := urbandictionary.New("Golang",1)
	data, err := urban.Data()
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range data {
		fmt.Println(v.Definition)
	}
}
```

### License

This repo is released under the [MIT License](https://github.com/thechampagne/urbandictionary-go/blob/main/LICENSE).