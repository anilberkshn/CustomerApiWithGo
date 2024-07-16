package services

import (
	"CustomerApi/Customer/src/entities"
	upRequestModel "CustomerApi/Customer/src/entities/requestModels"
	"CustomerApi/Customer/src/entities/responseModels"
	"CustomerApi/Customer/src/repositories"
	sharedentities "CustomerApi/shared/entities"
	"fmt"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

type Service struct {
	repo *repositories.Repo
}

func NewService(repo *repositories.Repo) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetAll(limit, offset int64) ([]*responseModels.CustomerResponseModel, error) {
	customers, err := s.repo.GetAll(limit, offset)
	if err != nil {
		fmt.Println(err, "-GetAll Service")
		return nil, err
	}
	return customers, nil
}

func (s *Service) Create(customer entities.Customer) (string string, e error) {

	// todo : alttaraf handlerda
	customerAddress := sharedentities.Address{
		AddressLine: customer.Address.AddressLine,
		City:        customer.Address.City,
		CityCode:    customer.Address.CityCode,
		Country:     customer.Address.Country,
	}

	// todo : repositorye taşınacak.
	var customerId = uuid.New().String()
	var timeNow = time.Now().String() // todo : stringe çekilmemeli

	var customerProp = bson.M{
		"_id":        customerId,
		"first_name": customer.FirstName,
		"last_name":  customer.LastName,
		"email":      customer.Email,
		"phone":      customer.Phone,
		"address":    customerAddress, // Todo: mantık yürüttüm hata alıyordu düzeldi.
		"created_at": timeNow,
		"updated_at": timeNow,
	}

	insertedId, err := s.repo.Create(customerProp) // todo _ yapınca hata görünmedi
	if err != nil {
		fmt.Println(err, "- Create Service")
		return "error", e
	}
	return *insertedId, nil
}

// s.repo.Create(customer) // Too many arguments in call to 's.repo.Create' hatası gitti

func (s *Service) GetByID(id string) (*responseModels.CustomerResponseModel, error) {
	customer, err := s.repo.GetById(id)
	if err != nil {
		fmt.Println(err, "- Service: GetByID")
		return nil, err
	}
	return customer, nil
}

func (s *Service) Delete(id string) (bool, error) {

	success, err := s.repo.Delete(id)
	if err != nil {
		fmt.Println(err, "- Service: Delete")
		return false, err
	}

	return success, nil
}

func (s *Service) Update(id string, updateData *upRequestModel.UpdateRequestModel) error {

	err := s.repo.Update(id, updateData)
	if err != nil {
		return err
	}

	return nil // todo id direk response edilebilir.
}
