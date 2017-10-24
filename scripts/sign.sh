file=$1
key=$2
output=$3

echo "$key" | base64 -d > keyfile
openssl dgst -sha256 -sign keyfile -out $output $file
rm -rf keyfile