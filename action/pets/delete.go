package pets

import (
	"github.com/jinzhu/gorm"
	"tyun.cn/terraform-ali-petstore/model/pet"
)

//DeletePetRequest request struct
type DeletePetRequest struct {
	ID string
}

//DeletePet deletes a pet from database
func DeletePet(db *gorm.DB, req *DeletePetRequest) (error) {
	err := pet.Delete(db, req.ID)
	return err
}