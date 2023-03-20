NAME = dicecalc
ZIP_NAME = handler.zip

create:
	@yc serverless function create --name=$(NAME)

delete:
	@yc serverless function delete --name=$(NAME)

display:
	@yc serverless function get --name=$(NAME)
	@yc serverless function version list --function-name=$(NAME)
	
pack: $(ZIP_NAME)
	@zip -r $(ZIP_NAME) *.go
	@touch $(ZIP_NAME)

deploy: pack
	@yc serverless function version create \
		--function-name=$(NAME) \
		--runtime golang119 \
		--entrypoint handler.Handle \
		--memory 128m \
		--execution-timeout 2s \
		--source-path $(ZIP_NAME)

invoke:
	@yc serverless function invoke --name=$(NAME)

logs:
	@yc serverless function logs --name=$(NAME)
