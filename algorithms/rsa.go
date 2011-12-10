package main

import(
   "fmt"
   "os"
 //  "big"
   "strconv"
   )

func main( ) {
   var buf []byte
   // Read in command line input, open the file given as the first argument
   // to the program and save it in byte[] buf.
   if len(os.Args) < 3 {
      fmt.Fprintf( os.Stderr, "RSA: Not enough arguments \n" )
   } else if st, err := os.Stat( os.Args[2] ); err != nil {
      fmt.Fprintf( os.Stderr, "RSA: %s\n", err );
   } else if f, err := os.Open( os.Args[2] ); err != nil {
      fmt.Fprintf( os.Stderr, "RSA: %s\n", err );
   } else {
      buf = make( []byte, st.Size )
      f.Read( buf )
   }

   // Reads the desired bitlength of the private keys and stores it in pleangth
   // int.
    temp := os.Args[3]
    plength, err := strconv.Atoi(temp); if err != nil {
      fmt.Fprintf( os.Stderr, "%s\n", err )
    }

    fmt.Printf("plength: %d\n", plength)
    fmt.Println( string( buf ))

    return
}

/*// Generates large (probably) prime number that is of the specified bit length
func privateKeyGenerator( plength int ) *big.Int {
   return
}

// Encrypts each letter in plaintext []int, using n and e as the public exponents. Returns a []big.Int with the ciphertext of each letter.
func encrypt( plaintext []int, n, e *big.Int, textLen int ) []big.Int {
   return
}

// Decrypts each letter in ciphertext []big.Int, using public exponent n and
// private exponenet d. Returns a []int with the decrypted text of each letter.
func decrypt( ciphertext []*big.Int, n, d *big.Int, textLen int ) []int {
   return
}
*/


