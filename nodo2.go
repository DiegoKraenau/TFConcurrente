package main

import (
    "bufio"
    "fmt"
    "net"
    "strings"
    "strconv"
    "flag"
    "os"
    "encoding/csv"
    "io"
    "math"
    "net/http"
)


type LinearRegression struct {
	x, y []float64
	xTotal, yTotal, xSqTotal, xYTotal, numerator, denominator,
	yInt float64
}



func leerCSV(c chan []float64) {

	fileUrl := "https://github.com/DiegoKraenau/TA2Concurrente/raw/main/StudentsPerformance.csv"
	err := DescargarArchivo("StudentsPerformance.csv", fileUrl)
	if err != nil {
		panic(err)
	}

	path := flag.String("f", "./StudentsPerformance.csv", "CSV filepath")
	dataFile, err := os.Open(*path)

	reader := csv.NewReader(dataFile)

	var x []float64
	var y []float64

	for {
		record, err := reader.Read()

		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Printf("Error al leer el csv: %v", err)

		}

		// convertir todas las filas a float64
		for i, val := range record {
			float, err := strconv.ParseFloat(val, 64)

			if err != nil {
				fmt.Printf("Error al convertir el valor <%v> afloat64: %v", val, err)
				break
			}

			if i < len(record)-1 {
				x = append(x, float)
			} else {
				y = append(y, float)
			}
		}

	}
	c <- x
	c <- y

}

func (l *LinearRegression) ReturnEcuacionParts(c chan []float64,num int) {


	l.LinearRegressionInit(c)



}

func (l *LinearRegression) LinearRegressionInit(c chan []float64) {

	l.x = <-c
	l.y = <-c
	//fmt.Println(len(l.X))
	for i, _ := range l.x {
		l.xTotal = l.xTotal + l.x[i]
		l.yTotal = l.yTotal + l.y[i]
		l.xSqTotal = l.xSqTotal + math.Pow(l.x[i], 2)
		l.xYTotal = l.xYTotal + (l.x[i] * l.y[i])
	}


	l.denominator = (float64(len(l.x)) * l.xSqTotal) - (math.Pow(l.xTotal, 2))

}



func DescargarArchivo(filepath string, url string) error {

	// Obtiene la data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Crea el archivo
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Lo escribe al archivo
	_, err = io.Copy(out, resp.Body)
	return err
}


var remotehost string

func main() {

    hostname := fmt.Sprintf("localhost:%s", "9002")


    remotehost = fmt.Sprintf("localhost:%s", "9005")

    // Listener!
    ln, _ := net.Listen("tcp", hostname)
    defer ln.Close()
    for {
        conn, _ := ln.Accept()
        go handle(conn)
    }
}

func handle(conn net.Conn) {
    defer conn.Close()
    r := bufio.NewReader(conn)
    str, _ := r.ReadString('\n')
    num, _ := strconv.Atoi(strings.TrimSpace(str))
    //Aqui va la logica
    ch := make(chan []float64, 2)

    leerCSV(ch)
    
    var partEcuacion LinearRegression
    partEcuacion.ReturnEcuacionParts(ch,num)
    send(partEcuacion.denominator)//El segundo nodo retorna la segunda parte de la ecuacion general

}

func send(num float64) {
    conn, _ := net.Dial("tcp", remotehost)
    defer conn.Close()
    fmt.Print(num)
    fmt.Fprintf(conn, "%7.5f",num)
}


