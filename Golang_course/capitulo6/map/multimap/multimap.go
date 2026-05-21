// -*- coding: utf-8 -*-
// multimap.go
// -----------------------------------------------------------------------------
//
// Started on <lun 23-02-2026 21:37:51.297428825 (1771879071)>
// Carlos Linares López <carlos.linares@uc3m.es>
//

// Implementación de un multimapa de cadenas a enteros
package multimap

import "errors"

// devuelve el índice de la primera ocurrencia de item en slice. Si no
// existe, devuelve -1
func index(item int, slice []int) int {

	// para todos los elementos del slice
	for k, v := range slice {

		// si este es el que se busca, se devuelve su posición
		if v == item {
			return k
		}
	}

	// en este punto, el item no existe demostradamente en slice
	return -1
}

// elimina el elemento en el índice indicado. Devuelve un error si el índice
// no existe; en otro caso devuelve el slice resultante
func remove(index int, slice []int) ([]int, error) {

	if index < 0 || index >= len(slice) {
		return slice, errors.New("Removing out of bounds")
	}

	// elimina item en O(1) intercambiándolo con la última posición, y
	// eliminando después el último elemento
	slice[index] = slice[len(slice)-1]
	return slice[:len(slice)-1], nil
}

// inserta value en el contenedor indexado por key
func Set(key string, value int, multimap map[string][]int) {

	// Si la clave no existe
	if _, ok := multimap[key]; !ok {

		// inicializa un contenedor con el valor indicado
		multimap[key] = []int{value}
	} else {

		// en otro caso, añade el nuevo valor
		multimap[key] = append(multimap[key], value)
	}
}

// devuelve todos los valores indexados por key y ok a true; si no hay valores
// indexados por key devuelve un slice nulo y ok=false
func Get(key string, multimap map[string][]int) ([]int, bool) {

	// si hay valores indexados por la clave
	if values, ok := multimap[key]; !ok {

		// si no, devuelve un slice nulo con false
		return []int{}, false
	} else {

		// en otro caso, devuelve los valores encontrados y true
		return values, true
	}
}

// elimina value del contenedor indexado por key. Devuelve true si
// efectivamente ha sido eliminado, y falso en otro caso
func Remove(key string, value int, multimap map[string][]int) bool {

	if values, ok := multimap[key]; ok {

		idx := index(value, values)
		found := idx >= 0
		if found {
			multimap[key], _ = remove(idx, multimap[key])
			if len(multimap[key]) == 0 {
				delete(multimap, key)
			}
			return true
		}
	}

	return false
}

// Local Variables:
// mode:go
// fill-column:80
// End:
