package main

import(
   "fmt"
   "os"
    "big"
   "strconv"
   "flag"
   "rand"
   "time"
   "bufio"
   )


// Flags for encryption / decryption
var e = flag.Bool("e", false, "encrypt text")
var d = flag.Bool("d", false, "decrypt text")

var zero = big.NewInt(0)
var one = big.NewInt(1)
var two = big.NewInt(2)
var three = big.NewInt(3)


func main( ) {
   var plaintext []int
   var bCiphertext []*big.Int

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
      // public exponent
      bN := new( big.Int )

      // private exponent
      bD := new( big.Int )

      // If encrypt flag true
      if( *e ) {

         // Read the plaintext into a temp buffer
         intBuf := make( []byte, st.Size )
         f.Read( intBuf )

         // Reads the desired bitlength of the private keys and stores it in plength
         // int.
         temp := os.Args[3]
         plength, err := strconv.Atoi( temp ); if err != nil {
            fmt.Fprintf( os.Stderr, "%s\n", err )
         }

         // Copies the temp buffer over to plaintext []int and sets up variables for
         // private and public keys.
         plaintext = []int( string( intBuf ))
         bP       := new( big.Int )
         bQ       := new( big.Int )
         bE       := new( big.Int )
         bPMinus1 := new( big.Int )
         bQMinus1 := new( big.Int )
         bPhi     := new( big.Int )
         rem      := new( big.Int )


         // Generates private keys ensuring that they are not equal to each other
         bP = privateKeyGenerator( plength )
         bQ = privateKeyGenerator( plength )

         fmt.Println( "p", bP )
         fmt.Println( "q", bQ )
         for bP.Cmp(bQ) == 0 {
            fmt.Println( "ji" )
            bQ = privateKeyGenerator( plength )
         }

         // Calculates public exponent n and private exponent d

         bN.Mul( bP, bQ )
         bPMinus1.Sub( bP, one )
         bQMinus1.Sub( bQ , one )
         bPhi.Mul( bPMinus1, bQMinus1 )

         source := rand.NewSource( time.Nanoseconds() )
         r := rand.New( source )
         bE.Rand( r, bPhi )
         for( rem.Rem( bE, bPhi ) == zero ) {
            bE.Rand( r, bPhi )
         }

         bD.ModInverse( bE, bPhi )

         fmt.Println( "private exponents:" )
         fmt.Println( "p", bP )
         fmt.Println( "q", bQ )
         fmt.Println( "phi", bPhi )
         fmt.Println( "n", bN )
         fmt.Println( "e", bE )
         fmt.Println( "d", bD, "\n" )

         // Encrypts text
         bCiphertext = encrypt( plaintext, bN, bE, len( plaintext ))

      // If decrypt flag true
      } else if( *d ) {
         var stringBuf []string
         r := bufio.NewReader( f )
         for {
            if temp, err := r.ReadString( '\n' ); err != nil {
               break
            } else {
               stringBuf = append( stringBuf, temp )
            }
         }

         // Reads the secret key and stores it in bD
         temp_d := os.Args[3]
         bD.SetString( temp_d, 10 )

         // Reads the secret key and stores it in bD
         temp_n := os.Args[4]
         bN.SetString( temp_n, 10 )

         // Copies letter in temporary buffer to a *big.Int[]
         bCiphertext = make( []*big.Int, len( stringBuf ) )

         for i := 0; i < len( stringBuf ); i++ {
            bCiphertext[i] = new( big.Int )
            bCiphertext[i].SetString( string( stringBuf[i] ), 10)
         }

         fmt.Println( "n", bN, "d", bD )

         plaintext = decrypt( bCiphertext, bN, bD, len( bCiphertext ))
         // If no flag is set, exits with error
      } else {
         fmt.Fprintf( os.Stderr, "RSA: No flag set\n" )
         os.Exit(0)
      }
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
func encrypt( plaintext []int, n, e *big.Int, textLen int ) []*big.Int {
   ciphertext := make( []*big.Int, textLen )
   for i := 0; i < textLen; i++ {
      plainB := big.NewInt( int64( plaintext[i] ))
      ciphertext[i] = new( big.Int )
      ciphertext[i].Exp( plainB, e, n )
      fmt.Println( ciphertext[i] )
   }
   return ciphertext
}

// Decrypts each letter in ciphertext []big.Int, using public exponent n and
// private exponenet d. Returns a []int with the decrypted text of each letter.
func decrypt( ciphertext []*big.Int, n, d *big.Int, textLen int ) []int {
   bDecrypt := make( []*big.Int, textLen )
   plain := make( []int, textLen )
   for i := 0; i < textLen; i++ {
      bDecrypt[i] = new( big.Int )
      bDecrypt[i].Exp( ciphertext[i], d, n )
      plain[i] = int( bDecrypt[i].Int64() )
      fmt.Println( string( plain[i] ) )
   }
   return plain
}
