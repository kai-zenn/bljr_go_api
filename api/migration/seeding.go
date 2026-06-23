package migration

import (
	"log"

	"github.com/kai-zenn/bljr_go_api/api/configs"
	"github.com/kai-zenn/bljr_go_api/api/model"
  "github.com/google/uuid"
  "golang.org/x/crypto/bcrypt"
)


func SeedingDB() {
	// Cek apakah tabel Accesses masih kosong
	var countAccess int64
	configs.DB.Model(&model.Access{}).Count(&countAccess)

	accesses := []model.Access{
		{AccessId: 1, AccessName: "users:create"}, // 0
		{AccessId: 2, AccessName: "users:read"}, // 1
		{AccessId: 3, AccessName: "users:update"}, // 2
		{AccessId: 4, AccessName: "users:delete"}, // 3
		{AccessId: 5, AccessName: "books:write"}, // 4
		{AccessId: 6, AccessName: "books:read"}, // 5
		{AccessId: 7, AccessName: "books:update"}, // 6
		{AccessId: 8, AccessName: "books:delete"}, // 7
	}

	for _, acc := range accesses {
		configs.DB.Save(&acc)
	}
	var countRole int64
	configs.DB.Model(&model.Role{}).Count(&countRole)

	if countRole == 0 {
		roles := []model.Role{
			{
				RoleId:   1,
				RoleName: "admin",
				Accesses: accesses,
			},
			{
				RoleId:   2,
				RoleName: "member",
				Accesses: []model.Access{accesses[1], accesses[3], accesses[5]},
			},
			{
				RoleId:   3,
				RoleName: "penulis",
				Accesses: []model.Access{accesses[5], accesses[3], accesses[6], accesses[4], accesses[7]},
			},
		}

		// Gunakan Create untuk data baru yang segar
		if err := configs.DB.Create(&roles).Error; err != nil {
			log.Println("Gagal seeding Roles & Relasi:", err)
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
				Roles:     []model.Role{adminRole}, 
			},
			{
				ID:        uuid.New(), 
				FirstName: "Renal",
				LastName:  "Tupai",
				Username:  "MasTUpai",
				Email:     "tupai@example.com",
				Password:  string(hashedPassword),
				Roles:     []model.Role{adminRole}, 
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
