package go_factors
import (
		"math"
		"errors"
)

func MathFactors(number float64)([]float64, error){
	factors:=[]float64
	latest_index=0
	for num:=1; i< math.Sqrt(number)+1; i++{
		if math.Mod(number,num)==0{
			append(factors,num)
			latest_index++
		}
		if math.Floor(number,num) != factors[latest_index]{
			append(factors,math.Floor(number,num))
		}else{
			
			return factors, nil
		}

		return factors, nil
	}
}