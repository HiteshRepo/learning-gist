## HTTPS illustrated
1. Without 'https' the communication between a browser and server is simply plain text.
2. This means the password or credit card details sent over the internet can be intercepted by whomever has the ability to do so.
3. 'https' is designed to overcome the above issue. This makes data unreadable by anyone other than sender/receiver.
4. 'https' is an extension of 'http' protocol. 
5. In 'https', data is sent in an encrypted form using 'tls'.
6. TLS: transport layer security.
7. If data is intercepted by a hacker, all they can see is jumbled data.

## TLS Handshake steps:
1. TCP handshake between browser (client) and server. TCP SYN (client -> server), TCP SYN + ACK (server -> client), TCP ACK (client -> server).
2. Client sends a 'client hello' to server. In this client sends following things: 
   1. Supported TLS version (tls1.2, tls1.3, ...). 
   2. Supported Cipher suite (algorithm used to encrypt data).
3. Server now chooses the TLS version and Cipher suite based on options provided by the client.
4. Server then sends the chosen constructs via 'server hello' back to the client.
5. Server the sends a certificate to the client. Certificate includes a lot of things but one of the important thing is the public key for the server.
6. Client makes use of the public key to do 'Asymmetric Encryption'.
7. 'Asymmetric Encryption' - a data encrypted by public key can only be decrypted by a private key.
8. This concludes 'server hello' and sever sends 'server hello done' request to client.
9. At this point, client has a server certificate and public key. Also both client and server have agreed on TLS version and Cipher suite to use.
10. Now in this subsequent step, the client and server come up with a shared encryption key to encrypt data.
11. Using 'Asymmetric Encryption', client sends and encryption key to server safely.
12. This is done via 'client key exchange' message. The details differ on the usage of CipherSuite.
13. For example if RSA cipher suite is used, below exchange happens:
    1. Client to Server: 'Client Key Exchange'
    2. Client to Server: 'Change Cipher Spec'
    3. Client to Server: 'Finished'
    4. Server to Client: 'Change Cipher Spec'
    5. Server to Client: 'Finished'
14. RSA encryption key generation steps:
    1. Client generates an encryption key 'Session Key'.
    2. Encrypts the 'Session Key' with server's public key.
    3. Sends the encrypted 'Session Key' to server.
    4. Server receives the encrypted 'Session Key' and decrypts it with private key.
    5. Now both sides have 'Session Key'
15. Now the data transmission happens where they encrypt the data using agreed upon 'Session key' and CipherSuite and exchange data back and forth securely.
16. Why not use 'Asymmetric Encryption' for data transmission as well? Because it is computationally expensive. Not suitable for Bulk data transmission.
17. The handshake we discussed is applied in TLS1.2 and at the moment TLS1.3 is used.
18. The number of n/w round trip us TLS1.2 is 2, which is reduced in TLS1.3 to just 1.
19. RSA is not supported for 'Asymmetric Encryption' in TLS1.3, rather 'Diffie-Hellman' is a more common way to exchange 'Session Key'.
20. 'Diffie-Hellman' is a complicated advanced math involving large prime numbers to derive shared encryption key without the need to transmit a public key over the n/w.