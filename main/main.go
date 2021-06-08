package main

import (
	"log"
	"system-design/strategy"
)

func main() {
	log.Println("main started")
	invoice := &strategy.Invoice{
		GrossAmount: 0,
		Vendor: strategy.Vendor{
			VendorName: "BlackHawk",
		},
		ContractorName: "Black",
	}
	strategyService := strategy.NewStrategyService(invoice)
	strategyService.StartStrategyService()
}