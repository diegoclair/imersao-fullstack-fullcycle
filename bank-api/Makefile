create-migration:
	echo "creating migration for $${name:?please, you need to pass the migration name like name=migration_name \$$name}"
	npm run typeorm migration:create -- -n ${name}	

run-migrations:
	npm run typeorm migration:run

create-model:
	echo "creating model for $${name:?please, you need to pass the model name like name=model_name \$$name}"
	npm run typeorm entity:create -- -n ${name}
	./rename-files.sh ./src/models/ ${name}