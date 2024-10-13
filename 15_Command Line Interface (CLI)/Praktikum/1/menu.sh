#!/bin/bash

mkdir -p foods drinks snacks

touch foods/menu.txt drinks/menu.txt snacks/menu.txt

echo "nasi goreng" >> foods/menu.txt
echo "ayam goreng" >> foods/menu.txt
echo "bubur ayam" >> foods/menu.txt

echo "kopi susu" >> drinks/menu.txt
echo "susu oat" >> drinks/menu.txt
echo "es teh" >> drinks/menu.txt

echo "kentang goreng - 15000" >> snacks/menu.txt
echo "roti bakar coklat - 25000" >> snacks/menu.txt
echo "roti bakar keju - 25000" >> snacks/menu.txt
echo "mochi - 7500" >> snacks/menu.txt
echo "es krim vanilla - 4000" >> snacks/menu.txt
echo "es krim stroberi - 4500" >> snacks/menu.txt
echo "coklat ayam thunder - 2000" >> snacks/menu.txt
echo "wafer pisang - 2000" >> snacks/menu.txt
echo "ciki sate kambing - 2000" >> snacks/menu.txt
echo "kebab - 23000" >> snacks/menu.txt

echo "Folders and menu files created successfully!"
