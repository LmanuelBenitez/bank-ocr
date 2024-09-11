En las historia de usuario existe una con los numeros: 000000051 (La cual la marca con el estado de 'OK')

Pero al momento de aplicar la formula (1*d1 + 2*d2 + 3*d3 + ...) mod 11 = 0

Usando:
numberPosition := map[string]int{
		"d9": 3,
		"d8": 4,
		"d7": 5,
		"d6": 8,
		"d5": 8,
		"d4": 2,
		"d3": 8,
		"d2": 6,
		"d1": 5,
	}
 
El resultado que arroja es 23 ya que (5*4 + 1*3) = 23  --> 23%11 no es igual a 0

------------------------------------------------------------------------------------------
Asi como en la historia: 123456789 (La marca como 'ERR')

El resultado que arroja es (1*5 + 2*6 + 3*8 + 4*2 + 5*8 + 6*8 + 7*5 + 8*4 + 9*3) --> (5 + 12 + 24 + 8 + 40 + 48 + 35 + 32 + 27) = 231 --> 231%11 = 0
