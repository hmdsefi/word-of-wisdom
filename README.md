# Word of Wisdom

## Challenge-response Protocol
Challenge-response protocols involve sending a challenge to a client, which must then compute a 
response that can be used to authenticate the client to a server. There are various algorithms 
that can be used to implement challenge-response protocols, depending on the specific requirements
of the system.

Word of Wisdom TCP server is using HMAC algorithm to protect the server from the DDOS attack.

## Why using HMAC for CRAM?
One commonly used algorithm for challenge-response protocols is the HMAC algorithm, which is a 
type of message authentication code (MAC) that uses a cryptographic hash function in combination
with a secret key to verify the integrity and authenticity of a message. In this algorithm, the 
server generates a challenge, and sends it to the client along with a secret key. The client then
computes an HMAC using the challenge and the secret key, and sends the HMAC back to the server
as the response. The server can then verify the response by computing the HMAC using the same
secret key and comparing it to the response sent by the client. If the HMACs match, the client 
is authenticated.

Another reason for using HMAC was because there were not much details in the requirements I received. 
As a result, I chose the most common and simple algorithm to implement.