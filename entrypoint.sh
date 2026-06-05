#!/bin/sh

echo "Aguardando MySQL..."

until nc -z mysql 3306
do
  sleep 2
done

echo "MySQL disponível"

echo "Executando migrations..."

make migrate

echo "Iniciando aplicação..."

./app