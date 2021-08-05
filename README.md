
Update eu Trust list
wget https://ec.europa.eu/tools/lotl/eu-lotl.xml
go run cmd/update-trustlist-eu/main.go files/eu-lotl.xml internal/gen/certs.go.templ internal/gen/eucerts.go 
