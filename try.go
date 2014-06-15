package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func loadNameWithSynset(synset string) string {
	db, err := sql.Open("sqlite3", "wnjpn.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sql := "select name from synset where synset=?;"
	stmt, err := db.Prepare(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	var name string
	err = stmt.QueryRow(synset).Scan(&name)

	if err != nil {
		log.Fatal(err, ": DBにwordidが見つかりません")
	}
	return name
}

func loadSensesWithWordid(wordid int) []string {
	db, err := sql.Open("sqlite3", "wnjpn.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sql := "select synset from sense where wordid=?;"
	stmt, err := db.Prepare(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(wordid)
	defer rows.Close()
	synsets := []string{}
	for i := 0; rows.Next(); i++ {
		var synset string
		rows.Scan(&synset)
		synsets = append(synsets, synset)
	}

	if err != nil {
		log.Fatal(err, ": DBにwordidが見つかりません")
	}
	return synsets
}

func loadWordsWithLemmma(lemma string) int {
	db, err := sql.Open("sqlite3", "wnjpn.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sql := "select wordid from word where lemma = ?;"
	stmt, err := db.Prepare(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	var wordid int
	err = stmt.QueryRow("夢").Scan(&wordid)

	if err != nil {
		log.Fatal(err, ": DBに単語が見つかりません")
	}
	return wordid
}

func main() {
	lemma := "夢"
	wordid := loadWordsWithLemmma(lemma)
	synsets := loadSensesWithWordid(wordid)
	for i := 0; i < len(synsets); i++ {
		fmt.Println(loadNameWithSynset(synsets[i]))
	}
	fmt.Println(synsets)
}
