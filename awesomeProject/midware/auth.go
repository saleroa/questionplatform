package midware

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func Auth(c *gin.Context) (id int, err error) {
	uid, err := c.Cookie("name")
	id, err = strconv.Atoi(uid)
	return
}
