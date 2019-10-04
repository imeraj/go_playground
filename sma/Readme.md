# Data structure for moving average

```
type sma struct {
	size uint
	data []float32
	sum  float64
}  
```

Here -   
    **size**: stores window size N  
    **data**: is a go **slice** type that stores last N data points   
    **sum**: stores sum of last N data points stored in **data**  

# Interface for moving average 

```
type smaOps interface {
	addData(val float32)
	getAvg() float64
	toString() string
}
```

This is a simple interface that defines operations on the data structure. 

Here -   
    **addData**   - updates sum, stores new val in **data** and adjusts **data** so that only last N data points are stored in **data**  
    **getAvg**    - computes average of last N data points stored in **data**  
    **toString**  - provides access to the elements in **data** slice by returning string representation of last N data points stored in **data**  

# Additional

Basic error handling has been added in case of input processing and in case there is not enough data in the **data** slice.

# Output from sample run 

Enter window size: 3
Enter numbers for calculating average - 
1
Not enough data points!

3
Not enough data points!

5

 Average: 3.000000 (1.000000 3.000000 5.000000)

6

 Average: 4.666667 (3.000000 5.000000 6.000000)

8

 Average: 6.333333 (5.000000 6.000000 8.000000)