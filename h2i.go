package main

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"math"
	"sort"
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

//Q10
// recherche dichotomique dans la table les premières et dernières lignes dont
// la seconde colonne est égale à idx
//   - table : table arc-en-ciel
//   - hauteur : nombre de chaines dans la table
//   - idx : indice à rechercher dans la dernière (deuxième) colonne
//   - a et b : (résultats) numéros des premières et dernières lignes dont les dernières colonnes sont égale à idx
// Si index n'est pas dans table, alors la fonction renvoie la longueur de table
func recherche(table [][2]uint64, height uint64, index uint64) (a uint64, b uint64) {
	 test := sort.Search(int(height), func(i int) bool {
		return table[i][1] == index
	})
	fmt.Printf("test : %d \n\n", test)
	 a= uint64(test)

	//Notre recherche trouve l'index d'un élément == index, il faut maintenant
	// trouver le premier élément égal à index et le dernier
	if a < uint64(len(table)) {
		b = a
		for table[a][1] == index && a>0{
			a--
		}
		for table[b][1] == index && b < height - 1{
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
		index = i2i(uint64(i), index)
	}
	clair = i2c(index)

	h2 := hash(string(clair))
	//fmt.Printf("empreinte : %v |  h2 : %v\n", empreinte, h2)
	return bytes.Equal(h2, empreinte), clair
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
	for t := largeur - 1; t > 0; t-- {
		idx := h2i(empreinte, t)
		for i := t + 1; i < largeur; i++ {
			idx = i2i(i, idx)
		}
		a, b := recherche(table, hauteur, idx)
		if !(a >= hauteur && b >= hauteur) {
			for i := a; i <= b; i++ {
				estObtenu, clair := verifieCandidat(empreinte, t, table[i][0])
				//fmt.Printf("est obtenu : %v clair : %s\n", estObtenu, clair)
				if estObtenu {
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
func estimerCouverture(largeur uint64, hauteur uint64) (couverture float64) {
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
