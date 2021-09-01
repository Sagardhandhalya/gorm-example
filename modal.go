package main

import (
	"gorm.io/gorm"
)

/*
validation with gorm
1. Name :
	- col name shoud be user_name
	- size 255 byte
	- not null

2. Age :
	- not null
	- have to be < 100

3. Is Admin :
	- type havr to bool
	- there other tags refer https://gorm.io/docs/models.html

4. CreditCard
	- ont to one relation with credit card

5. Cars
	- ont to many relation with car
	- will not creat column just a back reference

6. Languages
	- many to many relation with language
	- just a back reference
*/

type User struct {
	gorm.Model
	Name       string     `gorm:"size:255;not null"`
	Age        uint8      `gorm:"not null;check:age<100"`
	IsAdmin    bool       `gorm:"type:bool"`
	CreditCard CreditCard `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Cars       []Car      `gorm:"constraint:OnUpdate:SET NULL,OnDelete:SET NULL;"`
	Languages  []Language `gorm:"many2many:user_languages;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

/*
====== ONE TO ONE =====
one to one relation between user and credit card
every user has one credit card
so every row in credit card table will contain one userId.
put credit card in user modal and put userId uint in creditcard modal.
if you want to change foreign key or reference you will do it in User modal
by adding this  `gorm:"foreignkey:UserName;references:name"`
*/

type CreditCard struct {
	gorm.Model
	CardNo   string `gorm:"size:12;not null;unique"`
	UserId   uint
	UserName string `gorm:"size:255;not null"`
}

/*
======= ONE TO MANY =======
one relation between user and car
so each row in car will have a user id
add car array i user and add user id in car modal
you can change foreign key and reference same as one to one.
*/

type Car struct {
	gorm.Model
	Name   string `gorm:"size:255"`
	UserId uint
}

/*
====== MANY TO MANY =======
many to many relationsheep between user and language because one user can speak many language
and we can fetch all user who speack hindi.
To make this relation gorm will add new table that will map user and language.
in user add Languages []languages `gorm:"many2many:user_languages"`
in language add Users []Users `gorm:"many2many:user_languages"`

*/

type Language struct {
	gorm.Model
	Name  string `gorm:"size:50;unique"`
	Users []User `gorm:"many2many:user_languages"`
}
