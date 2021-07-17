run:
	docker-compose run payment build

inside:
	docker-compose run payment bash

stop:
	docker-compose down

clean:
	docker exec -it localmongo1 mongo my-sales-champion --eval "printjson(db.dropDatabase())"