# HTML Link Parser

### This code parses a given html and collect all the a tags return them as a list of Link structs.

Link struct:

```go
type Link struct {
	Href string
	Text string
}
```
