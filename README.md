# go-stuff

```
go build myprog.go
```
```
gccgo -O3 -march=core-avx2 -ffast-math myprog.go
```
### Benchmark
```
go test -bench=.
```
```
import (
  "fmt"
	"testing"
)
func BenchmarkPrint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("huhu")
	}
}
```
