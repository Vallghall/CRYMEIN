package alphabet

var russianIndexes = map[rune]int64{
	'А': 1, 'Б': 2, 'В': 3, 'Г': 4, 'Д': 5,
	'Е': 6, 'Ё': 7, 'Ж': 8, 'З': 9, 'И': 10,
	'Й': 11, 'К': 12, 'Л': 13, 'М': 14, 'Н': 15,
	'О': 16, 'П': 17, 'Р': 18, 'С': 19, 'Т': 20,
	'У': 21, 'Ф': 22, 'Х': 23, 'Ц': 24, 'Ч': 25,
	'Ш': 26, 'Щ': 27, 'Ъ': 28, 'Ы': 29, 'Ь': 30,
	'Э': 31, 'Ю': 32, 'Я': 33,
}

var russianRunes map[int64]rune

func init() {
	russianRunes = make(map[int64]rune, len(russianIndexes))
	for key, value := range russianIndexes {
		russianRunes[value] = key
	}
	russianRunes[0] = 'Я'
}

func ToRussianIndexes(txt string) (res []int64) {
	for _, letter := range txt {
		res = append(res, russianIndexes[letter])
	}
	return res
}

func ToRussianRunes(txt []int64) []rune {
	res := make([]rune, len(txt))
	for i, num := range txt {
		res[i] = russianRunes[num%33]
	}
	return res
}
