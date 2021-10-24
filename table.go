package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
)

//Q8
func indexAleatoire() uint64 {
	return rand.Uint64() % N
}

//Q8
func sortTableByLastColumn(table [][2]uint64) {
	sort.Slice(table, func(i, j int) bool {
		return table[i][1] < table[j][1]
	})

}

//Q8
func creerTable(width uint64, height uint64) [][2]uint64 {

	// ====
	fmt.Printf(">> Création de la table de largeur %d et de hauteur %d :\n", width, height)
	// ====

	var tabRes [][2]uint64

	for i := 0; i < int(height); i++ {
		var index = indexAleatoire()
		var newString = newString(index, width)
		var element = [2]uint64{index, newString}

		tabRes = append(tabRes, element)
		// ==== affichage du loading
		pourcentage := int((float64(i) / float64(height)) * 100)
		fmt.Printf(">> Création en cours : %d%%\n", pourcentage)
		// ====
	}
	sortTableByLastColumn(tabRes)

	fmt.Printf(">> Création terminée.\n")
	return tabRes
}

//Q9
func sauveTable(table [][2]uint64, width uint64, height uint64, filename string) error {
	//créer le fichier
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	//Entête du fichier
	f.WriteString(fmt.Sprintf("fonction de hachage : %s\nalphabet : %s\ntaille_min : %d\n taille_max : %d\nlargeur de la table : %d\nhauteur de la table : %d\n", hashMethod, alphabet, sizeMin, sizeMax, width, height))

	//écrire la table
	for i := 0; i < int(height); i++ {
		f.WriteString(fmt.Sprintf("%d : ", i))
		f.WriteString(fmt.Sprintf("%d %d ", table[i][0], table[i][1]))
		f.WriteString("\n")
	}

	return nil
}

//Q9
func ouvreTable(filename string) [][2]uint64 {
	f, _ := os.Open(filename)
	defer f.Close()

	sc := bufio.NewScanner(f)

	if sc.Scan() { //fonction de hachage
		line := sc.Text()
		hashMethod = strings.Split(line, " : ")[1]
	}
	if sc.Scan() { //alphabet
		line := sc.Text()
		alphabet = strings.Split(line, " : ")[1]
	}
	if sc.Scan() { //taille minimum
		line := sc.Text()
		sizeMinTemp, _ := strconv.Atoi(strings.Split(line, " : ")[1])
		sizeMin = uint64(sizeMinTemp)
	}
	if sc.Scan() { //taille maximale
		line := sc.Text()
		sizeMaxTemp, _ := strconv.Atoi(strings.Split(line, " : ")[1])
		sizeMax = uint64(sizeMaxTemp)
	}
	if sc.Scan() { //largeur
		line := sc.Text()
		widthTemp, _ := strconv.Atoi(strings.Split(line, " : ")[1])
		gWidth = uint64(widthTemp)
	}
	if sc.Scan() { //hauteur
		line := sc.Text()
		heightTemp, _ := strconv.Atoi(strings.Split(line, " : ")[1])
		gHeight = uint64(heightTemp)
	}

	//Calcul de la longueur de l'alphabet
	lenAlphabet = uint64(len(alphabet))
	N = nbPossibilities(lenAlphabet, sizeMin, sizeMax)

	//Création de la table
	var tabRes [][2]uint64

	//Scan du reste des lignes du fichier
	for sc.Scan() {
		line := sc.Text()
		lineTemp := strings.Split(line, " : ")[1]

		table := strings.Split(lineTemp, " ")
		index, _ := strconv.Atoi(table[0])
		value, _ := strconv.Atoi(table[1])

		tabRes = append(tabRes, [2]uint64{uint64(index), uint64(value)})
	}
	return tabRes

}

//Q9
// affiche la table en entrée ainsi que
// les entêtes qui sont en variable globales
// 	- poolSize : donne le nombre d'éléments que la fonction affiche
// 	au début et à la fin du tableau
func afficheTable(table [][2]uint64, poolSize uint64) {

	//Imprimer les entêtes
	fmt.Printf("fonction de hachage : %s\nalphabet : %s\ntaille_min : %d\n taille_max : %d\nlargeur de la table : %d\nhauteur de la table : %d\n", hashMethod, alphabet, sizeMin, sizeMax, gWidth, gHeight)
	fmt.Printf("On affiche les %d premiers et le %d derniers : \n", poolSize, poolSize)

	firstValues := table[:poolSize]
	lastValues := table[len(table)-int(poolSize):]
	fmt.Printf("content: \n %v\n...\n %v", firstValues, lastValues)
}
