package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/kai-zenn/bljr_go_api/api/configs"
	"github.com/kai-zenn/bljr_go_api/api/controller"
	"github.com/kai-zenn/bljr_go_api/api/middlewares"
	"github.com/kai-zenn/bljr_go_api/api/migration"
	"github.com/kai-zenn/bljr_go_api/api/model"
	"github.com/stretchr/testify/assert"
)

func SetupTestRoutes(r *gin.Engine) {
  // Public routes (accessible by anyone)
   r.POST("/login", controller.LoginUserHandler)
   r.POST("/register", controller.RegisterUserHandler)

   // Protected routes (secured with RBAC)
   authGroup := r.Group("/")
   authGroup.Use(middlewares.AuthMiddleware())

   // Role-based access control for specific routes
   authGroup.GET("/users", middlewares.RBACMiddleware("read"), controller.GetUsers)
   authGroup.GET("/users/:id", middlewares.RBACMiddleware("read"), controller.GetUserById)
   authGroup.PUT("/users/:id", middlewares.RBACMiddleware("update"), controller.UpdateUser)
	
}

func SetupTestRouter() *gin.Engine {
  gin.SetMode(gin.TestMode)
  r := gin.Default() 

  os.Setenv("JWT_SECRET", "rahasia_jwt_kamu_di_env")

	configs.InitDB()

	configs.DB.Exec("TRUNCATE TABLE user_role, role_access, users, roles, accesses CASCADE;")
	configs.DB.AutoMigrate(&model.Access{}, &model.Role{}, &model.User{})
	
	migration.SeedingDB() 
	migration.SeedUsers()

  SetupTestRoutes(r)
  return r
}

// 1. Tes Rute Publik: Login Gagal (Password Salah)
func TestLoginAPI_Failure(t *testing.T) {
	configs.InitDB() 
	router := SetupTestRouter()

	loginData := map[string]string{
		"email":    "admin@example.com",
		"password": "password_salah_nih",
	}
	jsonValue, _ := json.Marshal(loginData)

	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Validasi hasil menggunakan library testify/assert
	// Karena password salah, harusnya dapet status 401 Unauthorized atau 400 Bad Request (tergantung controllermu)
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

// 2. Tes Rute Protected: Akses /users Tanpa Token JWT (Harus Tertolak)
func TestGetUsers_WithoutToken(t *testing.T) {
	router := SetupTestRouter()

	// Tembak langsung ke /users tanpa pasang header Authorization
	req, _ := http.NewRequest("GET", "/users", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Karena tidak bawa token, AuthMiddleware kamu harusnya nembak 401 Unauthorized
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}


type LoginResponse struct {
	Token string `json:"token"`
}

func TestAdminAccessUsersAPI_Success(t *testing.T) {
	configs.InitDB()
	router := SetupTestRouter()

	adminCredentials := map[string]string{
		"email":    "admin@example.com",
		"password": "rahasia124",
	}
	jsonLogin, _ := json.Marshal(adminCredentials)

	reqLogin, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonLogin))
	reqLogin.Header.Set("Content-Type", "application/json")
	
	wLogin := httptest.NewRecorder()
	router.ServeHTTP(wLogin, reqLogin)

	assert.Equal(t, http.StatusOK, wLogin.Code)

	var loginResp LoginResponse
	err := json.Unmarshal(wLogin.Body.Bytes(), &loginResp)
	assert.Nil(t, err)
	assert.NotEmpty(t, loginResp.Token) 
	
	// LANGKAH 2: Tembak API /users menggunakan Token Admin tersebut
	reqGetUsers, _ := http.NewRequest("GET", "/users", nil)
	
	// Pasang token di Header Authorization (Standar JWT Bearer Token)
	reqGetUsers.Header.Set("Authorization", "Bearer "+loginResp.Token)

	wGetUsers := httptest.NewRecorder()
	router.ServeHTTP(wGetUsers, reqGetUsers)

	// Karena dia Admin (punya access 'read'), maka statusnya HARUS 200 OK (diizinkan)
	assert.Equal(t, http.StatusOK, wGetUsers.Code)
}

func TestMemberAccessUsersAPI_Forbidden(t *testing.T) {
	configs.InitDB()
	router := SetupTestRouter()

	memberCredentials := map[string]string{
		"email":    "kusnadi@smkn46.com",
		"password": "rahasia124",
	}
	jsonLogin, _ := json.Marshal(memberCredentials)

	reqLogin, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonLogin))
	reqLogin.Header.Set("Content-Type", "application/json")
	
	wLogin := httptest.NewRecorder()
	router.ServeHTTP(wLogin, reqLogin)

	var loginResp LoginResponse
	json.Unmarshal(wLogin.Body.Bytes(), &loginResp)

	reqGetUsers, _ := http.NewRequest("GET", "/users", nil)
	reqGetUsers.Header.Set("Authorization", "Bearer "+loginResp.Token)

	wGetUsers := httptest.NewRecorder()
	router.ServeHTTP(wGetUsers, reqGetUsers)

	// Karena member tidak punya izin, RBACMiddleware harus me-return 403 Forbidden!
	assert.Equal(t, http.StatusForbidden, wGetUsers.Code)
}
