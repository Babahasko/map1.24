package main

import (
    "fmt"
)

func main() {
    // Исходный хэш
    hash := uint64(0xABCDEF1234) // Пример хэша
	fmt.Printf("Hash (binary): %b\n", hash)
	fmt.Printf("Hash (hex): %X\n", hash)

    // Получение H2 (первые 7 бит)
    h2 := hash & 0x7F // Маска для первых 7 бит
    fmt.Printf("H2 (binary): %07b\n", h2) // Вывод в двоичной системе
    fmt.Printf("H2 (hex): %02X\n", h2)  // Вывод в шестнадцатеричной системе

    // Получение H1 (оставшиеся биты)
    h1 := hash >> 7 // Сдвигаем вправо на 7 бит
	fmt.Printf("H1 (binary): %b\n", h1)
    fmt.Printf("H1 (hex): %X\n", h1)    // Вывод в шестнадцатеричной системе
}