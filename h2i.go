package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"math"
	"math/rand"
	"os"
	"sort"
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
func sortTableByLastColumn(table [][2]uint64) {
	sort.Slice(table, func(i, j int) bool {
		return table[i][1] < table[j][1]
	})

}

//Q8
func creerTable(width uint64, height uint64) [][2]uint64 {

	var tabRes [][2]uint64

	for i := 0; i < int(height); i++ {
		var index = indexAleatoire(N)
		var newString = newString(index, width)
		var element [2]uint64 = [2]uint64{index, newString}

		tabRes = append(tabRes, element)
	}
	sortTableByLastColumn(tabRes)

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

//Q10
// recherche dichotomique dans la table les premières et dernières lignes dont
// la seconde colonne est égale à idx
//   - table : table arc-en-ciel
//   - hauteur : nombre de chaines dans la table
//   - idx : indice à rechercher dans la dernière (deuxième) colonne
//   - a et b : (résultats) numéros des premières et dernières lignes dont les dernières colonnes sont égale à idx
// Si index n'est pas dans table, alors la fonction renvoie la longueur de table
func recherche(table [][2]uint64, height uint64, index uint64) (a uint64, b uint64) {
	//recherche dichotomique
	debut := 0
	fin := len(table) - 1

	for debut <= fin {
		mediane := (debut + fin) / 2

		if table[mediane][1] < index {
			debut = mediane + 1
		} else {
			fin = mediane - 1
		}
	}

	if debut == len(table) || table[debut][1] != index {
		return gHeight, gHeight
		// Renvoie la longueur du tableau si index n'est pas dans le tableau.
	} else {
		a = uint64(debut)
	}

	//Notre recherche trouve l'index d'un élément == index, il faut maintenant
	// trouver le premier élément égal à index et le dernier
	// fmt.Printf("truc %d %d\n", a, table[a][1])

	if a < uint64(len(table)) {
		b = a
		for table[a][1] == index {
			a--
		}
		for table[b][1] == index {
			b++
		}
		fmt.Println("a et b", a, b)
		return a, b
	}

	return gHeight, gHeight
}

// Q10
// vérifie si un candidat est correct
//   - empreinte : empreinte à inverser
//   - t : numéro de la colonne où a été trouvé le candidat
//   - index : indice candidat (de la colonne t)
//   - clair : résultat : contient le texte clair obtenu
func verifieCandidat(empreinte []byte, t uint64, index uint64) (estObtenu bool, clair string) {
	for i := 1; i < int(t); i++ {
		index = i2i(index, index)
	}
	// clairTemp, _ := strconv.Atoi()
	clair = i2c(index)

	h2 := hash(string(clair))
	// h2_data := binary.LittleEndian.Uint64(h2)
	return bytes.Equal(h2, empreinte), clair
	// h2 == empreinte, clair

	// first8o := hash[:8]
	// data := binary.LittleEndian.Uint64(first8o)

}

//Q10
// essaie d'inverser l'empreinte h
//   - table : table arc-en-ciel
//   - hauteur : nombre de chaines dans la table
//   - largeur : longueur des chaines
//   - empreinte : empreinte à inverser
//   - clair : (résultat) texte clair dont l'empreinte est h
func inverse(table [][2]uint64, hauteur uint64, largeur uint64, empreinte []byte) (clair string, err error) {
	var nb_candidats uint64 = 0
	// byte_empreinte := make([]byte, 16)
	// binary.LittleEndian.PutUint64(byte_empreinte, empreinte)
	for t := largeur - 1; t > 0; t-- {
		idx := h2i(empreinte, t)
		for i := t + 1; i < largeur; i++ {
			idx = i2i(idx, i)
		}
		a, b := recherche(table, hauteur, idx)
		if !(a >= hauteur && b >= hauteur) {
			for i := a; i <= b; i++ {
				estCandidat, clair := verifieCandidat(empreinte, t, table[i][0])
				if estCandidat {
					return clair, nil
				} else {
					nb_candidats++
				}
			}
		}
	}
	return "", errors.New("pas trouvé de candidats")

}

//Q12
func estimerCouverture(table [][2]uint64, largeur uint64, hauteur uint64) (couverture float64) {
	m := float64(hauteur)
	N := float64(nbPossibilities(lenAlphabet, sizeMin, sizeMax))
	v := 1.0
	for i := 0; i < int(largeur); i++ {
		v = v * (1 - m/N)
		m = N * (1 - math.Exp(-m/N))
	}
	couverture = 100 * (1 - v)
	return couverture
}
