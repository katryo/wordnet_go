package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func loadNameWithSynsetFromSynset(synset string) string {
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

func loadSynsetsWithWordidFromSense(wordid string) []string {
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
		log.Fatal(err, ": senseにwordidが見つかりません")
	}
	return synsets
}

func loadWordsWithLemmmaFromWord(lemma string) []string {
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
	rows, err := stmt.Query(lemma)
	defer rows.Close()
	var wordids []string
	for i := 0; rows.Next(); i++ {
		var wordid string
		rows.Scan(&wordid)
		wordids = append(wordids, wordid)
	}

	if err != nil {
		log.Fatal(err, ": wordにwordidが見つかりません")
	}
	return wordids
}

func loadSynsetsWithSynsetFromSynlink(synset string) []string {
	db, err := sql.Open("sqlite3", "wnjpn.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sql := "select synset2 from synlink where synset1=? and link='hype';"
	stmt, err := db.Prepare(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(synset)
	defer rows.Close()
	synsets := []string{}
	for i := 0; rows.Next(); i++ {
		var synset2 string
		rows.Scan(&synset2)
		synsets = append(synsets, synset2)
	}

	if err != nil {
		log.Fatal(err, ": synlinkにwordidが見つかりません")
	}
	return synsets
}

func loadLemmaWithWordidFromWord(wordid string) string {
	db, err := sql.Open("sqlite3", "wnjpn.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sql := "select lemma from word where wordid = ?;"
	stmt, err := db.Prepare(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	var lemma string
	err = stmt.QueryRow(wordid).Scan(&lemma)

	if err != nil {
		log.Fatal(err, ": wordに単語が見つかりません")
	}
	return lemma
}

func loadWordidWithSynsetFromSense(synset string) string {
	db, err := sql.Open("sqlite3", "wnjpn.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sql := "select wordid from sense where synset = ?;"
	stmt, err := db.Prepare(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	var wordid string
	err = stmt.QueryRow(synset).Scan(&wordid)

	if err != nil {
		log.Fatal(err, ": senseに単語が見つかりません", synset)
	}
	return wordid
}
func printSynlinksRecursively(wordids []string, depth int) {
	for j := 0; j < len(wordids); j++ {
		synsets := loadSynsetsWithWordidFromSense(wordids[j])
		synlinksSynsets := []string{}
		for i := 0; i < len(synsets); i++ {
			synlinksSynsets = append(synlinksSynsets, loadSynsetsWithSynsetFromSynlink(synsets[i])...)
		}
		deeperWordids := []string{}
		if len(synlinksSynsets) != 0 {
			for i := 0; i < depth+1; i++ {
				fmt.Print("  ")
			}
			lemma := loadLemmaWithWordidFromWord(wordids[j])
			fmt.Println("  ", lemma)
			for i := 0; i < len(synlinksSynsets); i++ {
				newWordid := loadWordidWithSynsetFromSense(synlinksSynsets[i])
				deeperWordids = append(deeperWordids, newWordid)
			}
			printSynlinksRecursively(deeperWordids, depth+1)
		}
	}
}

func main() {
	lemma := "夢"
	wordids := loadWordsWithLemmmaFromWord(lemma)
	printSynlinksRecursively(wordids, 0)
}
