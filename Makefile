.PHONY: run-go run-vite dev install-web clean build build-vite build-go AS_OF
AS_OF ?= 2026-10-21

# Pass AS_OF to Vite through the environment. This works in any shell, including cmd.exe on Windows.
export VITE_AS_OF = $(AS_OF)

# Install and build the frontend before starting both development servers.
dev:
	$(MAKE) -j2 run-go run-vite

# Install frontend dependencies from the lockfile.
install-web:
	npm ci --prefix ./web

# Run the Go backend server 
run-go:
	go run ./api/main.go $(AS_OF)
	
# Run the Vite frontend dev server
run-vite:
	npm run dev --prefix ./web
	
# Build both applications for production
build: build-vite build-go

build-vite:
	npm run build --prefix ./web

build-go:
	go build -o bin/server ./api/main.go

clean:
	rm -rf bin/
	rm -rf ./web/dist