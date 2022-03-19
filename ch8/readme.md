## benchmark 100个连接100000并发测试结果

10字节
```
redis-benchmark -d 10  -c 100 -n 100000 -t get,set
```

```
====== SET ======
  100000 requests completed in 1.12 seconds
  100 parallel clients
  10 bytes payload
  89445.44 requests per second

====== GET ======
  100000 requests completed in 1.17 seconds
  100 parallel clients
  10 bytes payload
  85397.09 requests per second
```

20字节
```
redis-benchmark -d 20  -c 100 -n 100000 -t get,set
```

```
====== SET ======
  100000 requests completed in 1.17 seconds
  100 parallel clients
  20 bytes payload
  85616.44 requests per second

====== GET ======
  100000 requests completed in 1.12 seconds
  100 parallel clients
  20 bytes payload
  89686.10 requests per second
```

50字节
```
redis-benchmark -d 50  -c 100 -n 100000 -t get,set
```

```
====== SET ======
  100000 requests completed in 1.18 seconds
  100 parallel clients
  50 bytes payload
  84961.77 requests per second

====== GET ======
  100000 requests completed in 1.18 seconds
  100 parallel clients
  50 bytes payload
  84459.46 requests per second
```

100字节
```
redis-benchmark -d 100  -c 100 -n 100000 -t get,set
```

```
====== SET ======
  100000 requests completed in 1.18 seconds
  100 parallel clients
  100 bytes payload
  84530.86 requests per second

====== GET ======
  100000 requests completed in 1.35 seconds
  100 parallel clients
  100 bytes payload
  73964.50 requests per second
```

200字节
```
redis-benchmark -d 100  -c 100 -n 100000 -t get,set
```

```
====== SET ======
  100000 requests completed in 1.18 seconds
  100 parallel clients
  200 bytes payload
  84745.77 requests per second

====== GET ======
  100000 requests completed in 1.70 seconds
  100 parallel clients
  200 bytes payload
  58962.27 requests per second
```

1024字节=1k
```
redis-benchmark -d 1024  -c 100 -n 100000 -t get,set
```

```
====== SET ======
  100000 requests completed in 1.15 seconds
  100 parallel clients
  1024 bytes payload
  86730.27 requests per second

====== GET ======
  100000 requests completed in 1.40 seconds
  100 parallel clients
  1024 bytes payload
  71225.07 requests per second
```


5120字节=5k

```
redis-benchmark -d 5120  -c 100 -n 100000 -t get,set
```

```
====== SET ======
  100000 requests completed in 1.32 seconds
  100 parallel clients
  5120 bytes payload
  75815.01 requests per second

====== GET ======
  100000 requests completed in 1.16 seconds
  100 parallel clients
  5120 bytes payload
  85984.52 requests per second
```

51200字节=50k

```
redis-benchmark -d 51200  -c 100 -n 100000 -t get,set
```
```
====== SET ======
  100000 requests completed in 6.41 seconds
  100 parallel clients
  51200 bytes payload
  15600.62 requests per second

====== GET ======
  100000 requests completed in 5.73 seconds
  100 parallel clients
  51200 bytes payload
  17448.96 requests per second
```


### info memory信息

初始使用1.00M
```shell script
# Memory
used_memory:1050816
used_memory_human:1.00M
used_memory_rss:2002944
used_memory_rss_human:1.91M
used_memory_peak:1050816
used_memory_peak_human:1.00M
used_memory_peak_perc:106.38%
used_memory_overhead:1037534
used_memory_startup:987840
used_memory_dataset:13282
used_memory_dataset_perc:21.09%
allocator_allocated:987840
allocator_active:1965056
allocator_resident:1965056
total_system_memory:8589934592
total_system_memory_human:8.00G
used_memory_lua:37888
used_memory_lua_human:37.00K
used_memory_scripts:0
used_memory_scripts_human:0B
number_of_cached_scripts:0
maxmemory:0
maxmemory_human:0B
maxmemory_policy:noeviction
allocator_frag_ratio:1.99
allocator_frag_bytes:977216
allocator_rss_ratio:1.00
allocator_rss_bytes:0
rss_overhead_ratio:1.02
rss_overhead_bytes:37888
mem_fragmentation_ratio:2.03
mem_fragmentation_bytes:1015104
mem_not_counted_for_evict:0
mem_replication_backlog:0
mem_clients_slaves:0
mem_clients_normal:49694
mem_aof_buffer:0
mem_allocator:libc
active_defrag_running:0
lazyfree_pending_objects:0

```

写入2w个key,value键值对，占2.47M，key为key_{#d}，#d从1~20000，value为随机整数，0~100000000000范围内，平均每个key占130字节
```shell script
# Memory
# Memory
used_memory:2592960
used_memory_human:2.47M
used_memory_rss:3260416
used_memory_rss_human:3.11M
used_memory_peak:2592960
used_memory_peak_human:2.47M
used_memory_peak_perc:101.77%
used_memory_overhead:2099678
used_memory_startup:987840
used_memory_dataset:493282
used_memory_dataset_perc:30.73%
allocator_allocated:2547840
allocator_active:3222528
allocator_resident:3222528
total_system_memory:8589934592
total_system_memory_human:8.00G
used_memory_lua:37888
used_memory_lua_human:37.00K
used_memory_scripts:0
used_memory_scripts_human:0B
number_of_cached_scripts:0
maxmemory:0
maxmemory_human:0B
maxmemory_policy:noeviction
allocator_frag_ratio:1.26
allocator_frag_bytes:674688
allocator_rss_ratio:1.00
allocator_rss_bytes:0
rss_overhead_ratio:1.01
rss_overhead_bytes:37888
mem_fragmentation_ratio:1.28
mem_fragmentation_bytes:712576
mem_not_counted_for_evict:0
mem_replication_backlog:0
mem_clients_slaves:0
mem_clients_normal:49694
mem_aof_buffer:0
mem_allocator:libc
active_defrag_running:0
lazyfree_pending_objects:0
```
写入2w个key,value键值对，占用3.08M，key为key_{#d}，#d从1~20000，value为字符串，长度14~17之间，平均每个key占据162字节

```shell script
# Memory
used_memory:3232960
used_memory_human:3.08M
used_memory_rss:4472832
used_memory_rss_human:4.27M
used_memory_peak:3232960
used_memory_peak_human:3.08M
used_memory_peak_perc:101.42%
used_memory_overhead:2100702
used_memory_startup:988864
used_memory_dataset:1132258
used_memory_dataset_perc:50.45%
allocator_allocated:3187840
allocator_active:4434944
allocator_resident:4434944
total_system_memory:8589934592
total_system_memory_human:8.00G
used_memory_lua:37888
used_memory_lua_human:37.00K
used_memory_scripts:0
used_memory_scripts_human:0B
number_of_cached_scripts:0
maxmemory:0
maxmemory_human:0B
maxmemory_policy:noeviction
allocator_frag_ratio:1.39
allocator_frag_bytes:1247104
allocator_rss_ratio:1.00
allocator_rss_bytes:0
rss_overhead_ratio:1.01
rss_overhead_bytes:37888
mem_fragmentation_ratio:1.40
mem_fragmentation_bytes:1284992
mem_not_counted_for_evict:0
mem_replication_backlog:0
mem_clients_slaves:0
mem_clients_normal:49694
mem_aof_buffer:0
mem_allocator:libc
active_defrag_running:0
lazyfree_pending_objects:0
```
