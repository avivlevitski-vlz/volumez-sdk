# Capacityoptimization

Choosing “Capacity” directs Volumez to prefer using capacity-saving methods (such as compression, deduplication, erasure coding and thin provisioning) where relevant, in order to consume the minimum amount of raw media. Using such methods might take some CPU cycles, and might reduce the performance of your volumes (it will still be within the range you specified). Choosing “Balanced” directs Volumez to prefer using some capacity-saving methods where relevant, in order to use less raw media, while consuming a small amount of CPU cycles. “Performance Optimized” directs Volumez to avoid using capacity-saving any methods (such as compression and deduplication) that reduce media consumption. This way applications can get the optimal performance from their media, however more raw media might be consumed to provision Performance-Optimized volumes.


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `CapacityoptimizationCapacity`    | capacity                          |
| `CapacityoptimizationBalanced`    | balanced                          |
| `CapacityoptimizationPerformance` | performance                       |