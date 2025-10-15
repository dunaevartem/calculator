package main

import (
    "github.com/gin-gonic/gin"
    "strconv"
)

func main() {
    r := gin.Default()
    r.LoadHTMLGlob("templates/*.html")

    r.GET("/", func(c *gin.Context) {
        c.HTML(200, "index.html", gin.H{})
    })

    r.POST("/", func(c *gin.Context) {
        a := c.PostForm("a")
        b := c.PostForm("b")
        op := c.PostForm("op")

        if a == "" || b == "" || op == "" {
            c.HTML(400, "error.html", gin.H{"error": "Поля не заполнены"})
            return
        }

        numA, err := strconv.Atoi(a)
        numB, err := strconv.Atoi(b)
        if err != nil {
            c.HTML(400, "error.html", gin.H{"error": "Неверный формат чисел"})
            return
        }

        var result float64
        switch op {
        case "add":
            result = float64(numA) + float64(numB)
        case "sub":
            result = float64(numA) - float64(numB)
        case "mul":
            result = float64(numA) * float64(numB)
        case "div":
            if numB == 0 {
                c.HTML(400, "error.html", gin.H{"error": "Деление на ноль"})
                return
            }
            result = float64(numA) / float64(numB)
        default:
            c.HTML(400, "error.html", gin.H{"error": "Неизвестная операция"})
            return
        }

        c.HTML(200, "result.html", gin.H{"result": result})
    })

    r.Run(":8080")
}
