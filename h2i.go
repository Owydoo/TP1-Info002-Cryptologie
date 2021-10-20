package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func i2c(index uint64) string {
	var nbPossibilitiesMinSize = nbPossibilities(lenAlphabet, 1, sizeMin-1)
	var nbLetters uint64 = 0
	var indexTemp = index + nbPossibilitiesMinSize
	var powStock = lenAlphabet
	for indexTemp >= powStock {
		indexTemp -= powStock
		powStock *= lenAlphabet
		nbLetters++
	}

	var result = ""
	for i := 0; i <= int(nbLetters); i++ {
		result = result + string(alphabet[indexTemp%lenAlphabet])
		indexTemp /= lenAlphabet
	}
	return reverse(result)
}

//Q5
func h2i(hash []byte, t uint64) uint64 {
	first8o := hash[:8]
	data := binary.LittleEndian.Uint64(first8o)
	return (data + t) % N
}

//Q6
func i2i(t uint64, index uint64) uint64 {
	return h2i(hash(i2c(index)), t)
}

//Q6
func newString(index uint64, width uint64) uint64 {
	var newIndex uint64 = index
	for i := 1; i < int(width); i++ {
		newIndex = i2i(uint64(i), newIndex)
	}
	return newIndex
}

//Q8
func indexAleatoire(nb_max uint64) uint64 {
	// fmt.Printf("%d\n", rand.Int63n(int64(nb_max)))
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return uint64(r1.Intn(int(nb_max)))
}

//Q8
func creerTable(width uint64, height uint64) [][2]uint64 {
	// var texte_clair uint64 = index_aleatoire(N);

	var tabRes [][2]uint64

	for i := 0; i < int(height); i++ {
		// tabRes = append(tabRes, newString(indexAleatoire(N), width))
		var index = indexAleatoire(N)
		var newString = newString(index, width)
		var element [2]uint64 = [2]uint64{index, newString}

		tabRes = append(tabRes, element)
	}

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
		line := (sc.Text())
		sizeMinTemp, _ := strconv.Atoi(strings.Split(line, " : ")[1])
		sizeMin = uint64(sizeMinTemp)
	}
	if sc.Scan() { //taille maximale
		line := (sc.Text())
		sizeMaxTemp, _ := strconv.Atoi(strings.Split(line, " : ")[1])
		sizeMax = uint64(sizeMaxTemp)
	}
	if sc.Scan() { //largeur
		line := (sc.Text())
		widthTemp, _ := strconv.Atoi(strings.Split(line, " : ")[1])
		gWidth = uint64(widthTemp)
	}
	if sc.Scan() { //hauteur
		line := (sc.Text())
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
