variable "REPO" {
  default = "k33g"
}

variable "TAG" {
  default = "0.0.0"
}

group "default" {
  targets = ["tiny-service-image"]
}

target "tiny-service-image" {
  context = "."
  platforms = [
    "linux/amd64",
    "linux/arm64"
  ]
  tags = ["${REPO}/tiny-service:${TAG}"]
}

/*
docker buildx ls
docker login
docker buildx bake --push --file docker-bake.hcl
*/
