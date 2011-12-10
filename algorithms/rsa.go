package main

import(
   "fmt"
   "os"
    "big"
   "strconv"
   "flag"
   "rand"
   "time"
   )


// Flags for encryption / decryption
var e = flag.Bool("e", false, "encrypt text")
var d = flag.Bool("d", false, "decrypt text")

var one = big.NewInt(1)
var two = big.NewInt(2)
var three = big.NewInt(3)


func main( ) {
   var buf []byte
   var plaintext []int
   var bCiphertext []big.Int

   flag.Parse()
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

   // public exponent
   bN := new( big.Int )

   // private exponent
   bD := new( big.Int )

   // If encrypt flag true
   if( *e ) {

      // Copies the temp buffer over to plaintext []int and sets up variables for
      // private and public keys.
      plaintext = []int(string(buf))
      bP       := new( big.Int )
      bQ       := new( big.Int )
      bE       := new( big.Int )
      bPMinus1 := new( big.Int )
      bQMinus1 := new( big.Int )
      bPhi     := new( big.Int )

      // Setting public exponenet e to three does not affect security
      bE = three

   // Reads the desired bitlength of the private keys and stores it in plength
   // int.
   temp := os.Args[3]
   plength, err := strconv.Atoi( temp ); if err != nil {
     fmt.Fprintf( os.Stderr, "%s\n", err )
   }

      // Generates private keys ensuring that they are not equal to each other
      bP = privateKeyGenerator( plength )
      bQ = privateKeyGenerator( plength )
      if bP == bQ {
         bQ = privateKeyGenerator( plength )
      }

      // Calculats public exponenet n and private exponent d
      bN.Mul( bP, bQ )
      bPMinus1.Sub( bP, one )
      bQMinus1.Sub( bQ , one )
      bPhi.Mul( bPMinus1, bQMinus1 )
      bD.ModInverse( bE, bPhi )

      fmt.Println( "private exponenets:" )
      fmt.Println( bP )
      fmt.Println( bQ )
      fmt.Println( bPhi )
      fmt.Println( bD )
      fmt.Println(  )

      // Encrypts text
      bCiphertext = encrypt( plaintext, bN, bE, len( plaintext ))

   // If decrypt flag true
   } else if( *d ) {
   fmt.Println( "in decrypt" )
   // Reads the secret key and stores it in bD
   temp := os.Args[3]
   bD.SetString( temp, 10 )
      // Copies letter in temporary buffer to a *big.Int[]
      bCiphertext = make( []big.Int, len( buf ) )
      for i := 0; i < len( buf ); i++ {
         bCiphertext[i].SetString( string( buf[i] ), 10)
         plaintext = decrypt( bCiphertext, bN, bD, len( bCiphertext ))
      }
   // If no flag is set, exits with error
   } else {
      fmt.Fprintf( os.Stderr, "RSA: No flag set\n" )
      os.Exit(0)
   }

   return
}

// Generates large (probably) prime number that is of the specified bit length
func privateKeyGenerator( plength int ) *big.Int {
   var tempB = new( big.Int )
   var mlength = new( big.Int )
   var prime = new( big.Int )
   isPrime := false
   plengthB := big.NewInt( int64( plength ) )

   // Calculates the length of the required random numbers as <  2^(plength)
   mlength.Exp( two, plengthB, nil )

   // Seends rand
   source := rand.NewSource( time.Nanoseconds() )
   r := rand.New( source )

   // Ensures that the number is (probably) prime
   for( !isPrime ) {
      tempB.Rand( r, mlength )
      // sets the 2^(plength) bit to ensure that the number it atleast that big
      prime.SetBit( tempB, plength, 1) 
      isPrime = big.ProbablyPrime( prime, 5 )
   }
   return prime
}


// Encrypts each letter in plaintext []int, using n and e as the public exponents. Returns a []big.Int with the ciphertext of each letter.
func encrypt( plaintext []int, n, e *big.Int, textLen int ) []big.Int {
   ciphertext := make( []big.Int, textLen )
   for i := 0; i < textLen; i++ {
      plainB := big.NewInt( int64( plaintext[i] ))
      ciphertext[i].Exp( plainB, e, n )
      fmt.Println( &ciphertext[i] )
   }
   return ciphertext
}

// Decrypts each letter in ciphertext []big.Int, using public exponent n and
// private exponenet d. Returns a []int with the decrypted text of each letter.
func decrypt( ciphertext []big.Int, n, d *big.Int, textLen int ) []int {
   bDecrypt := make( []big.Int, textLen )
   plaintext := make( []int, textLen )
   for i := 0; i < textLen; i++ {
      fmt.Println( i )
      bDecrypt[i].Exp( &ciphertext[i], d, n )
      plaintext[i] = int( bDecrypt[i].Int64() )
      fmt.Println( plaintext[i] )
   }
   return plaintext
}


