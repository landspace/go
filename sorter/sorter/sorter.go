// sorter.go
package main

import "bufio"
import "flag"
import "fmt"
import "io"
import "os"
import "strconv"
import "time"

import "./algorithms/bubblesort"

//import "algorithms/qsort"

var infile *string = flag.String("i", "unsorted.dat", "File contains values for sorting")
var outfile *string = flag.String("o", "sorted.dat", "File to receive sorted values")
var algorithm *string = flag.String("a", "qsort", "Sort algorithm")

func readValues(infile string) (values []int, err error) {
	file, err := os.Open(infile)
	if err != nil {
		fmt.Println("Failed to open the input file ", file)
		return
	}
	defer file.Close()

	br := bufio.NewReader(file)

	values = make([]int, 0)

	for {
		line, isPerfix, err1 := br.ReadLine()
		if err1 != nil {
			if err1 != io.EOF {
				err = err1
			}
			break
		}

		if isPerfix {
			fmt.Println("A too log line, seem unexpected.")
			return
		}

		str := string(line) // 转换字符数组为字符串

		value, err1 := strconv.Atoi(str)

		if err1 != nil {
			err = err1
			return
		}

		values = append(values, value)
	}
	return
}

func writeValues(values []int, outfile string) error {
	file, err := os.Create(outfile)
	if err != nil {
		fmt.Println("Faild to create the output file ", file)
		return err
	}
	defer file.Close()

	for _, value := range values {
		str := strconv.Itoa(value)
		file.WriteString(str + "\n")
	}
	return nil
}

func main() {

	if infile != nil {
		fmt.Println("infile =", *infile, "outfile =", *outfile, "algorithm =", *algorithm)
	}

	values, err := readValues("unsorted.dat")
	if err == nil {
		t1 := time.Now()
		//		switch *algorithm {
		//		case "bubblesort":
		bubblesort.BubbleSort(values)
		//		default:
		//			fmt.Println("Sorting algorithm", *algorithm, "is either unknow or unsupported.")
		//		}

		t2 := time.Now()
		fmt.Println("The sorting process costs", t2.Sub(t1), "to complete.")

		writeValues(values, "sorted.dat")
	} else {
		fmt.Println(err)
	}
}
