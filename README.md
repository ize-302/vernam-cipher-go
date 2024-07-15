# Go implentation of Vernam Cipher

### About Vernam cypher

The Vernam cipher, also known as the one-time pad, is a symmetric key encryption technique that was invented by Gilbert Vernam in 1917. It's known for being theoretically unbreakable when used correctly.

## How it works

### Key generation

A key is generated that is as long as the message to be encrypted. This key must be completely random

### Encryption

- Convert the message: The message is converted to a binary representation i.e using ASCII code for each character.
- Convert the key: The key is converted into a binary representation i.e using ASCII code for each character.
- Produce the ciphertext: Each bit of the message is XORed with the corresponding bit of the key to produce the ciphertext

> Mathematically, this can be represented as: Ci = Pi ⊕ Ki
> where Ci is the ith bit of ciphertext, Pi is the ith bit of the message (plain text), and Ki is the ith bit of the key

##### ASCII TABLE

```
Dec  |  Char  |  Dec  |  Char  |  Dec  |  Char  |  Dec  |  Char
-----+--------+-------+--------+-------+--------+-------+--------
  0  |  NUL   |   32  |  Space |   64  |   @    |   96  |   `
  1  |  SOH   |   33  |   !    |   65  |   A    |   97  |   a
  2  |  STX   |   34  |   "    |   66  |   B    |   98  |   b
  3  |  ETX   |   35  |   #    |   67  |   C    |   99  |   c
  4  |  EOT   |   36  |   $    |   68  |   D    |  100  |   d
  5  |  ENQ   |   37  |   %    |   69  |   E    |  101  |   e
  6  |  ACK   |   38  |   &    |   70  |   F    |  102  |   f
  7  |  BEL   |   39  |   '    |   71  |   G    |  103  |   g
  8  |  BS    |   40  |   (    |   72  |   H    |  104  |   h
  9  |  TAB   |   41  |   )    |   73  |   I    |  105  |   i
 10  |  LF    |   42  |   *    |   74  |   J    |  106  |   j
 11  |  VT    |   43  |   +    |   75  |   K    |  107  |   k
 12  |  FF    |   44  |   ,    |   76  |   L    |  108  |   l
 13  |  CR    |   45  |   -    |   77  |   M    |  109  |   m
 14  |  SO    |   46  |   .    |   78  |   N    |  110  |   n
 15  |  SI    |   47  |   /    |   79  |   O    |  111  |   o
 16  |  DLE   |   48  |   0    |   80  |   P    |  112  |   p
 17  |  DC1   |   49  |   1    |   81  |   Q    |  113  |   q
 18  |  DC2   |   50  |   2    |   82  |   R    |  114  |   r
 19  |  DC3   |   51  |   3    |   83  |   S    |  115  |   s
 20  |  DC4   |   52  |   4    |   84  |   T    |  116  |   t
 21  |  NAK   |   53  |   5    |   85  |   U    |  117  |   u
 22  |  SYN   |   54  |   6    |   86  |   V    |  118  |   v
 23  |  ETB   |   55  |   7    |   87  |   W    |  119  |   w
 24  |  CAN   |   56  |   8    |   88  |   X    |  120  |   x
 25  |  EM    |   57  |   9    |   89  |   Y    |  121  |   y
 26  |  SUB   |   58  |   :    |   90  |   Z    |  122  |   z
 27  |  ESC   |   59  |   ;    |   91  |   [    |  123  |   {
 28  |  FS    |   60  |   <    |   92  |   \    |  124  |   |
 29  |  GS    |   61  |   =    |   93  |   ]    |  125  |   }
 30  |  RS    |   62  |   >    |   94  |   ^    |  126  |   ~
 31  |  US    |   63  |   ?    |   95  |   _    |  127  |  DEL
```

### Decryption

- The ciphertext is converted to a binary representation
- The same key is used to decrypt the ciphertext. Each bit of the ciphertext is XORed with the corresponding bit of the key to retrieve the original plaintext.

> Mathematically, this can be represented as: Pi = Ci ⊕ Ki
> where Pi is the ith bit of the message (plain text),Ci is the ith bit of ciphertext, and Ki is the ith bit of the key

### Key Properties and Security

- Randomness: The key must be truly random and at least as long as the message. Any predictability in the key compromises security.
- Key Use: The key must be used only once (hence the name "one-time pad"). Reusing the key can allow an attacker to deduce the message
- Key Distribution: The key must be securely distributed to both the sender and the receiver. This is often the most challenging aspect of using the Vernam cipher.

### Research materials

- [Gilbert Vernam (wikipedia)](<https://en.wikipedia.org/wiki/Gilbert_Vernam#:~:text=and%20kept%20secret.-,The%20Vernam%20cipher,or%22%20(XOR)%20function.>)
- [Bitwise (XOR) - Mozilla web docs](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Operators/Bitwise_XOR)
- [How to convert binary to hexadecimal](https://www.youtube.com/watch?v=jVPT2Lwelko)
- And a lot of Googling :D

---

### Example

Given the message "LOGIC"

Step 1. Get the ASCII value of each character (Looking at the ASCII table)

```
ASCII value of L is 76
```

Step 2. Convert each ASCII value to binary

```
 76 / 2 = 38 r 0
 38 / 2 = 19 r 0
 19 / 2 =  9 r 1
  9 / 2 =  4 r 1
  4 / 2 =  2 r 0
  2 / 2 =  1 r 0
  1 / 2 =  0 r 1
               0

Binary representation of L = 01001100
```

Step 3. Generate a random key in binary format

```
Assuming the generated random binary key for L is 10101111
```

Step 4. Generate the ciphertext for each character by calculating the XOR of the in binary and the random key of the character

```sh
# Using the rule of XOR which states that:
# - Two same values result to 0
# - Two different values result to 1

Binary representation of L = 01001100
Random key for L           = 10101111
-------------------------------------
XOR of L i.e ciphertext    = 11100011
```

##### Complete result

```
Character | ASCII value | Binary representation (Pi) | Random key (Ki) (binary) | Ciphertext(Ci) (XOR)
----------+-------------+----------------------------+--------------------------+------------------
L         | 76          | 01001100                   | 10101111                 | 11100011
O         | 79          | 01001111                   | 01101101                 | 00100010
G         | 71          | 01000111                   | 10010101                 | 11010010
I         | 73          | 01001001                   | 01101011                 | 00100010
C         | 67          | 01000011                   | 11010100                 | 10010111
```

> Note: The ciphertext is the XOR of the binary representation and random key of the character

Result:

```sh
# The secret key
10101111 01101101 10010101 01101011 11010100

# The ciphertext or encryption of "LOGIC" is:
11100011 00100010 11010010 001000010 10010111

-----------------------------------------------------

# Optionally, we can go further by converting the key and ciphertext into readable format e.g hexadecimal format

# Hexadecimal representation of the secret key
AF 6D 95 6B D4

# Hexadecimal representation of the ciphertext
E3 22 D2 42 97
```

#### Let's decrypt the message

Using the secret key
`10101111 01101101 10010101 01101011 11010100`

Lets decrypt the message from the below binary representation of the ciphertext:
`11100011 00100010 11010010 001000010 10010111`

```
Ciphertext(Ci) | Secret key (Ki) | Message (XOR) | ASCII value | Character
---------------+-----------------+---------------+-------------+------------+
11100011       | 10101111        | 01001100      | 76          | L
00100010       | 01101101        | 01001111      | 79          | O
11010010       | 10010101        | 01000111      | 71          | G
00100010       | 01101011        | 01001001      | 73          | I
10010111       | 11010100        | 01000011      | 67          | C
```
