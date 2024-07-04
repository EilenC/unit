package main

import (
	"fmt"
	"gee/core"
	"net/http"
	"os"
)

func permute(nums []int) [][]int {
	var res, path = make([][]int, 0), make([]int, 0)
	var used = make([]bool, len(nums))

	var dfs func()
	dfs = func() {
		if len(path) == len(nums) {
			var temp = make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}

		for i := range nums {
			if used[i] {
				continue
			}
			path = append(path, nums[i])
			used[i] = true
			dfs()
			// 回溯的过程中，将当前的节点从 path 中删除
			path = path[:len(path)-1]
			used[i] = false
		}
	}

	dfs()

	return res
}

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
	os.Exit(1)
	fr := core.New()
	fr.GET("/", func(c *core.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})
	fr.GET("/hello", func(c *core.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	fr.GET("/hello/:name", func(c *core.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	fr.GET("/assets/*filepath", func(c *core.Context) {
		c.JSON(http.StatusOK, core.H{"filepath": c.Param("filepath")})
	})

	fr.POST("/login", func(c *core.Context) {
		c.JSON(http.StatusOK, core.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	_ = fr.Run(":9999")
}
