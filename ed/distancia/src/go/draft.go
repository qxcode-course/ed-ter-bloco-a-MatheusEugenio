package main

import "fmt"

func temProximo(str string, indexAtual int, numAlvo int, limite int) bool{

    if str[indexAtual] == byte(numAlvo){
        return true
    }

    if temProximo(str, indexAtual+1, numAlvo, limite) ||
       temProximo(str, indexAtual-1, numAlvo, limite) {

    }

    return false
}

func formatStr (str string, limite int){

    numAlvo := 0
    for i, r := range str{

        if r != '.'{
            continue
        }

        if !temProximo(str, i, numAlvo+1, limite){
            str = str[:i] + fmt.Sprintf("%d", numAlvo) + str[i+1:]
        }

    }
}


func main() {

    var str string
    fmt.Scan(&str)
    
    var limite int
    fmt.Scan(&limite)

    formatStr(str, limite)
    fmt.Println(str)
}