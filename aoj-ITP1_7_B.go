
import "fmt"

func main() {
	for {
		var n, x, cnt int
		fmt.Scan(&n, &x)
		if n == 0 && x == 0 {
			break
		}
		for i := 1; i <= n; i ++ {
			for j := i+1; j <= n; j++ {
				for k := j+1; k <= n; k++ {
					if i + j + k == x {
						// fmt.Println(i, j, k)
						cnt++
					}
				}
			}
		}
		fmt.Println(cnt)
	}
}
