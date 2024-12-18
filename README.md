# ArchDemo
Need public repo for hackday multi arch demo

docker buildx build --push --platform linux/amd64,linux/arm64 --builder=multi-arch -t ttl.sh/my-arch-test-image:latest .

docker manifest create ttl.sh/ben-multi-arch-image:multi \
  --amend ttl.sh/ben-multi-arch-image:amd64 \
  --amend ttl.sh/ben-multi-arch-image:arm64

docker manifest push ttl.sh/ben-multi-arch-image:multi
