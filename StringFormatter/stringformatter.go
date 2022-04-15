
package StringFormatter

import (
    "fmt"
)

func ShowTransactionId(trnxID string) string {
	if trnxID != "" {
		return trnxID[0:10] //RETURN 1ST 10 CHARACTER
	}

	return trnxID
}


func ScientificNotation() {

    val := 1273.78888769000

    fmt.Printf("%f\n", val)
    fmt.Printf("%e\n", val)
    fmt.Printf("%g\n", val)
    fmt.Printf("%E\n", val)
    fmt.Printf("%G\n", val)

    fmt.Println("-------------------------")

    fmt.Printf("%.10f\n", val)
    fmt.Printf("%.10e\n", val)
    fmt.Printf("%.10g\n", val)
    fmt.Printf("%.10E\n", val)
    fmt.Printf("%.10G\n", val)

    fmt.Println("-------------------------")

    val2 := 66_000_000_000.1200

    fmt.Printf("%f\n", val2)
    fmt.Printf("%e\n", val2)
    fmt.Printf("%g\n", val2)
    fmt.Printf("%E\n", val2)
    fmt.Printf("%G\n", val2)
}
