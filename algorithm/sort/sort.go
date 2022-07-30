package sort

//冒泡排序
func Bsort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := 1; i < len(arr)-i; i++ {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
}
