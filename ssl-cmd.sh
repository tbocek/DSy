# create CA
openssl req -new -nodes -x509 -keyout ca.key -newkey rsa:4096 -out ca.pem -days 3650

# create signing request
openssl req -new -nodes -keyout user.key -newkey rsa:4096 -out user.pem
# sign the previous request
openssl x509 -req -days 365 -in user.pem -CA ca.pem -CAkey ca.key -set_serial 01 -out user.crt
# convert to pkcs12 format
openssl pkcs12 -export -out user.pfx -inkey user.key -in user.crt -certfile ca.pem

openssl req -new -nodes -keyout server.key -newkey rsa:4096 -out server.pem
openssl x509 -req -days 365 -in server.pem -CA ca.pem -CAkey ca.key -set_serial 02 -out server.crt -extfile extensions.ext