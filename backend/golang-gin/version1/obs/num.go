package main

import (
	"fmt"
	"strconv"
	"time"
)

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

func main() {

	str := "1234"
	//var num uint64
	var err error
	var e Empleado

	/** converting the str1 variable into an unsigned int using ParseUint method */
	e.IDEmpleado, err = strconv.ParseUint(str, 10, 64)
	if err == nil {
		fmt.Printf("Type: %T \n", e.IDEmpleado)
		fmt.Println(e.IDEmpleado)
	}
}
