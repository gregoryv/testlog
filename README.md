Package testlog to catch logs while testing

## Quick start

    $ go get github.com/gregoryv/testlog
	
in your code

```go
func TestYourthing(t *testing.T) {
	buf := testlog.Catch(t) 
	
	// do something
	
	got := buf.String()	
	exp := "some log phrase"
	if !strings.Contains(got, exp) {
		t.Error(got, " missing expected ", exp)
	}
}
```
