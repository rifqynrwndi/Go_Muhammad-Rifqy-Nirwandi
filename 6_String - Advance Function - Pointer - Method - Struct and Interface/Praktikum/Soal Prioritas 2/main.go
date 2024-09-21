package main

import (
	"rental/customer"
	"rental/vehicle"
)

func main() {
	rs := customer.RentalSystem{}

	rs.AddCustomer("Rifqy", "Gunungpati")
	rs.AddCustomer("Mulyono", "Solo")

	rs.RentVehicle("Mulyono", vehicle.Vehicle{Brand: "Toyota", Type: "Avanza", RentalCost: 500000})
	rs.RentVehicle("Rifqy", vehicle.Vehicle{Brand: "Honda", Type: "Civic", RentalCost: 300000})

	rs.ViewCustomerInfo("Rifqy")
	rs.ViewCustomerInfo("Mulyono")
	rs.ViewRentedVehicles("Rifqy")
	rs.ViewRentedVehicles("Mulyono")

	rs.ReturnVehicle("Rifqy", "Civic")
	rs.ViewRentedVehicles("Rifqy")
}
