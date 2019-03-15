package main

/*
	Estan definidos MySQL y SQLite, para hacer pruebas.
	Se utiliza Gorm.
*/
import (
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

//Persona tipo de dato
//Para usar IDPersona en lugar de ID, se debe añadir gorm:"primary_key".
type Persona struct {
	IDPersona uint `json:"idPersona" gorm:"primary_key"`
	//ID			 	uint   `json:"idPersona"`
	DniPersona       uint   `json:"dniPersona"`
	CuilPersona      uint   `json:"cuilPersona"`
	NombresPersona   string `json:"nombresPersona"`
	ApellidosPersona string `json:"apellidosPersona"`
	SexoPersona      string `json:"sexoPersona"`
}

//Empleado tipo de dato
type Empleado struct {
	IDEmpleado     uint   `json:"idEmpleado"`
	PuestoEmpleado string `json:"puestoEmpleado"`
	Movil          string `json:"movilEmpleado"`
	NumeroLegajo   uint   `json:"numeroLegajoEmpleado"`
}

//UsuarioEmpleado tipo de dato
//Por convencion gorm usa ID (con mayusculas como clave primaria).
//Si a ID lo nombro con otra etiqueta, gorm no lo leerá. Si se hará el incremento en MySQL pero no sera reflejado en la devolucion de gorm.
//Base model definition gorm.Model, including fields ID, CreatedAt, UpdatedAt, DeletedAt, you could embed it in your model, or only write those fields you want
type UsuarioEmpleado struct {
	//gorm.Model
	IDUsuarioEmpleado         uint   `json:"idUsuarioEmpleado" gorm:"primary_key"`
	IDPersona                 uint   `json:"idPersona"`
	NombreUsuarioEmpleado     string `json:"nombreUsuarioEmpleado"`
	ContraseñaUsuarioEmpleado string `json:"contraseñaUsuarioEmpleado"`
}

//MovilEmpresa tipo de dato
type MovilEmpresa struct {
	IDMovilEmpresa            string `json:"idMovilEmpresa"`
	PatenteMovilEmpresa       string `json:"patenteMovilEmpresa"`
	ModeloMovilEmpresa        string `json:"modeloMovilEmpresa"`
	IdentificadorMovilEmpresa string `json:"identificadorMovilEmpresa"`
}

//Esta es la manera mas eficiente (que encontré) de utilizar la conexion a una db.
var db *gorm.DB

//Variable tipo error
var err error

func main() {

	db, _ = gorm.Open("mysql", "pablo:rocky@tcp(127.0.0.1:3306)/pruebas?charset=utf8&parseTime=True&loc=Local")
	//db, err = gorm.Open("mysql", "pablo:rocky@tcp(127.0.0.1:3306)/pruebas?charset=utf8&parseTime=True&loc=Local")
	//db, err = gorm.Open("sqlite3", "./gorm.db")
	//db, err = gorm.Open("mysql", "pablo:rocky@localhost/pruebas?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	//AutoMigrate crea una tabla con el nombre del struct de esta manera:
	//	· EmpleadoUsuario -> empleado_usuarios
	//Cambia todo a minisculas y a partird el segundo camelcase añade un guion bajo y añade al final (y solo al final) una s.
	//Aparece un problema, si el nombre es compuesto, y se quiere que cada parte se trasnforme en plural.
	//Dada la configuracion por defecto, solo se agregará una s al final del nombre.
	//Para evitar estos problemas se pude utilizar el siguiente comando, donde se espefica el nombre de la tabla deaseado, a partir de
	//una estructura elegida.
	//db.AutoMigrate(&Persona{})
	//db.Table("usuarios_empleados").AutoMigrate(&UsuarioEmpleado{})

	//Tb crear una table con un nombre dado.
	//db.Table("usuarios_empleados").CreateTable(&UsuarioEmpleado{})

	//	Default returns an Engine instance with the Logger and Recovery middleware already attached.
	router := gin.Default()
	router.GET("/signin", GetUsuario)
	router.POST("/crear-persona", CrearPersona)
	router.POST("/crear-usuario-empleado", CrearUsuarioEmpleado)

	router.Run(":8887")
}

/*
	GetUsuario: Context es la parte más importante de la Gin. Nos permite pasar variables entre middleware, administrar el flujo, validar el JSON de una solicitud y, por ejemplo, generar una respuesta JSON.
*/
func GetUsuario(c *gin.Context) {

	/*
		func (*Context) Param
		func (c *Context) Param(key string) string
		Param retorna el valor del parametro URL. Es un shortcut para c.Params.ByName(key)

		router.GET("/user/:id", func(c *gin.Context) {
			a GET request to /user/john
			id := c.Param("id") // id == "john"
		})
	*/
	id := c.Params.ByName("id")
	//id := c.Params("id")
	var usuario UsuarioEmpleado

	//? SELECT * FROM usuarios WHERE id = 'id';
	err := db.Where("id = ?", id).First(&usuario).Error

	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		/*
			JSON serializa la estructura dada como JSON en el cuerpo de la respuesta. También establece el tipo de contenido como "application / json".
		*/
		c.JSON(200, usuario)
	}
}

//curl -i -X  POST http://localhost:8887/crear-persona -d '{ "DNIPersona": 12344567, "CuilPersona":20192912, "NombresPersona":"Juan", "ApellidosPersona":"Perez", "SexoPersona":"masculino" }'
func CrearPersona(c *gin.Context) {
	var persona Persona
	c.BindJSON(&persona)
	db.Create(&persona)
	c.JSON(200, persona)
}

//ALTER TABLE `usuarios_empleados` AUTO_INCREMENT = 1
//curl -i -X  POST http://localhost:8887/crear-usuario-empleado -d '{ "IdPersona": 1, "NombreUsuarioEmpleado": "pablo", "ContraseñaUsuarioEmpleado":"lalala" }'
func CrearUsuarioEmpleado(c *gin.Context) {
	var usuario UsuarioEmpleado

	/*
		Bind comprueba el Content-Type para seleccionar un motor de enlace (binding engine) automáticamente. Según el encabezado "Content-Type" se utilizan diferentes enlaces:

		"application/json" --> JSON binding
		"application/xml"  --> XML binding

		de lo contrario -> devuelve un error. Analiza el cuerpo de la solicitud como JSON si Content-Type == "application / json" usa JSON o XML como entrada JSON. Decodifica la carga json en la estructura especificada como puntero. Escribe un error 400 y establece el encabezado de tipo de contenido "texto / plano" en la respuesta si la entrada no es válida.
	*/
	// 	BindJSON is a shortcut for c.MustBindWith(obj, binding.JSON).
	c.BindJSON(&usuario)

	//user := User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}
	//db.NewRecord(user) // => returns `true` as primary key is blank
	//db.Create(&user)
	//db.NewRecord(user) // => return `false` after `user` created

	//db.Create(&usuario) sería suficiente si la configuacion por defecto de gorm no hiciese plural solo los ultimos terminos.
	//quizas se pueda configurar eso, pero es mejor atajar el problema de esta forma para que corra en una instalacion estandar.
	db.Table("usuarios_empleados").Create(&usuario)
	c.JSON(200, usuario)
}

// ! Alets
// ?
// *
