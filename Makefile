# thor
TAG=$(shell git tag | sort -bt"." -k1,1 -k2,2n -k3,3n |tail -n 1)

fmt:
	find . -name "*.go" -type f -not -path "./vendor/*" -not -path "./.idea/*" -not -path "./.git/*" | xargs gofmt -w -s -d
	find . -name "*.go" -type f -not -path "./vendor/*" -not -path "./.idea/*" -not -path "./.git/*" | xargs goimports -w
	find . -name "*.go" -type f -not -path "./vendor/*" -not -path "./.idea/*" -not -path "./.git/*" -print0 | xargs -0 misspell
