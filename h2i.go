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
	for i := uint64(0); i <= nbLetters; i++ {
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
	newIndex := index
	for i := uint64(1); i < width; i++ {
		newIndex = i2i(i, newIndex)
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
	// fmt.Printf("a : %d \n\n", test)
	a = uint64(test)

	//Notre recherche trouve l'index d'un élément == index, il faut maintenant
	// trouver le premier élément égal à index et le dernier
	if a < height {
		b = a
		for table[a][1] == index && a > 0 {
			a--
		}
		for table[b][1] == index && b < height-1 {
			b++
		}
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

	h2 := hash(clair)
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
	// ====
	fmt.Printf(">> Début de la recherche de l'empreinte %v dans la table :\n", empreinte)
	// ====

	var nbCandidats uint64 = 0
	for t := largeur - 1; t > 0; t-- {

		// fmt.Println("largeur et t", largeur, t, largeur-t, float64(largeur-t)/float64(largeur))
		// ==== affichage du loading
		fmt.Printf(">> Inversement en cours : %d%%\n", int((float64(largeur-t)/float64(largeur))*100))
		// ====

		idx := h2i(empreinte, t)
		for i := t + 1; i < largeur; i++ {
			idx = i2i(i, idx)
		}
		a, b := recherche(table, hauteur, idx)
		if !(a >= hauteur && b >= hauteur) {
			for i := a; i <= b; i++ {
				estObtenu, clair := verifieCandidat(empreinte, t, table[i][0])
				if estObtenu {
					// ====
					fmt.Printf(">> Un résultat a été obtenu : %s\n", clair)
					// ====
					return clair, nil
				} else {
					nbCandidats++
				}
			}
		}
	}
	return "", errors.New(">> L'inversion n'est pas possible.\n")

}

//Q12
func estimerCouverture(largeur uint64, hauteur uint64) (couverture float64) {
	m := float64(hauteur)
	nFloat := float64(N)
	v := 1.0
	for i := uint64(0); i < largeur; i++ {
		v = v * (1 - m/nFloat)
		m = nFloat * (1 - math.Exp(-m/nFloat))
	}
	couverture = 100 * (1 - v)
	return couverture
}
