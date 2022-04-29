package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// dataType()
	// goVariable()
	// defaultValue()

	// // ví dụ ve rune
	// name := "Hello World"
	// printBytes(name)
	// fmt.Printf("\n")
	// printChars(name)
	// fmt.Printf("\n\n")
	// name = "Señor"
	// printBytes(name)
	// fmt.Printf("\n")
	// printChars(name)

	// // tạo string từ slice chứa các rune
	// runeSlice := []rune{0x0053, 0x0065, 0x00f1, 0x006f, 0x0072}
	// str := string(runeSlice)
	// fmt.Println(str)
	// // vòng lặp for range của 1 string
	// text := "Señor"
	// for index, rune := range text {
	// 	fmt.Printf("%c starts at byte %d\n", rune, index)
	// }
	// // tạo string từ slice chức các byte
	// byteSlice := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9}
	// str1 := string(byteSlice)
	// fmt.Println(str1)

	// byteSlice1 := []byte{67, 97, 102, 195, 169}
	// str2 := string(byteSlice1)
	// fmt.Println(str2)

	// // ví dụ mô tả 2 text
	// word1 := "ｿｹｯﾄを作成する"
	// length(word1)
	// fmt.Println("len work1: ", len(word1))
	// word2 := "abcdefghi"
	// length(word2)
	// fmt.Println("len work2: ", len(word2))

	// word1Rune := []rune(word1)
	// fmt.Print("sub-string 1 : ")
	// printChars(string(word1Rune[1:]))

	// fmt.Printf("\nsub-string 2 : ")
	// printChars(string(word2[1:]))
	// subString(1, 3, "ﾄを作成する")

	var slice1 = []int{7, 8, 4, 3, 3, 4, 9}
	fmt.Println("mang sau khi sort", sortArr(slice1))
	fmt.Println("rank :", rankSort(slice1))
	fmt.Printf("\n\n\n")
	var slice2 = []int{7, 8, 4, 3, 3, 4, 9}
	fmt.Println(rankSort2(slice2))
}

type Infor struct {
	Name, Address string
	Age           int
}

func dataType() {
	/*
		- Các kiểu dữ liệu trong golang :
			+ Basic type: Numbers, strings, and booleans
			+ Aggregate type: Array and structs
			+ Reference type: Pointers, slices, maps, functions, and channels
			+ Interface type
	*/
	var (
		firstName, lastName string
		age                 int
		salary              float64
		isConfirmed         bool
	)

	fmt.Printf("firstName: %T, lastName: %T, age: %T, salary: %T, isConfirmed: %T\n",
		firstName, lastName, age, salary, isConfirmed)
	// struct
	tuanInfo := Infor{Name: "Tuan", Address: "155 khuong trung", Age: 26}

	fmt.Println(" Infor name : ", tuanInfo.Name)
	fmt.Println(" Infor adress : ", tuanInfo.Address)
	fmt.Println(" Infor age : ", tuanInfo.Age)
	// array

	var arr = [3]int{}
	fmt.Printf("arr %T\n", arr)
	var sl = []int{}
	fmt.Printf("slice %T\n", sl)

	//slice
	a := [5]int{76, 77, 78, 79, 80}
	var b []int = a[1:4] //creates a slice from a[1] to a[3]
	fmt.Println(b)
	// tạo slice sử dụng make : make([]T, len, cap) []T có thể được sử dụng để tạo slice bằng cách truyền vào type (loại), độ dài và sức chứa
	i := make([]int, 5, 5)
	fmt.Println(i)
	//slice không giới hạn dộ dài. muốn thêm kích thước sử dụng hàm append
	cars := []string{"Ferrari", "Honda", "Ford"}
	fmt.Println("cars:", cars, "has old length", len(cars), "and capacity", cap(cars))
	cars = append(cars, "Toyota")
	fmt.Println("cars:", cars, "has new length", len(cars), "and capacity", cap(cars))

	// Map : map[key]value{}
	sammy := map[string]string{"name": "Sammy", "animal": "shark", "color": "blue", "location": "ocean"}
	fmt.Println(sammy)
	fmt.Println(sammy["color"])

}

func goVariable() {
	/*
		- khai báo biến trong golang có 2 cách:
			+ cách 1 : cú pháp
				var variableName dataType = initialValue , initialValue có thể khai hoặc không
				var name string = "Tuan"
				var age int = 26

				var name string
				name = "Tuan"
				var age int
				age = 26

				var name = "Tuan"
				var age = 20
	*/
	var name = "Tuan"
	var age = 26

	fmt.Println("name: ", name, "age: ", age)

	/*
		+ cách 2 : cú pháp
				variableName := initialValue

				name := "Tuan"
				age := 26

			- cú pháp này sử dụng để khởi tạo và gán biến nhận kết quả trả về từ func khác
						result := calulateAge(abc)


			- Khai báo nhiều biến cùng 1 lúc:
				+ var var1, var2, var3 int
				+ var var1, var2, var3 int = 1, 2, 3
				+ var var1, var2, var3 = 1, 2, false
				+ var1, var2, var3 := 1, 2.2, false

			- Có thể sử dụng khai báo biến trong block var()
				var myName = "Tuan"						var(
				var myAge = 26                     => 		myName = "Tuan"
				var mySalary = 5.000.000					myAge = 26
															mySalary = 5.000.000
														)
	*/
	var1, var2, var3 := 1, 2.2, false
	fmt.Println("value var1 :", var1)
	fmt.Println("value var1 :", var2)
	fmt.Println("value var1 :", var3)
}

func defaultValue() {
	/*

		- Giá trị Default Value của các kiểu dữ liệu :
			+ Bất kì biến nào được khai báo nhưng không có có giá trị biến (initial value) thì sẽ có một giá trị 0 (zero value). Giá trị 0 này sẽ riêng cho từng loại kiểu dữ liệu.
			+ bool : false
			+ string : “”
			+ int, int8, int16 : 0
			+ float32, float64 : 0.0
			+ pointer : nil
	*/
	var (
		firstName, lastName string
		age                 int
		salary              float64
		isConfirmed         bool
	)

	fmt.Printf("firstName: %s, lastName: %s, age: %d, salary: %f, isConfirmed: %t\n",
		firstName, lastName, age, salary, isConfirmed)
}

// Start Rune
func typeRune() {
	// ra đời giúp hiển thị các ký tự đặc biệt không nằm trong ACSII có chiều dài lớn 1 byte
	// - một cách gọi khác của kiểu dữ liệu int32. Rune đại diện cho một điểm mã trong Go. Không quan trọng điểm mã chứa bao nhiêu byte, nó có thể đại diện bởi một rune.

	// string là 1 kiểu dữ liệu bất biến , khi muốn làm việc với string => chúng ta chuyển string thành slice chứa các rune, chúng ta thay đổi slice và sau đó chuyển lại thành string mới.
	h := "hello"
	fmt.Println(mutate([]rune(h)))
}

func mutate(s []rune) string {
	s[0] = 'a'
	return string(s)
}

func printBytes(s string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
}

func printChars(s string) {
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c ", runes[i])
	}
}

// End Rune

// so sanh 2 string "ｿｹｯﾄを作成する" , "abcdefghi"

func length(s string) {
	fmt.Printf("length of %s is %d\n", s, utf8.RuneCountInString(s))
}

func subString(startIndex int, endIndex int, s string) string {
	runes := []rune(s)
	fmt.Println(len(runes))
	if startIndex < 0 {
		startIndex = 0
	}
	if endIndex > len(runes) {
		endIndex = len(runes)
	}
	subS := string(runes[startIndex:endIndex])
	fmt.Printf("%d indexStart %d indexEnd data after sub : %s ", startIndex, endIndex, subS)
	return subS
}

// có 2 cách làm
// cách 1: sort -> tạo rank

func sortArr(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] >= arr[j] {
				temp := arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
		}
	}
	return arr
}

//sort rank of arrays

func rankSort(arr []int) []int {
	sliceRank := make([]int, len(arr))
	arrSort := sortArr(arr)
	// tạo 1 mảng chứa rank default tăng dần tới len(arr)
	for i := 0; i < len(arr); i++ {
		sliceRank[i] = i + 1
	}
	//đánh dấu mang rank dựa vào mảng giá trị
	for j := 0; j < len(arrSort)-1; j++ {
		if arrSort[j] == arrSort[j+1] {
			sliceRank[j+1] = j + 1
		}
	}
	// trả về mảng rank
	return sliceRank

}

// cách 2 tìm rank trên mảng: đang có chút lỗi => fix sau
func rankSort2(list []int) []int {
	rank_list := make([]int, len(list))
	sort_list := make([]int, len(list))

	for i := 0; i < len(list); i++ {
		for j := 0; j < i; j++ {
			if list[i] > list[j] {
				rank_list[i]++
			} else if list[i] == list[j] {
				continue
			} else {
				rank_list[j]++
			}
		}
	}
	fmt.Printf("hang cua moi phan tu mang\n")
	fmt.Printf("\n\t Phan tu\t\tHang\n")
	for i := 0; i < len(list); i++ {
		fmt.Printf("\t    %d\t\t\t %d\n", list[i], rank_list[i]+1)
	}

	for j := 0; j < len(list); j++ {
		sort_list[rank_list[j]] = list[j]
	}

	fmt.Printf("\nSap xep mang tang dan theo hang\n\t")
	for i := 0; i < len(list); i++ {
		fmt.Printf("%d ", sort_list[i])
	}

	return sort_list

}
