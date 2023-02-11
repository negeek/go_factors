package go_factors
import (
		"math"
)

func MathFactors(number int)[]int{
	factors := make([]int, number) 
	latest_index:=0
	for num:=1; num < int(math.Sqrt(float64(number))+1); num++{
		if number%num==0{
			factors[latest_index]=num
			latest_index++
		
			other_num:= int(math.Floor(float64(number)/float64(num)))
			if other_num != factors[latest_index-1]{
				factors[latest_index]=other_num
				latest_index++
			}else{

				return factors[:latest_index]
			}
		}
		
	}
	return factors[:latest_index]
}

func SquareFactors(number int)[]int{
	factors:= make([]int,number)
	latest_index:=0
	for num:=1; num < int(math.Sqrt(float64(number))+1); num++{
		if number%num==0{
			num_root:=int(math.Sqrt(float64(num)))
			// if perfect square
			if num_root*num_root==num{
				factors[latest_index]=num
				latest_index++
				//check if other num is perfect square
				other_num:= int(math.Floor(float64(number)/float64(num)))
				other_num_root:=int(math.Sqrt(float64(number)/float64(num)))
				if other_num != factors[latest_index-1] && other_num_root*other_num_root==other_num{
					factors[latest_index]=other_num
					latest_index++
				}
			}else{
				other_num:= int(math.Floor(float64(number)/float64(num)))
				other_num_root:=int(math.Sqrt(float64(number)/float64(num)))
				if other_num != factors[latest_index-1] && other_num_root*other_num_root==other_num{
					factors[latest_index]=other_num
					latest_index++
				}

			}
		}
	}
	return factors[:latest_index]
}

func EvenFactors(number int)[]int{
	if number%2!=0{
		factors:= make([]int,0)
		return factors
	}
	factors:= make([]int,number)
	latest_index:=0
	for num:=1; num < int(math.Sqrt(float64(number))+1); num++{
		if number%num==0{
			//if num is even
			if num%2==0{
				factors[latest_index]=num
				latest_index++
				other_num:= int(math.Floor(float64(number)/float64(num)))
				if other_num != factors[latest_index-1] && other_num%2==0{
					factors[latest_index]=other_num
					latest_index++
				}

			}else{
				other_num:= int(math.Floor(float64(number)/float64(num)))
				if other_num != factors[latest_index-1] && other_num%2==0{
					factors[latest_index]=other_num
					latest_index++
				}
			}
			
			
		}
	}
	return factors[:latest_index]
}

func OddFactors(number int)[]int{
	factors:= make([]int,number)
	latest_index:=0
	for num:=1; num < int(math.Sqrt(float64(number))+1); num++{
		if number%num==0{
			//if num is odd
			if num%2!=0{
				factors[latest_index]=num
				latest_index++
				other_num:= int(math.Floor(float64(number)/float64(num)))
				if other_num != factors[latest_index-1] && other_num%2!=0{
					factors[latest_index]=other_num
					latest_index++
				}
			
			}else{
				other_num:= int(math.Floor(float64(number)/float64(num)))
				if other_num != factors[latest_index-1] && other_num%2!=0{
					factors[latest_index]=other_num
					latest_index++
				}
			}
			
			
		}
	}
	return factors[:latest_index]

}

func PrimeFactors(number int)[]int{
	factors:=MathFactors(number)
	prime_factor_space=factors[2:]
	if len(factors)<3{
        return factors[1]
	}
	if  len(prime_factor_space)<2{
		return factors[0]
	}
	num:=prime_factor_space[0]
	var memory map[int]int
	next_num:=make([]int,number)
	next_num_index:=0

	if num >2{
		next_num[next_num_index]=num+2
		next_num_index++
        memory[num+2]=1

	}else{
		next_num[next_num_index]=num+1
		next_num_index++
        memory[num+1]=1
	}
	for{
		for i:=1; i < len(prime_factor_space); i++{
			if num!=prime_factor_space[i]{
				if prime_factor_space[i]%num==0{
					prime_factor_space[i]=0
				}else{
					there, ok= memory[prime_factor_space[i]]
					if ok==false{
						memory[prime_factor_space[i]]=1
						next_num[next_num_index]=prime_factor_space[i]
					}
				}
			}
		}
		if next_num_index-1<0{
			break
		}else{
			num=next_num[next_num_index-1]
			next_num[next_num_index-1]=0
			next_num_index-=1
		}
		
	}
	
}