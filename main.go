package main
import(
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/fatih/color"
	"fmt"
	"math"
	"strings"
	)
func makeBar(r int)string{
	ri := int(math.Round(float64(r) * 0.2))
	usedBar := strings.Repeat("■",ri)
	uf := 20 -ri
	FreeBar := strings.Repeat("■",uf)
	Bar := color.RedString(usedBar) + color.BlueString(FreeBar)
	return Bar
}
func main(){
	fmt.Println("Memory Bar (red=used,blue=free)")
	for{
		v,_ := mem.VirtualMemory()
		used := v.Used / 1024 / 1024 
		total := v.Total / 1024 / 1024
		r := int(float64(used) / float64(total)*100)
		BarString := makeBar(r)
		fmt.Printf("\r %s%%",BarString)
	}
}
