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
   "coolFunctions"
   )

// Flags indicating encryption / decryption
var e = flag.Bool("e", false, "encrypt text")
var d = flag.Bool("d", false, "decrypt text")

// big.Int representation for commonly used numbers
var zero    = big.NewInt(0)
var one     = big.NewInt(1)
var two     = big.NewInt(2)
var three   = big.NewInt(3)

// Naming convention for big.Int: Prefix b

func main( ) {
   var plaintext []int
   var bCiphertext []*big.Int

   // Flags indicate wether program should encrypt or decrypt
   flag.Parse()
      // public exponent n
      bN := new( big.Int )

      // private exponent d
      bD := new( big.Int )

      // If encrypt flag true
      if( *e ) {
         // Input should be: flag plaintext ciphertext bitlength
         if len(os.Args) != 4 {
            fmt.Fprintf( os.Stderr, "RSA: Not enough arguments \n" )
         } else if st, err := os.Stat( os.Args[2] ); err != nil {
            fmt.Fprintf( os.Stderr, "RSA: %s\n", err );
         } else if f, err := os.Open( os.Args[2] ); err != nil {
            fmt.Fprintf( os.Stderr, "RSA: %s\n", err );
         } else {
            // Reads the plaintext
            intBuf := make( []byte, st.Size )
            f.Read( intBuf )
            plaintext = []int( string( intBuf ))

            // Reads string bitlength ( of the private keys ) and stores it as
            // an int
            temp := os.Args[3]
            plength, err := strconv.Atoi( temp ); if err != nil {
               fmt.Fprintf( os.Stderr, "%s\n", err )
            }

            // Sets up variables for calculating keys
            bP       := new( big.Int )
            bQ       := new( big.Int )
            bE       := new( big.Int )
            bPMinus1 := new( big.Int )
            bQMinus1 := new( big.Int )
            bPhi     := new( big.Int )
            bRem     := new( big.Int )

            // Generates private keys p != q
            bP = privateKeyGenerator( plength )
            bQ = privateKeyGenerator( plength )

            fmt.Println( "p", bP )
            fmt.Println( "q", bQ )
            for bP.Cmp(bQ) == 0 {
               fmt.Println( "ji" )
               bQ = privateKeyGenerator( plength )
            }

            // Calculates public key n and private key d
            // n = pq
            bN.Mul( bP, bQ )
            bPMinus1.Sub( bP, one )
            bQMinus1.Sub( bQ , one )

            // Phi = (p - 1)(q - 1)
            bPhi.Mul( bPMinus1, bQMinus1 )

            // Generares random e < phi
            source := rand.NewSource( time.Nanoseconds() )
            r := rand.New( source )
            bE.Rand( r, bPhi )
            for( bRem.Rem( bE, bPhi ) == zero ) {
               bE.Rand( r, bPhi )
            }

            // d = e^{-1} (  phi )
            bD = coolFunctions.ModuloInverse( bE, bPhi )

            fmt.Println( "Public keys:" )
            fmt.Println( "n", bN )
            fmt.Println( "e", bE )

            fmt.Println( "Private keys:" )
            fmt.Println( "d", bD, "\n" )

            // Encrypts text
            bCiphertext = encrypt( plaintext, bN, bE, len( plaintext ))
         }

      // If decrypt flag true
      } else if( *d ) {
         // Input should be: flag cichertext decrypted-text n d
         if len(os.Args) != 5 {
            fmt.Fprintf( os.Stderr, "RSA: Not enough arguments \n" )
         } else if st, err := os.Stat( os.Args[2] ); err != nil {
            fmt.Fprintf( os.Stderr, "RSA: %s\n", err );
         } else if f, err := os.Open( os.Args[2] ); err != nil {
            fmt.Fprintf( os.Stderr, "RSA: %s\n", err );
         } else {
            stringBuf := make( []string, st.Size )

            // Ciphertexts are large. We need a buffered reader.
            r := bufio.NewReader( f )

            // Ciphertexts are seperated by new lines in the file. The next block
            // of code stores them into a temporary string slice.
            for {
               if temp, err := r.ReadString( '\n' ); err != nil {
                  break
               } else {
                  stringBuf = append( stringBuf, temp )
               }
            }

            // Now we create a slice of bigInts and copy the ciphertext over.
            bCiphertext = make( []*big.Int, len( stringBuf ) )
            for i := 0; i < len( stringBuf ); i++ {
               bCiphertext[i] = new( big.Int )
               bCiphertext[i].SetString( string( stringBuf[i] ), 10)
            }

            // Reads secret key d
            temp_d := os.Args[3]
            bD.SetString( temp_d, 10 )

            // Reads public key n
            temp_n := os.Args[4]
            bN.SetString( temp_n, 10 )

            // Decrypt!!
            plaintext = decrypt( bCiphertext, bN, bD, len( bCiphertext ))
         }
      // If no flag ( encrypt or decrypt ) is set, exits with error
      } else {
         fmt.Fprintf( os.Stderr, "RSA: No flag set\n" )
         os.Exit(0)
      }

   return
}

// Generates large (probably) prime number of the specified bit length
func privateKeyGenerator( plength int ) *big.Int {
   var tempB = new( big.Int )
   var mlength = new( big.Int )
   var prime = new( big.Int )
   isPrime := false

   // Exp function takes int64s as inputs for power
   plengthB := big.NewInt( int64( plength ) )

   // Calculates the upper bound of the required random numbers as <  2^(plength)
   // We actually need numbers of length 2^n, but we will take care of that
   // later
   mlength.Exp( two, plengthB, nil )

   // Seeds rand
   source := rand.NewSource( time.Nanoseconds() )
   r := rand.New( source )

   // Ensures that the number is (probably) prime
   for( !isPrime ) {
      tempB.Rand( r, mlength )
      // sets the 2^(plength) bit to ensure that our prime has 2^n bits
      prime.SetBit( tempB, plength, 1)
      // Miller-Rabin probibalistic primality test
      isPrime = big.ProbablyPrime( prime, 5 )
   }
   return prime
}


// Encrypts each letter in plaintext []int, using n and e as the public exponents. Returns a []big.Int with the ciphertext of each letter.
func encrypt( plaintext []int, n, e *big.Int, textLen int ) []*big.Int {
   ciphertext := make( []*big.Int, textLen )
   // Encrypt each character seperately
   for i := 0; i < textLen; i++ {
      // Caste each plaintext character to int64 and then big.Int
      bPlain := big.NewInt( int64( plaintext[i] ))
      // Initialize each element in ciphertext as a big.Int and set it to: c = m^e (mod n)
      ciphertext[i] = new( big.Int )
      ciphertext[i].Exp( bPlain, e, n )
   }
   return ciphertext
}

// Decrypts each letter in ciphertext []big.Int, using public exponent n and
// private exponenet d. Returns a []int with the decrypted text of each letter.
func decrypt( ciphertext []*big.Int, n, d *big.Int, textLen int ) []int {
   bDecrypt := make( []*big.Int, textLen )
   plaintext := make( []int, textLen )
   // Decrypt each character seperately
   for i := 0; i < textLen; i++ {
      // Initialize each element in decrypt as a big.Int and set it to: 
      // m = c^d (mod n)
      bDecrypt[i] = new( big.Int )
      bDecrypt[i].Exp( ciphertext[i], d, n )
      // Caste each element in decrypt to a ( int64 and then an ) int ( as the message originally
      // was ).
      plaintext[i] = int( bDecrypt[i].Int64() )
   }
   return plaintext
}

