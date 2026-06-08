package migration

import (
	"log"

	"github.com/kai-zenn/bljr_go_api/api/configs"
	"github.com/kai-zenn/bljr_go_api/api/model"
  "github.com/google/uuid"
  "golang.org/x/crypto/bcrypt"
)


func SeedingDB() {
	// 1. Cek dulu apakah tabel Accesses masih kosong
	var countAccess int64
	configs.DB.Model(&model.Access{}).Count(&countAccess)

	if countAccess == 0 {
		createAccess := model.Access{AccessId: 1, AccessName: "create"}
		readAccess   := model.Access{AccessId: 2, AccessName: "read"}
		updateAccess := model.Access{AccessId: 3, AccessName: "update"}
		deleteAccess := model.Access{AccessId: 4, AccessName: "delete"}

		configs.DB.Create(&[]model.Access{createAccess, readAccess, updateAccess, deleteAccess})

		roles := []model.Role{
			{
				RoleId:   1,
				RoleName: "admin",
				Accesses: []model.Access{createAccess, readAccess, updateAccess, deleteAccess},
			},
			{
				RoleId:   2,
				RoleName: "member",
				Accesses: []model.Access{readAccess},
			},
			{
				RoleId:   3,
				RoleName: "penulis",
				Accesses: []model.Access{createAccess, readAccess, updateAccess},
			},
		}

		if err := configs.DB.Create(&roles).Error; err != nil {
			log.Println("Gagal seeding Roles & Accesses:", err)
		} else {
			log.Println("Seeding Roles, Accesses, dan Role_Access BERHASIL!")
		}
	}
}

func SeedUsers() {
	var countUser int64
	configs.DB.Model(&model.User{}).Count(&countUser)

	if countUser == 0 {
		var adminRole, memberRole, penulisRole model.Role
		configs.DB.First(&adminRole, "role_name = ?", "admin")
		configs.DB.First(&memberRole, "role_name = ?", "member")
		configs.DB.First(&penulisRole, "role_name = ?", "penulis")

		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("rahasia124"), bcrypt.DefaultCost)

		users := []model.User{
			{
				ID:        uuid.New(), 
				FirstName: "Kai",
				LastName:  "Zen",
				Username:  "kaizen",
				Email:     "admin@example.com",
				Password:  string(hashedPassword),
				// 🌟 Triknya di sini: Masukkan objek role ke dalam slice Roles
				// User ini akan otomatis jadi Admin dan Member sekaligus di tabel user_role
				Roles:     []model.Role{adminRole, memberRole}, 
			},
			{
				ID:        uuid.New(),
				FirstName: "Kusnadi",
				LastName:  "Writed",
				Username:  "koesnadi",
				Email:     "kusnadi@smkn46.com",
				Password:  string(hashedPassword),

				Roles:     []model.Role{penulisRole},
			},
		}
		
		if err := configs.DB.Create(&users).Error; err != nil {
			log.Println("Gagal seeding Users:", err)
		} else {
			log.Println("Seeding Users dan User_Role berhasil!")
		}
	}
}
