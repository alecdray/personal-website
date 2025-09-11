run: build
	go run ./cmd/server/main.go

build/templ:
	templ generate

build/tailwind:
	npx @tailwindcss/cli -i ./static/src/main.css -o ./static/public/main.css

build/server:
	go build -v -o ./bin/app ./cmd/server

build:
	make build/templ build/tailwind build/server

# run templ generation in watch mode to detect all .templ files and 
# re-create _templ.txt files on change, then send reload event to browser. 
# Default url: http://localhost:7331
dev/templ:
	templ generate --watch --proxy="http://localhost:8080" --open-browser=false -v

# run air to detect any go file changes to re-build and re-run the server.
dev/server:
	go run github.com/air-verse/air@v1.63.0 \
	--build.cmd "go build -v -o ./bin/app ./cmd/server" --build.bin "./bin/app" --build.delay "100" \
	--build.exclude_dir "node_modules" \
	--build.include_ext "go" \
	--build.stop_on_error "false" \
	--misc.clean_on_exit true

# run tailwindcss to generate the styles.css bundle in watch mode.
dev/tailwind:
	npx --yes @tailwindcss/cli -i ./static/src/main.css -o ./static/public/main.css --watch

# watch for any js or css change in the assets/ folder, then reload the browser via templ proxy.
dev/sync_assets:
	go run github.com/air-verse/air@v1.63.0 \
	--build.cmd "templ generate --notify-proxy" \
	--build.bin "true" \
	--build.delay "100" \
	--build.exclude_dir "node_modules" \
	--build.include_ext "tmpl,css"

# start all watch processes in parallel.
dev: 
	make -j4 dev/templ dev/server dev/tailwind dev/sync_assets