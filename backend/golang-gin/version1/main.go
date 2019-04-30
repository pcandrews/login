package main

/*
	Version con gorm parcial.
		no uso:
			gorm.Model
			automigrate
*/

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

type Persona struct {
	//  gorm.Model
	IDPersona  uint64 `json:"idPersona" xml:"idPersona" form:"idPersona" gorm:"primary_key"`
	DniPersona uint64 `json:"dniPersona" form:"dniPersona" binding:"required"`

	CreadoEn      time.Time `gorm:"-"`
	ActualizadoEn time.Time `gorm:"-"`
	EliminadoEn   time.Time `gorm:"-"`

	CuilPersona          uint64 `json:"cuilPersona" form:"cuilPersona" binding:"required"`
	NombresPersona       string `json:"nombresPersona" form:"nombresPersona" binding:"required"`
	ApellidosPersona     string `json:"apellidosPersona" form:"apellidosPersona" binding:"required"`
	SexoPersona          string `json:"sexoPersona" form:"sexoPersona" binding:"required"`
	ObservacionesPersona string `json:"observacionesPersona" form:"observacionesPersona"`
}

type Empleado struct {
	// `gorm:"-"` excluye el campo
	Persona    `gorm:"-"`
	IDEmpleado uint64 `json:"idEmpleado" form:"idEmpleado" gorm:"primary_key"`
	IDPersona  uint64 `json:"idPersona" form:"idPersona gorm:"foreignkey:IDPersona"`

	CreadoEn      time.Time `gorm:"-"`
	ActualizadoEn time.Time `gorm:"-"`
	EliminadoEn   time.Time `gorm:"-"`

	PuestoEmpleado       string `json:"puestoEmpleado" form:"puestoEmpleado" binding:"required"`
	MovilEmpleado        string `json:"movilEmpleado" form:"movilEmpleado"`
	NumeroLegajoEmpleado uint64 `json:"numeroLegajoEmpleado" form:"numeroLegajoEmpleado" binding:"required"`
	CelularEmpleado      uint64 `json:"celularEmpleado" form:"celularEmpleado"`
}

// CRUD Tabla Empleados

// Create (CRUD)
func (e *Empleado) Crear() {
	db.Exec("ALTER TABLE `personas` AUTO_INCREMENT = 1;")
	db.Save(&e.Persona)
	db.Exec("ALTER TABLE `empleados` AUTO_INCREMENT = 1;")
	e.IDPersona = e.Persona.IDPersona
	db.Save(&e)
}

// Read (CRUD)
func (e *Empleado) Leer() {
	var p Persona
	db.First(&e, e.IDEmpleado)
	db.First(&p, e.IDPersona)
	e.Persona = p

	db.Find(&e)

	//fmt.Println("Leer:%d", e.IDEmpleado)

	/*
		out, err := json.Marshal(e)
		if err != nil {
			panic(err)
		}

		fmt.Println("\n\n\n\n\nEmpleado: " + string(out) + "\n\n\n\n\n")
	*/

	//return e
}

// Read (CRUD)
func (e *Empleado) LeerTodos() []Empleado {
	var todosE []Empleado
	var todosPE []Persona

	db.Exec("SELECT * FROM empleados e LEFT JOIN personas p ON e.id_persona = p.id_persona ORDER BY e.id_empleado ASC").Find(&todosE)

	db.Exec("SELECT * FROM empleados e LEFT JOIN personas p ON e.id_persona = p.id_persona ORDER BY e.id_empleado ASC").Find(&todosPE)

	for i := 0; i < len(todosE); i++ {
		todosE[i].Persona = todosPE[i]
	}

	//c.JSON(http.StatusOK, todosE)
	//c.JSON(http.StatusOK, todosPE)

	return todosE
}

// Update (CRUD)
func (e *Empleado) Actualizar() {
	var datetime = time.Now()

	datetime.Format(time.RFC3339)

	e.Leer()

	//fmt.Println(e)

	db.Exec("UPDATE empleados SET actualizado_en=? WHERE empleados.id_empleado=?", datetime, e.IDEmpleado).Find(&e)
}

// Delete (Softdelete) (CRUD)
func (e *Empleado) Eliminar() {
	var datetime = time.Now()

	datetime.Format(time.RFC3339)

	e.Leer()

	// Enable Logger, show detailed log
	//db.LogMode(true)

	// Debug a single operation, show detailed log for this operation
	//db.Debug().Where("name = ?", "jinzhu").First(&User{})

	//db.Debug().Exec("UPDATE empleados SET eliminado_en=? WHERE empleados.id_empleado=?", datetime, e.IDEmpleado).Find(&e)

	db.Exec("UPDATE empleados SET eliminado_en=? WHERE empleados.id_empleado=?", datetime, e.IDEmpleado).Find(&e)

	// Disable Logger, don't show any log even errors
	//db.LogMode(false)

}

func main() {
	var err error

	conex := "pablo:rocky@tcp(127.0.0.1:3306)/pruebas?charset=utf8&parseTime=True&loc=Local"
	db, err = gorm.Open("mysql", conex)

	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	defer db.Close()

	router := gin.Default()

	router.LoadHTMLGlob("pla/*")

	router.GET("/formulario-empleado", FormularioCrearEmpleado)
	router.POST("/mostrar-empleado", CrearEmpleado)
	router.GET("/obterner-empleado/:id", ObtenerEmpleado)
	router.GET("/obterner-todos-empleados/", ObtenerTodosLosEmpleados)
	router.PUT("/actualizar-empleado/:id", ActualizarEmpleado)
	router.DELETE("/eliminar-empleado/:id", EliminarEmpleado)

	/*
		sudo lsof -n -i :8080
		kill -9 <PID>
		sudo killall -9 main3
	*/
	router.Run(":8080")
}

/*
	Get
	http://localhost:8080/formulario-empleado
*/
func FormularioCrearEmpleado(c *gin.Context) {
	//var resultado gin.H
	//var err error

	c.HTML(http.StatusOK, "formulario-empleado.tmpl", nil)
}

//Post
func CrearEmpleado(c *gin.Context) {
	var resultado gin.H
	var err error
	var e Empleado

	if c.ShouldBind(&e) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	e.Crear()

	resultado = gin.H{
		"resultado": e,
		"cantidad":  1,
	}

	c.JSON(http.StatusOK, resultado)
}

/*
	Get
	http://localhost:8080/obterner-empleado/1
*/
func ObtenerEmpleado(c *gin.Context) {
	var resultado gin.H
	var err error
	var e Empleado
	var idParam string

	idParam = c.Param("id")

	e.IDEmpleado, err = strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		log.Fatalln(err)
	}
	e.Leer()

	resultado = gin.H{
		"resultado": e,
		"cantidad":  1,
	}

	c.JSON(http.StatusOK, resultado)

	/*
		c.HTML(http.StatusOK, "mostrar-datos-empleado.tmpl", gin.H{
		"nombres": e.Leer(e.IDEmpleado).Persona.NombresPersona})
	*/
}

/*
	Get
	http://localhost:8080/obterner-empleado/1
*/
func ObtenerTodosLosEmpleados(c *gin.Context) {
	var resultado gin.H
	//var err error
	var e Empleado
	var todosE []Empleado

	todosE = e.LeerTodos()

	resultado = gin.H{
		"resultado": todosE,
		"cantidad":  len(todosE),
	}

	c.JSON(http.StatusOK, resultado)
}

/*
	Get
	http://localhost:8080/eliminar-empleado/1
*/
func ActualizarEmpleado(c *gin.Context) {
	var resultado gin.H
	var err error
	var e Empleado
	var idParam string

	idParam = c.Param("id")

	e.IDEmpleado, err = strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		log.Fatalln(err)
	}

	e.Actualizar()

	resultado = gin.H{
		"resultado": e,
		"cantidad":  1,
	}

	c.JSON(http.StatusOK, resultado)
}

/*
	Get
	http://localhost:8080/eliminar-empleado/1
	Softdelete
*/
func EliminarEmpleado(c *gin.Context) {
	var resultado gin.H
	var err error
	var e Empleado
	var idParam string

	idParam = c.Param("id")

	e.IDEmpleado, err = strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		log.Fatalln(err)
	}

	e.Eliminar()

	resultado = gin.H{
		"resultado": e,
		"cantidad":  1,
	}

	c.JSON(http.StatusOK, resultado)

	/*c.HTML(http.StatusOK, "mostrar-datos-empleado.tmpl", gin.H{
	"nombres": e.Leer(e.IDEmpleado).Persona.NombresPersona})*/
}
