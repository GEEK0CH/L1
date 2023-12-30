package main

import "fmt"

func sliseCut(slice []string, numb int) (sliceN []string) { //немного некорректный вариант исполнения, т.к. теряется порядок
	slice[numb-1] = slice[len(slice)-1] //копировать последний элемент в индекс numb
	slice[len(slice)-1] = ""            //удалить последний элемент
	sliceN = slice[:len(slice)-1]       //урезать срез.
	return
}

func sliseCut2(slice []string, numb int) (sliceN []string) { //в данном случае порядок сохраняется
	for i := 0; i <= len(slice)-1; i++ { // цикл для перебора элементов слайся
		if i != numb { //исключение для записи
			sliceN = append(sliceN, slice[i])
		}
	}
	return
}

func sliseCut3(slice []string, numb int) []string { //в данном случае порядок сохраняется
	return append(slice[:numb], slice[numb+1:]...) //вырезается все что находит до и после элемента которого надо удалить
}

func main() {
	str := []string{"snow", "dog", "sun", "moon", "ball"}
	fmt.Println(sliseCut(str, 2))
	str = []string{"snow", "dog", "sun", "moon", "ball"}
	fmt.Println(sliseCut2(str, 2))
	str = []string{"snow", "dog", "sun", "moon", "ball"}
	fmt.Println(sliseCut3(str, 2))

}
