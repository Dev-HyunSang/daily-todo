server-run:
	go run server/main.go	
front-run:
	cd front && npm run dev
gen:
	go generate ./ent