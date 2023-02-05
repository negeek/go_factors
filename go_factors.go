package go_factors
import (
		"math"
)

func MathFactors(number int)([]int, error){
	factors := make([]int, number) 
	latest_index:=0
	for num:=1; num < int(math.Sqrt(float64(number))+1); num++{
		if number%num==0{
			factors[latest_index]=num
			latest_index++
		
			other_num:= int(math.Floor(float64(number)/float64(num)))
			if other_num != factors[latest_index]{
				factors[latest_index]=other_num
				latest_index++
			}else{

				return factors[:latest_index], nil
			}
		}
		
	}
	return factors[:latest_index], nil
}