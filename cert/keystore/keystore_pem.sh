keytool -importkeystore -srckeystore kafka.keystore.jks -destkeystore kafka.keystore.p12 -srcstoretype jks -deststoretype pkcs12 -srcstorepass \secret123 -deststorepass \secret123 -srcalias localhost -destalias localhost
openssl pkcs12 -in kafka.keystore.p12 -out kafka.pem -nodes -passin pass:\secret123
