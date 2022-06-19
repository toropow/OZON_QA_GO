.PHONY: load-example
load-example:
	docker run -v $(pwd):/var/loadtest direvius/yandex-tank -c ./load.yaml

.PHONY: load-device-api
load-device-api:
	docker run --network=host -v $(pwd):/var/loadtest direvius/yandex-tank -c ./load_device_api.yaml
