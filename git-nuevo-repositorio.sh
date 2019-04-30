#!/bin/bash
# 1: situarse en la carpeta del proyecto.
# 2: añadir nombre del proyecto, si lleva espacios escribir el nombre entre comillas simples
git init  &&
git remote add origin git@github.com:pcandrews/$1.git
