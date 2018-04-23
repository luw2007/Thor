# thor
TAG=$(shell git tag | sort -bt"." -k1,1 -k2,2n -k3,3n |tail -n 1)

fmt:
	find . -name "*.go" -type f -not -path "./vendor/*" -not -path "./.idea/*" -not -path "./.git/*" | xargs gofmt -w -s -d
	find . -name "*.go" -type f -not -path "./vendor/*" -not -path "./.idea/*" -not -path "./.git/*" | xargs goimports -w
	find . -name "*.go" -type f -not -path "./vendor/*" -not -path "./.idea/*" -not -path "./.git/*" -print0 | xargs -0 misspell

kit-manager:
	rm -rf manager/{cmd,pkg}
	kit n s manager
	cp manager/service.go manager/pkg/service/service.go
	kit g s manager --gorilla
	kit g c manager

kit-worker:
	rm -rf worker/{cmd,pkg}
	kit n s worker
	cp worker/service.go worker/pkg/service/service.go
	kit g s worker --gorilla
	kit g c worker