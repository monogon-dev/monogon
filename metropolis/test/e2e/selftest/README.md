self-test image
===

This image is used by the Metropolis E2E tests to perform some cluster-internal
tests. See //metropolis/test/e2e:main_test.go for usage.

The image should be run as a Kubernetes Job, and should return 0 if all tests
have passed. If the job fails, its last log line will be printed.