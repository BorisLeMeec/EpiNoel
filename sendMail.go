package main

import (
	"log"
	"os"
	"strings"

	gomail "gopkg.in/gomail.v2"
)

type student struct {
	email     string
	firstName string
	lastName  string
}

func sendMail(sender, getter []string) bool {

	sSender := student{sender[indexEmail], strings.Title(strings.Split(strings.Split(sender[indexEmail], "@")[0], ".")[0]), strings.Title(strings.Split(strings.Split(sender[indexEmail], "@")[0], ".")[1])}
	sGetter := student{getter[indexEmail], strings.Title(strings.Split(strings.Split(getter[indexEmail], "@")[0], ".")[0]), strings.Title(strings.Split(strings.Split(getter[indexEmail], "@")[0], ".")[1])}

	m := gomail.NewMessage()
	m.SetHeader("From", "bde.epitech.rennes@gmail.com")
	m.SetHeader("To", sSender.email)
	m.SetHeader("Subject", "[BDE] You have a mission!")
	m.SetBody("text/html", `<html><body>Bonjour <b>`+sSender.firstName+` `+sSender.lastName+`</b>,<br><br>Les fêtes de Noël approchent à grand pas, et avant la fin de la journée un magnifique sapin devrait faire son apparition dans la grande salle. À toi de le décorer ;)<br><br>
	Afin de ne pas couper la tradition,comme tous les ans, ton BDE préféré met en place un <b>échange de cadeaux entre étudiants</b>.<br><br>
	Le principe est tout ce qu'il y a de plus simple. Tu vas devoir offrir un cadeau (<b>maximum 5€</b>) à <b>`+sGetter.firstName+` `+sGetter.lastName+`</b>, et lui, de son côté devra en offrir un à un étudiant également (pas forcément toi).<br>
	Attention, c'est plus marrant si on ne sait pas qui doit offrir à qui, par conséquent garde la surprise pour le jour du déballage (16 décembre au soir) ;)<br><br>
	Qui participera ? Tek1, Tek3, AERs, BDE et ADM !<br>
	Nous procéderons tous ensemble au débalage avant que tout le monde parte en vacances. En attendant, cela sera l'occasion de voir plein de beaux cadeaux au pied du sapin.<br>
	L'essentiel est le cadeau, tout présent sera accepté (collier de pâtes, moulures en pate à sel...), tant que celui ci n'est pas irrespectueux.<br><br>
	Le jour du déballage nous organiserons également une grosse raclette des familles, et, si tu souhaites continuer la soirée, on partira se réchauffer le coeur et le gosier dans les bars de Rennes jusqu'à pas d'heures, mais on te rappelle tout ça dans le prochain mail ;)<br><br>
	<b>DONC</b>, récapitulatif, tu dois offrir un cadeau à <a href="https://intra.epitech.eu/user/`+sGetter.email+`/"><b>`+sGetter.firstName+` `+sGetter.lastName+`</b></a>, d'une valeur de <b>5€ max</b>, en le déposant au pied du sapin avant le 16 décembre au soir.<br><br>
	Bonne semaine à tous, et on se revoit pour le débalage des cadeaux + la grosse raclette + soirée du 16 décembre :)<br><br>
	<b>Boris Le Méec</b><br>Président de votre BDE <3</body></html>`)
	log.Println("Start sending mail to " + sSender.email + " => (" + sGetter.email + ")")

	d := gomail.NewPlainDialer("smtp.gmail.com", 587, "bde.epitech.rennes@gmail.com", os.Getenv("PASSWD_GMAIL_BDE"))
	if err := d.DialAndSend(m); err != nil {
		log.Fatal(err)
	}
	log.Println("Mail sent with success !")
	return true
}
