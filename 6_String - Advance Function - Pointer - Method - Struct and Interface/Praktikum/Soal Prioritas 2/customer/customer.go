package customer

import (
	"fmt"
	"rental/vehicle"
)

type Customer struct {
	Nama     string
	Alamat  string
	Rentals  []vehicle.Vehicle
}

type RentalSystem struct {
	Customers []Customer
}

func (rs *RentalSystem) AddCustomer(nama, alamat string) {
	rs.Customers = append(rs.Customers, Customer{Nama: nama, Alamat: alamat})
}

func (rs *RentalSystem) RentVehicle(customerName string, vehicle vehicle.Vehicle) {
	for i, customer := range rs.Customers {
		if customer.Nama == customerName {
			rs.Customers[i].Rentals = append(rs.Customers[i].Rentals, vehicle)
			break
		}
	}
}

func (rs *RentalSystem) ReturnVehicle(customerName, vehicleType string) {
	for i, customer := range rs.Customers {
		if customer.Nama == customerName {
			for j, v := range rs.Customers[i].Rentals {
				if v.Type == vehicleType {
					rs.Customers[i].Rentals = append(rs.Customers[i].Rentals[:j], rs.Customers[i].Rentals[j+1:]...)
					fmt.Println(customerName, "Mengembalikan Mobil", vehicleType)
					break
				}
			}
			break
		}
	}
}

func (rs *RentalSystem) ViewCustomerInfo(customerName string) {
	for _, customer := range rs.Customers {
		if customer.Nama == customerName {
			fmt.Println("Nama:", customer.Nama, ", Alamat:", customer.Alamat)
			break
		}
	}
}

func (rs *RentalSystem) ViewRentedVehicles(customerName string) {
	for _, customer := range rs.Customers {
		if customer.Nama == customerName {
			fmt.Println("Mobil Disewa untuk", customer.Nama)
			for _, v := range customer.Rentals {
				fmt.Println("-", v.Brand, v.Type, ", Harga Rental:", v.RentalCost)
			}
			break
		}
	}
}
