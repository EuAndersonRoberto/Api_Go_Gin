package models

import "gorm.io/gorm"

type Aluno struct {
	gorm.Model        // Este comando nos trás uma struct do GORM que já nos apresenta: id, create, update e delete.
	Nome       string `json:"nome"`
	CPF        string `json:"cpf"`
	RG         string `json:"rg"`
}

//var Alunos []Aluno //Aqui criamos uma variável Alunos que comtém um slice da struct "Aluno".
