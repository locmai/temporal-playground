package workflows

func bubble(arr[] int)[]int {
   for i:=0; i< len(arr)-1; i++ {
      for j:=0; j < len(arr)-i-1; j++ {
         if (arr[j] > arr[j+1]) {
            arr[j], arr[j+1] = arr[j+1], arr[j]
         }
      }
   }
   return arr
}

func quick(arr []int, low, high int) []int {
	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		arr = quick(arr, low, p-1)
		arr = quick(arr, p+1, high)
	}
	return arr
}

func partition(arr []int, low, high int) ([]int, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}
