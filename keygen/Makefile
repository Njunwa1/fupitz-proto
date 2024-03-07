protoc:
	@echo "Generating protocol buffer files..."
	@cd proto && protoc --go_out=../golang --go_opt=paths=source_relative \
              --go-grpc_out=../proto/golang --go-grpc_opt=paths=source_relative \
              --grpc-gateway_out=../proto/golang --grpc-gateway_opt paths=source_relative \
              --grpc-gateway_opt generate_unbound_methods=true \
              ./**/*.proto
	@echo "Done."

hellow
	@echo "Hellow"