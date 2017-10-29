FROM mongo
COPY customers.json /customers.json
CMD mongoimport --host mongodb --db exampleDb --collection contacts --type json --file /customers.json --jsonArray
