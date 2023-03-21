NAME = dicecalc
ZIP_NAME = handler.zip

create:
	@yc serverless function create --name=$(NAME)
	@yc serverless function allow-unauthenticated-invoke --name=$(NAME)

delete:
	@yc serverless function delete --name=$(NAME)

display:
	@yc serverless function get --name=$(NAME)
	@yc serverless function version list --function-name=$(NAME)

.PHONY: $(ZIP_NAME)
$(ZIP_NAME):
	@zip -r $(ZIP_NAME) *.go go.mod go.sum
	@touch $(ZIP_NAME)

deploy: $(ZIP_NAME)
	@yc serverless function version create \
		--function-name=$(NAME) \
		--runtime golang119 \
		--entrypoint handler.Handle \
		--memory 128m \
		--execution-timeout 2s \
		--source-path $(ZIP_NAME)

invoke:
	@yc serverless function invoke --name=$(NAME) --data '{"httpMethod":"GET"}'

logs:
	@yc serverless function logs --name=$(NAME)
