#!/bin/sh

# Generated by convert.go. DO NOT EDIT.

DIR=$(dirname $BASH_SOURCE)

mongoimport --uri=mongodb://localhost/monila --collection=actor --drop --maintainInsertionOrder $DIR/actor.json
mongoimport --uri=mongodb://localhost/monila --collection=address --drop --maintainInsertionOrder $DIR/address.json
mongoimport --uri=mongodb://localhost/monila --collection=category --drop --maintainInsertionOrder $DIR/category.json
mongoimport --uri=mongodb://localhost/monila --collection=city --drop --maintainInsertionOrder $DIR/city.json
mongoimport --uri=mongodb://localhost/monila --collection=country --drop --maintainInsertionOrder $DIR/country.json
mongoimport --uri=mongodb://localhost/monila --collection=customer --drop --maintainInsertionOrder $DIR/customer.json
mongoimport --uri=mongodb://localhost/monila --collection=film_actor --drop --maintainInsertionOrder $DIR/film_actor.json
mongoimport --uri=mongodb://localhost/monila --collection=film_category --drop --maintainInsertionOrder $DIR/film_category.json
mongoimport --uri=mongodb://localhost/monila --collection=film --drop --maintainInsertionOrder $DIR/film.json
mongoimport --uri=mongodb://localhost/monila --collection=inventory --drop --maintainInsertionOrder $DIR/inventory.json
mongoimport --uri=mongodb://localhost/monila --collection=language --drop --maintainInsertionOrder $DIR/language.json
mongoimport --uri=mongodb://localhost/monila --collection=rental --drop --maintainInsertionOrder $DIR/rental.json
mongoimport --uri=mongodb://localhost/monila --collection=staff --drop --maintainInsertionOrder $DIR/staff.json
mongoimport --uri=mongodb://localhost/monila --collection=store --drop --maintainInsertionOrder $DIR/store.json
