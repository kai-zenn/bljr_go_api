package utils

import (
	"github.com/google/uuid"
	"github.com/kai-zenn/bljr_go_api/api/configs"
)

func HasAccess(userId uuid.UUID, AccessName string) bool {
  var count int64
  

  query := `
    SELECT COUNT(*)
    FROM users u
    JOIN user_role ur ON u.id = ur.user_id
    JOIN role_access ra ON ur.role_id = ra.role_id
    JOIN accesses a ON ra.access_id = a.access_id
    WHERE a.access_name = ? AND u.id = ?
  `

  if err := configs.DB.Raw(query, AccessName, userId.String()).Scan(&count).Error; err != nil {
    return false
  }

  return count > 0
}
