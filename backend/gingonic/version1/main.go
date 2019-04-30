package main

/*
	Estan definidos MySQL y SQLite, para hacer pruebas.
	Se utiliza Gorm.
	Cuando se agrega una guion bajo delante de las dependencias, por ejemplo:
		_ "github.com/go-sql-driver/mysql"
	Se desactivan los warnings por no uso de la denpendencia, o sea, si no se utliza, no se indicará que no se utiliza.
*/
import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

//Persona tipo de dato
//Para usar IDPersona en lugar de ID, se debe añadir gorm:"primary_key".
/*type Persona struct {
	IDPersona uint `json:"idPersona" gorm:"primary_key"`
	//ID			 	 uint 	`json:"idPersona"`
	DniPersona           uint   `json:"dniPersona"`
	CuilPersona          uint   `json:"cuilPersona"`
	NombresPersona       string `json:"nombresPersona"`
	ApellidosPersona     string `json:"apellidosPersona"`
	SexoPersona          string `json:"sexoPersona"`
	ObservacionesPersona string `json:"observacionesPersona"`
}*/

type Persona struct {
	IDPersona uint `form:"idPersona" gorm:"primary_key"`
	//ID			 	 uint 	`form:"idPersona"`
	DniPersona           uint   `form:"dniPersona"`
	CuilPersona          uint   `form:"cuilPersona"`
	NombresPersona       string `form:"nombresPersona"`
	ApellidosPersona     string `form:"apellidosPersona"`
	SexoPersona          string `form:"sexoPersona"`
	ObservacionesPersona string `form:"observacionesPersona"`
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

	//Default returns an Engine instance with the Logger and Recovery middleware already attached.
	//Engine is the framework's instance, it contains the muxer, middlewares and configuration settings. Create an instance of Engine, by using New() or Default()
	router := gin.Default()

	//LoadHTMLFiles loads a slice of HTML files and associates the result with HTML renderer.
	//Es necesario invocar a los archivo template o html para que se puedan utilizar.
	router.LoadHTMLFiles("formulario-persona.tmpl")

	router.GET("/", Inicio)

	router.GET("/signin", GetUsuario)
	router.POST("/crear-persona", CrearPersona)
	router.POST("/crear-usuario-empleado", CrearUsuarioEmpleado)

	router.GET("/formulario-persona", GetFormularioPersona)
	router.POST("/mostrarDatosPersona", PostMostrarDatosPersona)

	/*
		sudo lsof -n -i :8080
		kill -9 <PID>
	*/
	router.Run(":8080")
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
	db.Exec("ALTER TABLE `personas` AUTO_INCREMENT =1;")
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

	db.Exec("ALTER TABLE `usuarios_empleados` AUTO_INCREMENT =1;")

	//db.Create(&usuario) sería suficiente si la configuacion por defecto de gorm no hiciese plural solo los ultimos terminos.
	//quizas se pueda configurar eso, pero es mejor atajar el problema de esta forma para que corra en una instalacion estandar.
	db.Table("usuarios_empleados").Create(&usuario)
	c.JSON(200, usuario)
}

/**
 * Mi metodo
 **Importante
 *! deprecated
 *? should this method be exposed?
 * TODO: abc
 * @param myparam parametro
 */

/*

	http://www.forosdelweb.com/f18/aporte-entendiendo-las-cabeceras-post-get-put-delete-920883/

	Relacion CRUD (Crear, recuperar, actualizar y eliminar) con la semántica con "Representational State Transfer" (REST, verbos definidos por la especificación de HTTP: GET, PUT, POST, DELETE, HEAD, etc.)

	Mientras que recuperar realmente se asigna a una solicitud GET HTTP, y también eliminar realmente se asigna a una operación HTTP DELETE, el mismo no puede decirse de crear y PUT o la actualización y POST. En algunos casos, crear es PUT, pero en otros casos se debe emplear POST. Del mismo modo, en algunos casos, actualización puede ser POST, mientras que en otros PUT.

	La esencia de la cuestión se reduce a un concepto conocido como *idempotencia. Una operación es idempotente si hay una secuencia de dos o más del mismo resultado de operación en el mismo estado de recurso, al igual que se trabaja con una clase que implementa singleton. De acuerdo con la especificación HTTP 1.1, GET, HEAD, PUT y DELETE son idempotentes, mientras que POST no lo es. Es decir, una secuencia de varios intentos de poner los datos a una URL se traducirá en el estado de los recursos lo mismo que un solo intento de poner los datos a esa URL, pero el mismo no se puede decir de una petición POST. Por ello, un navegador siempre le aparece un cuadro de diálogo de advertencia cuando se hace más de una petición en un formulario POST.

	? · Crear = PUT si y sólo si va a enviar todo el contenido del recurso especificado (dirección URL)
	? · Crear = POST si va a enviar un comando al servidor para crear un subordinado del recurso especificado mediante algún algoritmo del lado del servidor
	? · Recuperar = GET
	? · Actualizar = PUT si y sólo si va a actualizar el contenido completo del recurso especificado
	? · Actualizar = POST si usted está solicitando el servidor para actualizar uno o más subordinados del recurso especificado
	? · Eliminar = DELETE

	* idempotente = Se refiere a una operación que produce los mismos resultados sin importar cuántas veces se lleva a cabo. Por ejemplo, si una solicitud para eliminar un archivo se completa con éxito de un programa, todas las solicitudes posteriores a eliminar ese archivo de otros programas devolverán el mensaje de confirmación del primero como éxito si la función de borrado es idempotente. En una función que no es idempotente, un error se devuelve para la segunda y subsiguientes peticiones que indica que el archivo no estaba allí, y que la condición de error puede provocar que el programa se detuviera. Si todo lo que se desea es garantizar un determinado archivo se ha eliminado, una función idempotente de eliminar devolvería el resultado mismo, éxito, no importa cuántas veces se ha ejecutado para el mismo archivo.

	En todo esto se refiere a cuando por ejemplo, tenemos un formulario y ese formulario no vemos que haya una acción, continuamos presionando varias veces el botón de "submit" hasta que vemos resultado. Eso lo evitamos verificando la misma solucitud que se sometió consultando si ya existe el resultado en la base de datos. O cuando refrescamos la pantalla y si usamos POST vemos el cuadro de advertencia de que si desea enviar los datos nuevamente, mientras que PUT no aparece, sino que te muestra el resultado de la primera vez.
*/

func Inicio(c *gin.Context) {
	c.String(http.StatusOK, "Inicio")
}

/*
	GET (equivalente a READ de CRUD)
	El método GET se emplea para leer una representación de un resource. En caso de respuesta positiva (200 OK), GET devuelve la representación en un formato concreto: HTML, XML, JSON o imágenes, JavaScript, CSS, etc. En caso de respuesta negativa devuelve 404 (not found) o 400 (bad request). Por ejemplo en la carga de una página web, primero se carga la url solicitada:

	GET php.net/docs HTTP/1.1
	En este caso devolverá HTML. Y después los demás resources como CSS, JS, o imágenes:

	GET php.net/images/logo.png HTTP/1.1
	Los formularios también pueden usarse con el método GET, donde se añaden los keys y values buscados a la URL del header:

	<form action="formget.php" method="get">
	Nombre: <input type="text" name="nombre"><br>
	Email: <input type="text" name="email"><br>
	<input type="submit" value="Enviar">
	</form>
	La URL con los datos rellenados quedaría así:

	GET ejemplo.com/formget.php?nombre=pepe&email=pepe%40ejemplo.com HTTP/1.1
*/
/*
	GetPersona: recupera datos desde un formulario.
*/
func GetFormularioPersona(c *gin.Context) {
	/*dniPersona := c.Param("dni")
	cuilPersona := c.Param("cuil")
	nombresPersona := c.Param("nombres")
	apellidosPersona := c.Param("apellidos")
	sexoPersona := c.Param("sexo")
	observacionesPersona := c.Param("observaciones")*/

	/*c.HTML(http.StatusOK, "formulario-persona.tmpl", gin.H{
	"dni":           dniPersona,
	"cuil":          cuilPersona,
	"nombres":       nombresPersona,
	"apellidos":     apellidosPersona,
	"sexo":          sexoPersona,
	"observaciones": observacionesPersona})*/

	c.HTML(http.StatusOK, "formulario-persona.tmpl", nil)
}

/*
	POST
	Aunque se puedan enviar datos a través del método GET, en muchos casos se utiliza POST por las limitaciones de GET. En caso de respuesta positiva devuelve 201 (created). Los POST requests se envían normalmente con formularios:

	<form action="formpost.php" method="post">
		Nombre: <input type="text" name="nombre"><br>
		Email: <input type="text" name="email"><br>
		<input type="submit" value="Enviar">
	</form>
	Rellenar el formulario anterior crea un HTTP request con la request line:

	POST /formpost.php HTTP/1.1
	El contenido va en el body del request, no aparece nada en la URL, aunque se envía en el mismo formato que con el método GET. Si se quiere enviar texto largo o cualquier tipo de archivo este es el método apropiado.

	Le siguen los headers, donde se incluyen algunas líneas específicas con información de los datos enviados:

	Content-Type: application/x-www-form-urlencoded
	Content-Length: 43
	A los headers le siguen una línea en blanco y a continuación el contenido del request:

	formpost.php?nombre=pepe&email=pepe%40ejemplo.com
*/
/*
	PostMostrarDatosPersona: muestra los datos enviados desde un formulario.
*/
func PostMostrarDatosPersona(c *gin.Context) {
	//var persona Persona // es equivalente a  persona := new(Persona)

	//c.BindJSON(&persona)
	//c.Bind(&persona)
	//c.ShouldBind(&persona)

	//fmt.Println(persona)
	//fmt.Printf("bla")

	//c.JSON(http.StatusOK, persona)
	//c.JSON(http.StatusOK, gin.H{"nombres": persona.NombresPersona})

	dniPersona := c.PostForm("dni")
	cuilPersona := c.PostForm("cuil")
	nombresPersona := c.PostForm("nombres")
	apellidosPersona := c.PostForm("apellidos")
	sexoPersona := c.PostForm("sexo")
	observacionesPersona := c.PostForm("observaciones")

	c.JSON(http.StatusOK, gin.H{
		"dni":           dniPersona,
		"cuil":          cuilPersona,
		"nombres":       nombresPersona,
		"apellidos":     apellidosPersona,
		"sexo":          sexoPersona,
		"observaciones": observacionesPersona})
}
