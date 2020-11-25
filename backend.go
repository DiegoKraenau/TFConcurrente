package main

import (
	"net/http"
	"log"
	"fmt"
	"net"
	"bufio"
	"encoding/json"
	"strconv"
	"strings"
	"io"
)



type entrada struct {
	Mathscore  int
}

type Ecuacion struct {
	M   float64  `json:"M"`
	B   float64  `json:"B"`
	MathScore   int  `json:"MathScore"`
	Forma string `json:"Forma"`
	Prediccion   float64  `json:"PrediccionDeWritingScore"`
}



var ecuacion Ecuacion
var remotehost string
var numero int
var variablesRecibidas []float64


func main(){
	

	handleResquest()

}

func handleResquest() {
	http.HandleFunc("/regresionLinealConJson", homePage)
	http.HandleFunc("/regresionLinealConQuery", query)
	http.HandleFunc("/prueba", prueba)
	//http.HandleFunc("/alumno", returnSingleAlumno)
	log.Fatal(http.ListenAndServe(":9000", nil))
	

}

func homePage(res http.ResponseWriter, r *http.Request) {
	res.Header().Set("Access-Control-Allow-Origin", "*")

    res.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	
	res.Header().Set("Content-Type", "application/json")

	
	
	var variable entrada
	decoder := json.NewDecoder(r.Body)
	
	err := decoder.Decode(&variable)

	if err != nil {
		panic(err)
	}
	
	defer r.Body.Close()



	
	numero=variable.Mathscore
	fmt.Print(variable.Mathscore)
	remotehost = fmt.Sprintf("localhost:%s", "9001") //Envia la informacion al primer nodo
	send(numero,res)
	
	
	numero=variable.Mathscore
	remotehost = fmt.Sprintf("localhost:%s", "9002") //Envia la informacion al segundo nodo
	send(numero,res)

	numero=variable.Mathscore
	remotehost = fmt.Sprintf("localhost:%s", "9003") //Envia la informacion al tercer nodo
	send(numero,res)


	ecuacion.B=variablesRecibidas[0]/variablesRecibidas[1]
	ecuacion.M=variablesRecibidas[2]/variablesRecibidas[1]
	ecuacion.MathScore=numero
	bforma := fmt.Sprintf("%f", ecuacion.B) 
	mforma := fmt.Sprintf("%f", ecuacion.M) 
	ecuacion.Forma=bforma+" X + "+mforma
	ecuacion.Prediccion=ecuacion.B*float64(numero)+ecuacion.M
	jsonBytes, _ := json.MarshalIndent(ecuacion, "", " ")
	io.WriteString(res, string(jsonBytes))
	

}

func query(res http.ResponseWriter, r *http.Request) {
	res.Header().Set("Access-Control-Allow-Origin", "*")

    res.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	
	res.Header().Set("Content-Type", "application/json")

	
	

	mathScore := r.FormValue("mathScore")
	log.Println(mathScore)

	numero, err := strconv.Atoi(mathScore)
	if err == nil {
		fmt.Println(numero)
	}
	

	remotehost = fmt.Sprintf("localhost:%s", "9001") //Envia la informacion al primer nodo
	send(numero,res)
	
	
	numero, err = strconv.Atoi(mathScore)
	if err == nil {
		fmt.Println(numero)
	}
	
	remotehost = fmt.Sprintf("localhost:%s", "9002") //Envia la informacion al segundo nodo
	send(numero,res)

	numero, err = strconv.Atoi(mathScore)
	if err == nil {
		fmt.Println(numero)
	}
	
	remotehost = fmt.Sprintf("localhost:%s", "9003") //Envia la informacion al tercer nodo
	send(numero,res)


	ecuacion.B=variablesRecibidas[0]/variablesRecibidas[1]
	ecuacion.M=variablesRecibidas[2]/variablesRecibidas[1]
	ecuacion.MathScore=numero
	bforma := fmt.Sprintf("%f", ecuacion.B) 
	mforma := fmt.Sprintf("%f", ecuacion.M) 
	ecuacion.Forma=bforma+" X + "+mforma
	ecuacion.Prediccion=ecuacion.B*float64(numero)+ecuacion.M
	jsonBytes, _ := json.MarshalIndent(ecuacion, "", " ")
	io.WriteString(res, string(jsonBytes))
	

}

func prueba(res http.ResponseWriter, r *http.Request) {
	res.Header().Set("Access-Control-Allow-Origin", "*")

    res.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	fmt.Print("HOLAAAAA")
	
}

func send(num int,res http.ResponseWriter) {
    conn, _ := net.Dial("tcp", remotehost)
    defer conn.Close()
    fmt.Fprintf(conn, "%d\n", num)
	


	hostname := fmt.Sprintf("localhost:%s", "9004")
	ln, _ := net.Listen("tcp", hostname)
    defer ln.Close()
    conn, _ = ln.Accept()
    handle(conn)

	


}
	

func handle(conn net.Conn) {
    defer conn.Close()
    r := bufio.NewReader(conn)
    str, _ := r.ReadString('\n')
	num := strings.TrimSpace(str)
	if s, err := strconv.ParseFloat(num, 64); err == nil {
		fmt.Println(s) // 3.14159265
		variablesRecibidas=append(variablesRecibidas,s)
		//ecuacion.B=s
		//jsonBytes, _ := json.MarshalIndent(ecuacion, "", " ")
		//io.WriteString(res, string(jsonBytes))
	}

	

}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

