def bubble(arr):
    print(f"Bubble sort started")
    n = len(arr)
    swapped = False
    for i in range(n-1):
        for j in range(0, n-i-1):
            if arr[j] > arr[j + 1]:
                swapped = True
                arr[j], arr[j + 1] = arr[j + 1], arr[j]
        if not swapped:
            return

def quick(arr, low, high):
    print(f"Quick sort with {low}, {high}")
    if low < high:
        pi = partition(arr, low, high)
        quick(arr, low, pi - 1)
        quick(arr, pi + 1, high)

def partition(arr, low, high):
    pivot = arr[high]
    i = low - 1
    for j in range(low, high):
        if arr[j] <= pivot:
            i = i + 1
            (arr[i], arr[j]) = (arr[j], arr[i])

    (arr[i + 1], arr[high]) = (arr[high], arr[i + 1])

    return i + 1

import time
def sleep(arr):
    for i in range(1, len(arr) + 1):
        print(f"Sleep {i} time(s)")
        time.sleep(1)
