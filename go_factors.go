package go_factors
import (
		"math"
		"errors"
)

func MathFactors(number int)([]int, error){
	factors:=[]int
	latest_index=0
	for num:=1; i< math.Sqrt(number)+1; i++{
		if number%num==0{
			append(factors,num)
			latest_index+=1
		}
		if math.Floor(number,num) != factors[latest_index]{
			append(factors,math.Floor(number,num))
		}else{
			
			return factors, nil
		}

		return factors, nil
	}
}