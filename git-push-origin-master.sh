#!/bin/bash
# sube los cambios a git
# para usar espacios en el commit usar comillas simples
# sino, no dejar espacios
git add . &&
git commit -m "$1" &&
git push -u origin master
