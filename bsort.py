array = [13, 2, 43, 25, 4, 4, 52, 63, 12]
print("Starting the bubble sort function")
print(f"Initial state of the array: {array}")

arrayLen = len(array)

swapped = False

for i in range(arrayLen):
    swapped = False
    for j in range(arrayLen - i - 1):
        if array[j] > array[j+1]:
            aux = array[j+1]
            array[j+1] = array[j]
            array[j] = aux
            swapped = True

    if not swapped:
        break

print("End of sorting")
print(f"Sorted array: {array}")
