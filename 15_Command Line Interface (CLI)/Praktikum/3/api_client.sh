#!/bin/bash

read -p "Masukkan endpoint: " endpoint
read -p "Masukkan method HTTP (GET, POST, PUT, DELETE): " method
read -p "Masukkan request body (biarkan kosong jika tidak diperlukan): " body
read -p "Masukkan tipe dari request body (contoh: application/json): " content_type

if [[ -z "$body" ]]; then
    response=$(curl -s -X "$method" "$endpoint")
else
    response=$(curl -s -X "$method" -H "Content-Type: $content_type" -d "$body" "$endpoint")
fi

echo "Respons dari API:"
echo "$response"
