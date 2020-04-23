package weight

func KgToLB(kg KG) LB {
	return LB(kg * kgToLBRatio)
}
