Project status: On-hold

- Not usable yet, no releases
- Configuration file loading implemented
- Metric firing TODO

# Prometheus Canary
Intentionally trigger Prometheus alerts to test alerting.

# Table Of Contents
- [Overview](#overview)
- [Build](#build)
- [Release Checklist](#release-checklist)

# Overview
Publishes a single metric at a set interval. A Prometheus alert should be setup
to fire when this metric is found. 

This can be used to test the entire Prometheus alert flow.

# Build
Golang application. Build with:

```
% go build .
```

# Release Checklist
To create a new release:

1. Push to master
2. Build the tool for Linux x86_64:
   ```
   go build -o prometheus-canary-vmajor.minor.patch-linux-x86_64 .
   ```
3. Create a new GitHub release
    1. Tag it `vmajor.minor.patch`
	2. Name it `vmajor.minor.patch`
	3. Have a description like so:
	   ```
	   Short one sentence release description.
	   
	   # Change log
	   - More detailed change log
	   
	   # Files
	   | File | SHA256 Checksum |
       | - | - |
       | `prometheus-canary-vmajor.minor.patch-linux-x86_64` | `checksum` |
       ```
